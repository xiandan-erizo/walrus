// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/seal-io/walrus/pkg/dao/model/connector"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/schema/intercept"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/utils/json"
)

// ConnectorCreateInput holds the creation input of the Connector entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ConnectorCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create Connector entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Config whether enable finOps, will install prometheus and opencost while enable.
	EnableFinOps bool `path:"-" query:"-" json:"enableFinOps"`
	// Connector config version.
	ConfigVersion string `path:"-" query:"-" json:"configVersion"`
	// Environment type of the connector to apply.
	ApplicableEnvironmentType string `path:"-" query:"-" json:"applicableEnvironmentType"`
	// Type of the connector.
	Type string `path:"-" query:"-" json:"type,cli-table-column"`
	// Category of the connector.
	Category string `path:"-" query:"-" json:"category,cli-table-column"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Connector config data.
	ConfigData crypto.Properties `path:"-" query:"-" json:"configData,omitempty"`
	// Custom pricing user defined.
	FinOpsCustomPricing *types.FinOpsCustomPricing `path:"-" query:"-" json:"finOpsCustomPricing,omitempty"`
}

// Model returns the Connector entity for creating,
// after validating.
func (cci *ConnectorCreateInput) Model() *Connector {
	if cci == nil {
		return nil
	}

	_c := &Connector{
		EnableFinOps:              cci.EnableFinOps,
		ConfigVersion:             cci.ConfigVersion,
		ApplicableEnvironmentType: cci.ApplicableEnvironmentType,
		Type:                      cci.Type,
		Category:                  cci.Category,
		Name:                      cci.Name,
		Description:               cci.Description,
		Labels:                    cci.Labels,
		ConfigData:                cci.ConfigData,
		FinOpsCustomPricing:       cci.FinOpsCustomPricing,
	}

	if cci.Project != nil {
		_c.ProjectID = cci.Project.ID
	}

	return _c
}

// Validate checks the ConnectorCreateInput entity.
func (cci *ConnectorCreateInput) Validate() error {
	if cci == nil {
		return errors.New("nil receiver")
	}

	return cci.ValidateWith(cci.inputConfig.Context, cci.inputConfig.Client, nil)
}

// ValidateWith checks the ConnectorCreateInput entity with the given context and client set.
func (cci *ConnectorCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if cci.Project != nil {
		if err := cci.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				cci.Project = nil
			}
		}
	}

	return nil
}

// ConnectorCreateInputs holds the creation input item of the Connector entities.
type ConnectorCreateInputsItem struct {
	// Config whether enable finOps, will install prometheus and opencost while enable.
	EnableFinOps bool `path:"-" query:"-" json:"enableFinOps"`
	// Connector config version.
	ConfigVersion string `path:"-" query:"-" json:"configVersion"`
	// Environment type of the connector to apply.
	ApplicableEnvironmentType string `path:"-" query:"-" json:"applicableEnvironmentType"`
	// Type of the connector.
	Type string `path:"-" query:"-" json:"type,cli-table-column"`
	// Category of the connector.
	Category string `path:"-" query:"-" json:"category,cli-table-column"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Connector config data.
	ConfigData crypto.Properties `path:"-" query:"-" json:"configData,omitempty"`
	// Custom pricing user defined.
	FinOpsCustomPricing *types.FinOpsCustomPricing `path:"-" query:"-" json:"finOpsCustomPricing,omitempty"`
}

