package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/seal-io/seal/pkg/auths/session"
	"github.com/seal-io/seal/pkg/dao"
	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/types/object"
	"github.com/seal-io/seal/pkg/dao/types/status"
	deptypes "github.com/seal-io/seal/pkg/deployer/types"
	"github.com/seal-io/seal/utils/log"
	"github.com/seal-io/seal/utils/strs"
)

const annotationSubjectIDName = "seal.io/subject-id"

// Options for deploy or destroy.
type Options struct {
	TlsCertified bool
	Tags         []string
}

func Create(
	ctx context.Context,
	mc model.ClientSet,
	dp deptypes.Deployer,
	entity *model.Service,
	opts Options,
) (*model.ServiceOutput, error) {
	err := mc.WithTx(ctx, func(tx *model.Tx) (err error) {
		// TODO(thxCode): generated by entc.
		status.ServiceStatusDeployed.Unknown(entity, "")
		entity.Status.SetSummary(status.WalkService(&entity.Status))

		entity, err = tx.Services().Create().
			Set(entity).
			SaveE(ctx, dao.ServiceDependenciesEdgeSave)

		return err
	})
	if err != nil {
		return nil, err
	}

	// Deploy service.
	err = Apply(ctx, mc, dp, entity, Options{
		TlsCertified: opts.TlsCertified,
		Tags:         opts.Tags,
	})
	if err != nil {
		return nil, err
	}

	return model.ExposeService(entity), nil
}

func UpdateStatus(
	ctx context.Context,
	mc model.ClientSet,
	entity *model.Service,
) error {
	entity.Status.SetSummary(status.WalkService(&entity.Status))

	err := mc.Services().UpdateOne(entity).
		SetStatus(entity.Status).
		Exec(ctx)
	if err != nil && !model.IsNotFound(err) {
		return err
	}

	return nil
}

func Apply(
	ctx context.Context,
	mc model.ClientSet,
	dp deptypes.Deployer,
	entity *model.Service,
	opts Options,
) (err error) {
	logger := log.WithName("service")

	defer func() {
		if err == nil {
			return
		}
		// Update a failure status.
		status.ServiceStatusDeployed.False(entity, err.Error())

		uerr := UpdateStatus(ctx, mc, entity)
		if uerr != nil {
			logger.Errorf("error updating status of service %s: %v",
				entity.ID, uerr)
		}
	}()

	if !status.ServiceStatusDeployed.IsUnknown(entity) {
		return fmt.Errorf("service status is not deploying, service: %s", entity.ID)
	}

	if dp == nil {
		return errors.New("deployer is not set")
	}

	applyOpts := deptypes.ApplyOptions{
		SkipTLSVerify: !opts.TlsCertified,
		Tags:          opts.Tags,
	}

	return dp.Apply(ctx, entity, applyOpts)
}

func Destroy(
	ctx context.Context,
	mc model.ClientSet,
	dp deptypes.Deployer,
	entity *model.Service,
	opts Options,
) (err error) {
	logger := log.WithName("service")

	defer func() {
		if err == nil {
			return
		}
		// Update a failure status.
		status.ServiceStatusDeleted.False(entity, err.Error())

		uerr := UpdateStatus(ctx, mc, entity)
		if uerr != nil {
			logger.Errorf("error updating status of service %s: %v",
				entity.ID, uerr)
		}
	}()

	if dp == nil {
		return errors.New("deployer is not set")
	}

	// Check dependants.
	dependants, err := dao.GetServiceDependantNames(ctx, mc, entity)
	if err != nil {
		return err
	}

	if len(dependants) > 0 {
		msg := fmt.Sprintf("Waiting for dependants to be deleted: %s", strs.Join(", ", dependants...))
		if !status.ServiceStatusProgressing.IsUnknown(entity) ||
			status.ServiceStatusDeleted.GetMessage(entity) != msg {
			// Mark status to deleting with dependency message.
			status.ServiceStatusDeleted.Reset(entity, msg)
			status.ServiceStatusProgressing.Unknown(entity, "")

			if err = UpdateStatus(ctx, mc, entity); err != nil {
				return fmt.Errorf("failed to update service status: %w", err)
			}
		}

		return nil
	} else {
		// Mark status to deleting.
		status.ServiceStatusDeleted.Reset(entity, "")
		status.ServiceStatusProgressing.True(entity, "Resolved dependencies")

		if err = UpdateStatus(ctx, mc, entity); err != nil {
			return fmt.Errorf("failed to update service status: %w", err)
		}
	}

	destroyOpts := deptypes.DestroyOptions{
		SkipTLSVerify: !opts.TlsCertified,
	}

	return dp.Destroy(ctx, entity, destroyOpts)
}