// ValidateWith checks the ConnectorCreateInputsItem entity with the given context and client set.
func (cci *ConnectorCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// ConnectorCreateInputs holds the creation input of the Connector entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ConnectorCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create Connector entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ConnectorCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Connector entities for creating,
// after validating.
func (cci *ConnectorCreateInputs) Model() []*Connector {
	if cci == nil || len(cci.Items) == 0 {
		return nil
	}

	_cs := make([]*Connector, len(cci.Items))

	for i := range cci.Items {
		_c := &Connector{
			EnableFinOps:              cci.Items[i].EnableFinOps,
			ConfigVersion:             cci.Items[i].ConfigVersion,
			ApplicableEnvironmentType: cci.Items[i].ApplicableEnvironmentType,
			Type:                      cci.Items[i].Type,
			Category:                  cci.Items[i].Category,
			Name:                      cci.Items[i].Name,
			Description:               cci.Items[i].Description,
			Labels:                    cci.Items[i].Labels,
			ConfigData:                cci.Items[i].ConfigData,
			FinOpsCustomPricing:       cci.Items[i].FinOpsCustomPricing,
		}

		if cci.Project != nil {
			_c.ProjectID = cci.Project.ID
		}

		_cs[i] = _c
	}

	return _cs
}

// Validate checks the ConnectorCreateInputs entity .
func (cci *ConnectorCreateInputs) Validate() error {
	if cci == nil {
		return errors.New("nil receiver")
	}

	return cci.ValidateWith(cci.inputConfig.Context, cci.inputConfig.Client, nil)
}

// ValidateWith checks the ConnectorCreateInputs entity with the given context and client set.
func (cci *ConnectorCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cci == nil {
		return errors.New("nil receiver")
	}

	if len(cci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if cci.Project != nil {
		if err := cci.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				cci.Project = nil
			}
		}
	}

	for i := range cci.Items {
		if cci.Items[i] == nil {
			continue
		}

		if err := cci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ConnectorDeleteInput holds the deletion input of the Connector entity,
// please tags with `path:",inline"` if embedding.
type ConnectorDeleteInput struct {
	ConnectorQueryInput `path:",inline"`
}

// ConnectorDeleteInputs holds the deletion input item of the Connector entities.
type ConnectorDeleteInputsItem struct {
	// ID of the Connector entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Connector entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// ConnectorDeleteInputs holds the deletion input of the Connector entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ConnectorDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to delete Connector entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ConnectorDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Connector entities for deleting,
// after validating.
func (cdi *ConnectorDeleteInputs) Model() []*Connector {
	if cdi == nil || len(cdi.Items) == 0 {
		return nil
	}

	_cs := make([]*Connector, len(cdi.Items))
	for i := range cdi.Items {
		_cs[i] = &Connector{
			ID: cdi.Items[i].ID,
		}
	}
	return _cs
}

// IDs returns the ID list of the Connector entities for deleting,
// after validating.
func (cdi *ConnectorDeleteInputs) IDs() []object.ID {
	if cdi == nil || len(cdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(cdi.Items))
	for i := range cdi.Items {
		ids[i] = cdi.Items[i].ID
	}
	return ids
}

// Validate checks the ConnectorDeleteInputs entity.
func (cdi *ConnectorDeleteInputs) Validate() error {
	if cdi == nil {
		return errors.New("nil receiver")
	}

	return cdi.ValidateWith(cdi.inputConfig.Context, cdi.inputConfig.Client, nil)
}

// ValidateWith checks the ConnectorDeleteInputs entity with the given context and client set.
func (cdi *ConnectorDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cdi == nil {
		return errors.New("nil receiver")
	}

	if len(cdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Connectors().Query()

	// Validate when deleting under the Project route.
	if cdi.Project != nil {
		if err := cdi.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				cdi.Project = nil
				q.Where(
					connector.ProjectIDIsNil())
			}
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				connector.ProjectID(cdi.Project.ID))
		}
	} else {
		q.Where(
			connector.ProjectIDIsNil())
	}

	ids := make([]object.ID, 0, len(cdi.Items))
	ors := make([]predicate.Connector, 0, len(cdi.Items))
	indexers := make(map[any][]int)

	for i := range cdi.Items {
		if cdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if cdi.Items[i].ID != "" {
			ids = append(ids, cdi.Items[i].ID)
			ors = append(ors, connector.ID(cdi.Items[i].ID))
			indexers[cdi.Items[i].ID] = append(indexers[cdi.Items[i].ID], i)
		} else if cdi.Items[i].Name != "" {
			ors = append(ors, connector.And(
				connector.Name(cdi.Items[i].Name)))
			indexerKey := fmt.Sprint("/", cdi.Items[i].Name)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := connector.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = connector.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			connector.FieldID,
			connector.FieldName,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		indexer := indexers[es[i].ID]
		if indexer == nil {
			indexerKey := fmt.Sprint("/", es[i].Name)
			indexer = indexers[indexerKey]
		}
		for _, j := range indexer {
			cdi.Items[j].ID = es[i].ID
			cdi.Items[j].Name = es[i].Name
		}
	}

	return nil
}

// ConnectorPatchInput holds the patch input of the Connector entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ConnectorPatchInput struct {
	ConnectorQueryInput `path:",inline" query:"-" json:"-"`

	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Annotations holds the value of the "annotations" field.
	Annotations map[string]string `path:"-" query:"-" json:"annotations,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime *time.Time `path:"-" query:"-" json:"createTime,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime *time.Time `path:"-" query:"-" json:"updateTime,omitempty"`
	// Status holds the value of the "status" field.
	Status status.Status `path:"-" query:"-" json:"status,omitempty"`
	// Category of the connector.
	Category string `path:"-" query:"-" json:"category,omitempty"`
	// Type of the connector.
	Type string `path:"-" query:"-" json:"type,omitempty"`
	// Environment type of the connector to apply.
	ApplicableEnvironmentType string `path:"-" query:"-" json:"applicableEnvironmentType,omitempty"`
	// Connector config version.
	ConfigVersion string `path:"-" query:"-" json:"configVersion,omitempty"`
	// Connector config data.
	ConfigData crypto.Properties `path:"-" query:"-" json:"configData,omitempty"`
	// Config whether enable finOps, will install prometheus and opencost while enable.
	EnableFinOps bool `path:"-" query:"-" json:"enableFinOps,omitempty"`
	// Custom pricing user defined.
	FinOpsCustomPricing *types.FinOpsCustomPricing `path:"-" query:"-" json:"finOpsCustomPricing,omitempty"`

	patchedEntity *Connector `path:"-" query:"-" json:"-"`
}

// PatchModel returns the Connector partition entity for patching.
func (cpi *ConnectorPatchInput) PatchModel() *Connector {
	if cpi == nil {
		return nil
	}

	_c := &Connector{
		Name:                      cpi.Name,
		Description:               cpi.Description,
		Labels:                    cpi.Labels,
		Annotations:               cpi.Annotations,
		CreateTime:                cpi.CreateTime,
		UpdateTime:                cpi.UpdateTime,
		Status:                    cpi.Status,
		Category:                  cpi.Category,
		Type:                      cpi.Type,
		ApplicableEnvironmentType: cpi.ApplicableEnvironmentType,
		ConfigVersion:             cpi.ConfigVersion,
		ConfigData:                cpi.ConfigData,
		EnableFinOps:              cpi.EnableFinOps,
		FinOpsCustomPricing:       cpi.FinOpsCustomPricing,
	}

	if cpi.Project != nil {
		_c.ProjectID = cpi.Project.ID
	}

	return _c
}

// Model returns the Connector patched entity,
// after validating.
func (cpi *ConnectorPatchInput) Model() *Connector {
	if cpi == nil {
		return nil
	}

	return cpi.patchedEntity
}

// Validate checks the ConnectorPatchInput entity.
func (cpi *ConnectorPatchInput) Validate() error {
	if cpi == nil {
		return errors.New("nil receiver")
	}

	return cpi.ValidateWith(cpi.inputConfig.Context, cpi.inputConfig.Client, nil)
}

// ValidateWith checks the ConnectorPatchInput entity with the given context and client set.
func (cpi *ConnectorPatchInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := cpi.ConnectorQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	q := cs.Connectors().Query()

	// Validate when querying under the Project route.
	if cpi.Project != nil {
		if err := cpi.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				cpi.Project = nil
				q.Where(
					connector.ProjectIDIsNil())
			}
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				connector.ProjectID(cpi.Project.ID))
		}
	} else {
		q.Where(
			connector.ProjectIDIsNil())
	}

	if cpi.Refer != nil {
		if cpi.Refer.IsID() {
			q.Where(
				connector.ID(cpi.Refer.ID()))
		} else if refers := cpi.Refer.Split(1); len(refers) == 1 {
			q.Where(
				connector.Name(refers[0].String()))
		} else {
			return errors.New("invalid identify refer of connector")
		}
	} else if cpi.ID != "" {
		q.Where(
			connector.ID(cpi.ID))
	} else if cpi.Name != "" {
		q.Where(
			connector.Name(cpi.Name))
	} else {
		return errors.New("invalid identify of connector")
	}

	q.Select(
		connector.WithoutFields(
			connector.FieldAnnotations,
			connector.FieldCreateTime,
			connector.FieldUpdateTime,
			connector.FieldStatus,
		)...,
	)

	var e *Connector
	{
		// Get cache from previous validation.
		queryStmt, queryArgs := q.sqlQuery(setContextOp(ctx, q.ctx, "cache")).Query()
		ck := fmt.Sprintf("stmt=%v, args=%v", queryStmt, queryArgs)
		if cv, existed := cache[ck]; !existed {
			var err error
			e, err = q.Only(ctx)
			if err != nil {
				return err
			}

			// Set cache for other validation.
			cache[ck] = e
		} else {
			e = cv.(*Connector)
		}
	}

	_pm := cpi.PatchModel()

	_po, err := json.PatchObject(*e, *_pm)
	if err != nil {
		return err
	}

	_obj := _po.(*Connector)

	if e.Name != _obj.Name {
		return errors.New("field name is immutable")
	}
	if !reflect.DeepEqual(e.CreateTime, _obj.CreateTime) {
		return errors.New("field createTime is immutable")
	}
	if e.Category != _obj.Category {
		return errors.New("field category is immutable")
	}
	if e.Type != _obj.Type {
		return errors.New("field type is immutable")
	}
	if e.ApplicableEnvironmentType != _obj.ApplicableEnvironmentType {
		return errors.New("field applicableEnvironmentType is immutable")
	}

	cpi.patchedEntity = _obj
	return nil
}

// ConnectorQueryInput holds the query input of the Connector entity,
// please tags with `path:",inline"` if embedding.
type ConnectorQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query Connector entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"project,omitempty"`

	// Refer holds the route path reference of the Connector entity.
	Refer *object.Refer `path:"connector,default=" query:"-" json:"-"`
	// ID of the Connector entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Connector entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// Model returns the Connector entity for querying,
// after validating.
func (cqi *ConnectorQueryInput) Model() *Connector {
	if cqi == nil {
		return nil
	}

	return &Connector{
		ID:   cqi.ID,
		Name: cqi.Name,
	}
}

// Validate checks the ConnectorQueryInput entity.
func (cqi *ConnectorQueryInput) Validate() error {
	if cqi == nil {
		return errors.New("nil receiver")
	}

	return cqi.ValidateWith(cqi.inputConfig.Context, cqi.inputConfig.Client, nil)
}

// ValidateWith checks the ConnectorQueryInput entity with the given context and client set.
func (cqi *ConnectorQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cqi == nil {
		return errors.New("nil receiver")
	}

	if cqi.Refer != nil && *cqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", connector.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Connectors().Query()

	// Validate when querying under the Project route.
	if cqi.Project != nil {
		if err := cqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				cqi.Project = nil
				q.Where(
					connector.ProjectIDIsNil())
			}
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				connector.ProjectID(cqi.Project.ID))
		}
	} else {
		q.Where(
			connector.ProjectIDIsNil())
	}

	if cqi.Refer != nil {
		if cqi.Refer.IsID() {
			q.Where(
				connector.ID(cqi.Refer.ID()))
		} else if refers := cqi.Refer.Split(1); len(refers) == 1 {
			q.Where(
				connector.Name(refers[0].String()))
		} else {
			return errors.New("invalid identify refer of connector")
		}
	} else if cqi.ID != "" {
		q.Where(
			connector.ID(cqi.ID))
	} else if cqi.Name != "" {
		q.Where(
			connector.Name(cqi.Name))
	} else {
		return errors.New("invalid identify of connector")
	}

	q.Select(
		connector.FieldID,
		connector.FieldName,
	)

	var e *Connector
	{
		// Get cache from previous validation.
		queryStmt, queryArgs := q.sqlQuery(setContextOp(ctx, q.ctx, "cache")).Query()
		ck := fmt.Sprintf("stmt=%v, args=%v", queryStmt, queryArgs)
		if cv, existed := cache[ck]; !existed {
			var err error
			e, err = q.Only(ctx)
			if err != nil {
				return err
			}

			// Set cache for other validation.
			cache[ck] = e
		} else {
			e = cv.(*Connector)
		}
	}

	cqi.ID = e.ID
	cqi.Name = e.Name
	return nil
}

// ConnectorQueryInputs holds the query input of the Connector entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type ConnectorQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query Connector entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Category of the connector.
	Category string `path:"-" query:"category,omitempty" json:"-"`
	// Type of the connector.
	Type string `path:"-" query:"type,omitempty" json:"-"`
	// Environment type of the connector to apply.
	ApplicableEnvironmentType string `path:"-" query:"applicableEnvironmentType,omitempty" json:"-"`
}

// Validate checks the ConnectorQueryInputs entity.
func (cqi *ConnectorQueryInputs) Validate() error {
	if cqi == nil {
		return errors.New("nil receiver")
	}

	return cqi.ValidateWith(cqi.inputConfig.Context, cqi.inputConfig.Client, nil)
}

// ValidateWith checks the ConnectorQueryInputs entity with the given context and client set.
func (cqi *ConnectorQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when querying under the Project route.
	if cqi.Project != nil {
		if err := cqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				cqi.Project = nil
			}
		}
	}

	return nil
}

// ConnectorUpdateInput holds the modification input of the Connector entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ConnectorUpdateInput struct {
	ConnectorQueryInput `path:",inline" query:"-" json:"-"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Category of the connector.
	Category string `path:"-" query:"-" json:"category,omitempty"`
	// Type of the connector.
	Type string `path:"-" query:"-" json:"type,omitempty"`
	// Connector config version.
	ConfigVersion string `path:"-" query:"-" json:"configVersion,omitempty"`
	// Connector config data.
	ConfigData crypto.Properties `path:"-" query:"-" json:"configData,omitempty"`
	// Config whether enable finOps, will install prometheus and opencost while enable.
	EnableFinOps bool `path:"-" query:"-" json:"enableFinOps,omitempty"`
	// Custom pricing user defined.
	FinOpsCustomPricing *types.FinOpsCustomPricing `path:"-" query:"-" json:"finOpsCustomPricing,omitempty"`
}