func GetSubjectID(entity *model.Service) (object.ID, error) {
	if entity == nil {
		return "", fmt.Errorf("service is nil")
	}

	subjectIDStr := entity.Annotations[annotationSubjectIDName]

	return object.ID(subjectIDStr), nil
}

func SetSubjectID(ctx context.Context, services ...*model.Service) error {
	sj, err := session.GetSubject(ctx)
	if err != nil {
		return err
	}

	for i := range services {
		if services[i].Annotations == nil {
			services[i].Annotations = make(map[string]string)
		}
		services[i].Annotations[annotationSubjectIDName] = string(sj.ID)
	}

	return nil
}

// SetServiceStatusScheduled sets the status of the service to scheduled.
func SetServiceStatusScheduled(ctx context.Context, mc model.ClientSet, entity *model.Service) error {
	if entity == nil {
		return fmt.Errorf("service is nil")
	}

	dependencyNames := dao.ServiceRelationshipGetDependencyNames(entity)

	msg := ""
	if len(dependencyNames) > 0 {
		msg = fmt.Sprintf("Waiting for dependent services to be ready: %s", strs.Join(", ", dependencyNames...))
	}

	status.ServiceStatusProgressing.Reset(entity, msg)
	entity.Status.SetSummary(status.WalkService(&entity.Status))

	return mc.Services().UpdateOne(entity).
		SetStatus(entity.Status).
		Exec(ctx)
}

// CreateScheduledServices creates scheduled services.
func CreateScheduledServices(ctx context.Context, mc model.ClientSet, entities model.Services) (model.Services, error) {
	results := make(model.Services, 0, len(entities))

	sortedServices, err := TopologicalSortServices(entities)
	if err != nil {
		return nil, err
	}

	for i := range sortedServices {
		entity := sortedServices[i]

		err = mc.WithTx(ctx, func(tx *model.Tx) error {
			// TODO(thxCode): generated by entc.
			status.ServiceStatusDeployed.Unknown(entity, "")
			entity.Status.SetSummary(status.WalkService(&entity.Status))

			entity, err = tx.Services().Create().
				Set(entity).
				SaveE(ctx, dao.ServiceDependenciesEdgeSave)
			if err != nil {
				return err
			}

			return SetServiceStatusScheduled(ctx, tx, entity)
		})
		if err != nil {
			return nil, err
		}

		results = append(results, entity)
	}

	return results, nil
}

// IsStatusReady returns true if the service is ready.
func IsStatusReady(entity *model.Service) bool {
	switch entity.Status.SummaryStatus {
	case "Preparing", "NotReady", "Ready":
		return true
	}

	return false
}

// IsStatusFalse returns true if the service is in error status.
func IsStatusFalse(entity *model.Service) bool {
	switch entity.Status.SummaryStatus {
	case "DeployFailed", "DeleteFailed":
		return true
	case "Progressing":
		return entity.Status.Error
	}

	return false
}

// IsStatusDeleted returns true if the service is deleted.
func IsStatusDeleted(entity *model.Service) bool {
	switch entity.Status.SummaryStatus {
	case "Deleted", "Deleting":
		return true
	}

	return false
}