// Model returns the Connector entity for modifying,
// after validating.
func (cui *ConnectorUpdateInput) Model() *Connector {
	if cui == nil {
		return nil
	}

	_c := &Connector{
		ID:                  cui.ID,
		Name:                cui.Name,
		Description:         cui.Description,
		Labels:              cui.Labels,
		Category:            cui.Category,
		Type:                cui.Type,
		ConfigVersion:       cui.ConfigVersion,
		ConfigData:          cui.ConfigData,
		EnableFinOps:        cui.EnableFinOps,
		FinOpsCustomPricing: cui.FinOpsCustomPricing,
	}

	return _c
}

// Validate checks the ConnectorUpdateInput entity.
func (cui *ConnectorUpdateInput) Validate() error {
	if cui == nil {
		return errors.New("nil receiver")
	}

	return cui.ValidateWith(cui.inputConfig.Context, cui.inputConfig.Client, nil)
}

// ValidateWith checks the ConnectorUpdateInput entity with the given context and client set.
func (cui *ConnectorUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := cui.ConnectorQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	return nil
}

// ConnectorUpdateInputs holds the modification input item of the Connector entities.
type ConnectorUpdateInputsItem struct {
	// ID of the Connector entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Connector entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Category of the connector.
	Category string `path:"-" query:"-" json:"category,cli-table-column"`
	// Type of the connector.
	Type string `path:"-" query:"-" json:"type,cli-table-column"`
	// Connector config version.
	ConfigVersion string `path:"-" query:"-" json:"configVersion"`
	// Connector config data.
	ConfigData crypto.Properties `path:"-" query:"-" json:"configData,omitempty"`
	// Config whether enable finOps, will install prometheus and opencost while enable.
	EnableFinOps bool `path:"-" query:"-" json:"enableFinOps"`
	// Custom pricing user defined.
	FinOpsCustomPricing *types.FinOpsCustomPricing `path:"-" query:"-" json:"finOpsCustomPricing,omitempty"`
}

// ValidateWith checks the ConnectorUpdateInputsItem entity with the given context and client set.
func (cui *ConnectorUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// ConnectorUpdateInputs holds the modification input of the Connector entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ConnectorUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to update Connector entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ConnectorUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Connector entities for modifying,
// after validating.
func (cui *ConnectorUpdateInputs) Model() []*Connector {
	if cui == nil || len(cui.Items) == 0 {
		return nil
	}

	_cs := make([]*Connector, len(cui.Items))

	for i := range cui.Items {
		_c := &Connector{
			ID:                  cui.Items[i].ID,
			Name:                cui.Items[i].Name,
			Description:         cui.Items[i].Description,
			Labels:              cui.Items[i].Labels,
			Category:            cui.Items[i].Category,
			Type:                cui.Items[i].Type,
			ConfigVersion:       cui.Items[i].ConfigVersion,
			ConfigData:          cui.Items[i].ConfigData,
			EnableFinOps:        cui.Items[i].EnableFinOps,
			FinOpsCustomPricing: cui.Items[i].FinOpsCustomPricing,
		}

		_cs[i] = _c
	}

	return _cs
}

// IDs returns the ID list of the Connector entities for modifying,
// after validating.
func (cui *ConnectorUpdateInputs) IDs() []object.ID {
	if cui == nil || len(cui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(cui.Items))
	for i := range cui.Items {
		ids[i] = cui.Items[i].ID
	}
	return ids
}

// Validate checks the ConnectorUpdateInputs entity.
func (cui *ConnectorUpdateInputs) Validate() error {
	if cui == nil {
		return errors.New("nil receiver")
	}

	return cui.ValidateWith(cui.inputConfig.Context, cui.inputConfig.Client, nil)
}

// ValidateWith checks the ConnectorUpdateInputs entity with the given context and client set.
func (cui *ConnectorUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cui == nil {
		return errors.New("nil receiver")
	}

	if len(cui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Connectors().Query()

	// Validate when updating under the Project route.
	if cui.Project != nil {
		if err := cui.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				cui.Project = nil
				q.Where(
					connector.ProjectIDIsNil())
			}
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				connector.ProjectID(cui.Project.ID))
		}
	} else {
		q.Where(
			connector.ProjectIDIsNil())
	}

	ids := make([]object.ID, 0, len(cui.Items))
	ors := make([]predicate.Connector, 0, len(cui.Items))
	indexers := make(map[any][]int)

	for i := range cui.Items {
		if cui.Items[i] == nil {
			return errors.New("nil item")
		}

		if cui.Items[i].ID != "" {
			ids = append(ids, cui.Items[i].ID)
			ors = append(ors, connector.ID(cui.Items[i].ID))
			indexers[cui.Items[i].ID] = append(indexers[cui.Items[i].ID], i)
		} else if cui.Items[i].Name != "" {
			ors = append(ors, connector.And(
				connector.Name(cui.Items[i].Name)))
			indexerKey := fmt.Sprint("/", cui.Items[i].Name)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := connector.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = connector.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			connector.FieldID,
			connector.FieldName,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		indexer := indexers[es[i].ID]
		if indexer == nil {
			indexerKey := fmt.Sprint("/", es[i].Name)
			indexer = indexers[indexerKey]
		}
		for _, j := range indexer {
			cui.Items[j].ID = es[i].ID
			cui.Items[j].Name = es[i].Name
		}
	}

	for i := range cui.Items {
		if err := cui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ConnectorOutput holds the output of the Connector entity.
type ConnectorOutput struct {
	ID                        object.ID                  `json:"id,omitempty"`
	Name                      string                     `json:"name,omitempty"`
	Description               string                     `json:"description,omitempty"`
	Labels                    map[string]string          `json:"labels,omitempty"`
	CreateTime                *time.Time                 `json:"createTime,omitempty"`
	UpdateTime                *time.Time                 `json:"updateTime,omitempty"`
	Status                    status.Status              `json:"status,omitempty"`
	Category                  string                     `json:"category,cli-table-column,omitempty"`
	Type                      string                     `json:"type,cli-table-column,omitempty"`
	ApplicableEnvironmentType string                     `json:"applicableEnvironmentType,omitempty"`
	ConfigVersion             string                     `json:"configVersion,omitempty"`
	ConfigData                crypto.Properties          `json:"configData,omitempty"`
	EnableFinOps              bool                       `json:"enableFinOps,omitempty"`
	FinOpsCustomPricing       *types.FinOpsCustomPricing `json:"finOpsCustomPricing,omitempty"`

	Project *ProjectOutput `json:"project,omitempty"`
}

// View returns the output of Connector entity.
func (_c *Connector) View() *ConnectorOutput {
	return ExposeConnector(_c)
}

// View returns the output of Connector entities.
func (_cs Connectors) View() []*ConnectorOutput {
	return ExposeConnectors(_cs)
}

// ExposeConnector converts the Connector to ConnectorOutput.
func ExposeConnector(_c *Connector) *ConnectorOutput {
	if _c == nil {
		return nil
	}

	co := &ConnectorOutput{
		ID:                        _c.ID,
		Name:                      _c.Name,
		Description:               _c.Description,
		Labels:                    _c.Labels,
		CreateTime:                _c.CreateTime,
		UpdateTime:                _c.UpdateTime,
		Status:                    _c.Status,
		Category:                  _c.Category,
		Type:                      _c.Type,
		ApplicableEnvironmentType: _c.ApplicableEnvironmentType,
		ConfigVersion:             _c.ConfigVersion,
		ConfigData:                _c.ConfigData,
		EnableFinOps:              _c.EnableFinOps,
		FinOpsCustomPricing:       _c.FinOpsCustomPricing,
	}

	if _c.Edges.Project != nil {
		co.Project = ExposeProject(_c.Edges.Project)
	} else if _c.ProjectID != "" {
		co.Project = &ProjectOutput{
			ID: _c.ProjectID,
		}
	}
	return co
}

// ExposeConnectors converts the Connector slice to ConnectorOutput pointer slice.
func ExposeConnectors(_cs []*Connector) []*ConnectorOutput {
	if len(_cs) == 0 {
		return nil
	}

	cos := make([]*ConnectorOutput, len(_cs))
	for i := range _cs {
		cos[i] = ExposeConnector(_cs[i])
	}
	return cos
}
