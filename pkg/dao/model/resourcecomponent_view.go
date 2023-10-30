// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/walrus/pkg/dao/model/resourcecomponent"
	"github.com/seal-io/walrus/pkg/dao/schema/intercept"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// ResourceComponentCreateInput holds the creation input of the ResourceComponent entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceComponentCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create ResourceComponent entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to create ResourceComponent entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Resource indicates to create ResourceComponent entity MUST under the Resource route.
	Resource *ResourceQueryInput `path:",inline" query:"-" json:"-"`

	// Shape of the component, it can be class or instance shape.
	Shape string `path:"-" query:"-" json:"shape"`
	// Type of deployer.
	DeployerType string `path:"-" query:"-" json:"deployerType"`
	// Name of the generated component, it is the real identifier of the component, which provides by deployer.
	Name string `path:"-" query:"-" json:"name"`
	// Type of the generated component, it is the type of the resource which the deployer observes, which provides by deployer.
	Type string `path:"-" query:"-" json:"type"`
	// Mode that manages the generated component, it is the management way of the deployer to the component, which provides by deployer.
	Mode string `path:"-" query:"-" json:"mode"`
	// Status of the component.
	Status types.ServiceResourceStatus `path:"-" query:"-" json:"status,omitempty"`

	// Components specifies full inserting the new ResourceComponent entities of the ResourceComponent entity.
	Components []*ResourceComponentCreateInput `uri:"-" query:"-" json:"components,omitempty"`
	// Instances specifies full inserting the new ResourceComponent entities of the ResourceComponent entity.
	Instances []*ResourceComponentCreateInput `uri:"-" query:"-" json:"instances,omitempty"`
	// Dependencies specifies full inserting the new ResourceComponentRelationship entities of the ResourceComponent entity.
	Dependencies []*ResourceComponentRelationshipCreateInput `uri:"-" query:"-" json:"dependencies,omitempty"`
}

// Model returns the ResourceComponent entity for creating,
// after validating.
func (rcci *ResourceComponentCreateInput) Model() *ResourceComponent {
	if rcci == nil {
		return nil
	}

	_rc := &ResourceComponent{
		Shape:        rcci.Shape,
		DeployerType: rcci.DeployerType,
		Name:         rcci.Name,
		Type:         rcci.Type,
		Mode:         rcci.Mode,
		Status:       rcci.Status,
	}

	if rcci.Project != nil {
		_rc.ProjectID = rcci.Project.ID
	}
	if rcci.Environment != nil {
		_rc.EnvironmentID = rcci.Environment.ID
	}
	if rcci.Resource != nil {
		_rc.ResourceID = rcci.Resource.ID
	}

	if rcci.Components != nil {
		// Empty slice is used for clearing the edge.
		_rc.Edges.Components = make([]*ResourceComponent, 0, len(rcci.Components))
	}
	for j := range rcci.Components {
		if rcci.Components[j] == nil {
			continue
		}
		_rc.Edges.Components = append(_rc.Edges.Components,
			rcci.Components[j].Model())
	}
	if rcci.Instances != nil {
		// Empty slice is used for clearing the edge.
		_rc.Edges.Instances = make([]*ResourceComponent, 0, len(rcci.Instances))
	}
	for j := range rcci.Instances {
		if rcci.Instances[j] == nil {
			continue
		}
		_rc.Edges.Instances = append(_rc.Edges.Instances,
			rcci.Instances[j].Model())
	}
	if rcci.Dependencies != nil {
		// Empty slice is used for clearing the edge.
		_rc.Edges.Dependencies = make([]*ResourceComponentRelationship, 0, len(rcci.Dependencies))
	}
	for j := range rcci.Dependencies {
		if rcci.Dependencies[j] == nil {
			continue
		}
		_rc.Edges.Dependencies = append(_rc.Edges.Dependencies,
			rcci.Dependencies[j].Model())
	}
	return _rc
}

// Validate checks the ResourceComponentCreateInput entity.
func (rcci *ResourceComponentCreateInput) Validate() error {
	if rcci == nil {
		return errors.New("nil receiver")
	}

	return rcci.ValidateWith(rcci.inputConfig.Context, rcci.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceComponentCreateInput entity with the given context and client set.
func (rcci *ResourceComponentCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if rcci.Project != nil {
		if err := rcci.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}
	// Validate when creating under the Environment route.
	if rcci.Environment != nil {
		if err := rcci.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}
	// Validate when creating under the Resource route.
	if rcci.Resource != nil {
		if err := rcci.Resource.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	for i := range rcci.Components {
		if rcci.Components[i] == nil {
			continue
		}

		if err := rcci.Components[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcci.Components[i] = nil
			}
		}
	}

	for i := range rcci.Instances {
		if rcci.Instances[i] == nil {
			continue
		}

		if err := rcci.Instances[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcci.Instances[i] = nil
			}
		}
	}

	for i := range rcci.Dependencies {
		if rcci.Dependencies[i] == nil {
			continue
		}

		if err := rcci.Dependencies[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcci.Dependencies[i] = nil
			}
		}
	}

	return nil
}

// ResourceComponentCreateInputs holds the creation input item of the ResourceComponent entities.
type ResourceComponentCreateInputsItem struct {
	// Shape of the component, it can be class or instance shape.
	Shape string `path:"-" query:"-" json:"shape"`
	// Type of deployer.
	DeployerType string `path:"-" query:"-" json:"deployerType"`
	// Name of the generated component, it is the real identifier of the component, which provides by deployer.
	Name string `path:"-" query:"-" json:"name"`
	// Type of the generated component, it is the type of the resource which the deployer observes, which provides by deployer.
	Type string `path:"-" query:"-" json:"type"`
	// Mode that manages the generated component, it is the management way of the deployer to the component, which provides by deployer.
	Mode string `path:"-" query:"-" json:"mode"`
	// Status of the component.
	Status types.ServiceResourceStatus `path:"-" query:"-" json:"status,omitempty"`

	// Components specifies full inserting the new ResourceComponent entities.
	Components []*ResourceComponentCreateInput `uri:"-" query:"-" json:"components,omitempty"`
	// Instances specifies full inserting the new ResourceComponent entities.
	Instances []*ResourceComponentCreateInput `uri:"-" query:"-" json:"instances,omitempty"`
	// Dependencies specifies full inserting the new ResourceComponentRelationship entities.
	Dependencies []*ResourceComponentRelationshipCreateInput `uri:"-" query:"-" json:"dependencies,omitempty"`
}

// ValidateWith checks the ResourceComponentCreateInputsItem entity with the given context and client set.
func (rcci *ResourceComponentCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range rcci.Components {
		if rcci.Components[i] == nil {
			continue
		}

		if err := rcci.Components[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcci.Components[i] = nil
			}
		}
	}

	for i := range rcci.Instances {
		if rcci.Instances[i] == nil {
			continue
		}

		if err := rcci.Instances[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcci.Instances[i] = nil
			}
		}
	}

	for i := range rcci.Dependencies {
		if rcci.Dependencies[i] == nil {
			continue
		}

		if err := rcci.Dependencies[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcci.Dependencies[i] = nil
			}
		}
	}

	return nil
}

// ResourceComponentCreateInputs holds the creation input of the ResourceComponent entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceComponentCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create ResourceComponent entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to create ResourceComponent entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Resource indicates to create ResourceComponent entity MUST under the Resource route.
	Resource *ResourceQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ResourceComponentCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ResourceComponent entities for creating,
// after validating.
func (rcci *ResourceComponentCreateInputs) Model() []*ResourceComponent {
	if rcci == nil || len(rcci.Items) == 0 {
		return nil
	}

	_rcs := make([]*ResourceComponent, len(rcci.Items))

	for i := range rcci.Items {
		_rc := &ResourceComponent{
			Shape:        rcci.Items[i].Shape,
			DeployerType: rcci.Items[i].DeployerType,
			Name:         rcci.Items[i].Name,
			Type:         rcci.Items[i].Type,
			Mode:         rcci.Items[i].Mode,
			Status:       rcci.Items[i].Status,
		}

		if rcci.Project != nil {
			_rc.ProjectID = rcci.Project.ID
		}
		if rcci.Environment != nil {
			_rc.EnvironmentID = rcci.Environment.ID
		}
		if rcci.Resource != nil {
			_rc.ResourceID = rcci.Resource.ID
		}

		if rcci.Items[i].Components != nil {
			// Empty slice is used for clearing the edge.
			_rc.Edges.Components = make([]*ResourceComponent, 0, len(rcci.Items[i].Components))
		}
		for j := range rcci.Items[i].Components {
			if rcci.Items[i].Components[j] == nil {
				continue
			}
			_rc.Edges.Components = append(_rc.Edges.Components,
				rcci.Items[i].Components[j].Model())
		}
		if rcci.Items[i].Instances != nil {
			// Empty slice is used for clearing the edge.
			_rc.Edges.Instances = make([]*ResourceComponent, 0, len(rcci.Items[i].Instances))
		}
		for j := range rcci.Items[i].Instances {
			if rcci.Items[i].Instances[j] == nil {
				continue
			}
			_rc.Edges.Instances = append(_rc.Edges.Instances,
				rcci.Items[i].Instances[j].Model())
		}
		if rcci.Items[i].Dependencies != nil {
			// Empty slice is used for clearing the edge.
			_rc.Edges.Dependencies = make([]*ResourceComponentRelationship, 0, len(rcci.Items[i].Dependencies))
		}
		for j := range rcci.Items[i].Dependencies {
			if rcci.Items[i].Dependencies[j] == nil {
				continue
			}
			_rc.Edges.Dependencies = append(_rc.Edges.Dependencies,
				rcci.Items[i].Dependencies[j].Model())
		}

		_rcs[i] = _rc
	}

	return _rcs
}

// Validate checks the ResourceComponentCreateInputs entity .
func (rcci *ResourceComponentCreateInputs) Validate() error {
	if rcci == nil {
		return errors.New("nil receiver")
	}

	return rcci.ValidateWith(rcci.inputConfig.Context, rcci.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceComponentCreateInputs entity with the given context and client set.
func (rcci *ResourceComponentCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcci == nil {
		return errors.New("nil receiver")
	}

	if len(rcci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if rcci.Project != nil {
		if err := rcci.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcci.Project = nil
			}
		}
	}
	// Validate when creating under the Environment route.
	if rcci.Environment != nil {
		if err := rcci.Environment.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcci.Environment = nil
			}
		}
	}
	// Validate when creating under the Resource route.
	if rcci.Resource != nil {
		if err := rcci.Resource.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcci.Resource = nil
			}
		}
	}

	for i := range rcci.Items {
		if rcci.Items[i] == nil {
			continue
		}

		if err := rcci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ResourceComponentDeleteInput holds the deletion input of the ResourceComponent entity,
// please tags with `path:",inline"` if embedding.
type ResourceComponentDeleteInput struct {
	ResourceComponentQueryInput `path:",inline"`
}

// ResourceComponentDeleteInputs holds the deletion input item of the ResourceComponent entities.
type ResourceComponentDeleteInputsItem struct {
	// ID of the ResourceComponent entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// ResourceComponentDeleteInputs holds the deletion input of the ResourceComponent entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceComponentDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to delete ResourceComponent entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to delete ResourceComponent entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Resource indicates to delete ResourceComponent entity MUST under the Resource route.
	Resource *ResourceQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ResourceComponentDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ResourceComponent entities for deleting,
// after validating.
func (rcdi *ResourceComponentDeleteInputs) Model() []*ResourceComponent {
	if rcdi == nil || len(rcdi.Items) == 0 {
		return nil
	}

	_rcs := make([]*ResourceComponent, len(rcdi.Items))
	for i := range rcdi.Items {
		_rcs[i] = &ResourceComponent{
			ID: rcdi.Items[i].ID,
		}
	}
	return _rcs
}

// IDs returns the ID list of the ResourceComponent entities for deleting,
// after validating.
func (rcdi *ResourceComponentDeleteInputs) IDs() []object.ID {
	if rcdi == nil || len(rcdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(rcdi.Items))
	for i := range rcdi.Items {
		ids[i] = rcdi.Items[i].ID
	}
	return ids
}

// Validate checks the ResourceComponentDeleteInputs entity.
func (rcdi *ResourceComponentDeleteInputs) Validate() error {
	if rcdi == nil {
		return errors.New("nil receiver")
	}

	return rcdi.ValidateWith(rcdi.inputConfig.Context, rcdi.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceComponentDeleteInputs entity with the given context and client set.
func (rcdi *ResourceComponentDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcdi == nil {
		return errors.New("nil receiver")
	}

	if len(rcdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ResourceComponents().Query()

	// Validate when deleting under the Project route.
	if rcdi.Project != nil {
		if err := rcdi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				resourcecomponent.ProjectID(rcdi.Project.ID))
		}
	}

	// Validate when deleting under the Environment route.
	if rcdi.Environment != nil {
		if err := rcdi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				resourcecomponent.EnvironmentID(rcdi.Environment.ID))
		}
	}

	// Validate when deleting under the Resource route.
	if rcdi.Resource != nil {
		if err := rcdi.Resource.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				resourcecomponent.ResourceID(rcdi.Resource.ID))
		}
	}

	ids := make([]object.ID, 0, len(rcdi.Items))

	for i := range rcdi.Items {
		if rcdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if rcdi.Items[i].ID != "" {
			ids = append(ids, rcdi.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(resourcecomponent.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	return nil
}

// ResourceComponentQueryInput holds the query input of the ResourceComponent entity,
// please tags with `path:",inline"` if embedding.
type ResourceComponentQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query ResourceComponent entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"project"`
	// Environment indicates to query ResourceComponent entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"environment"`
	// Resource indicates to query ResourceComponent entity MUST under the Resource route.
	Resource *ResourceQueryInput `path:",inline" query:"-" json:"resource"`

	// Refer holds the route path reference of the ResourceComponent entity.
	Refer *object.Refer `path:"resourcecomponent,default=" query:"-" json:"-"`
	// ID of the ResourceComponent entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// Model returns the ResourceComponent entity for querying,
// after validating.
func (rcqi *ResourceComponentQueryInput) Model() *ResourceComponent {
	if rcqi == nil {
		return nil
	}

	return &ResourceComponent{
		ID: rcqi.ID,
	}
}

// Validate checks the ResourceComponentQueryInput entity.
func (rcqi *ResourceComponentQueryInput) Validate() error {
	if rcqi == nil {
		return errors.New("nil receiver")
	}

	return rcqi.ValidateWith(rcqi.inputConfig.Context, rcqi.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceComponentQueryInput entity with the given context and client set.
func (rcqi *ResourceComponentQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcqi == nil {
		return errors.New("nil receiver")
	}

	if rcqi.Refer != nil && *rcqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", resourcecomponent.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ResourceComponents().Query()

	// Validate when querying under the Project route.
	if rcqi.Project != nil {
		if err := rcqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				resourcecomponent.ProjectID(rcqi.Project.ID))
		}
	}

	// Validate when querying under the Environment route.
	if rcqi.Environment != nil {
		if err := rcqi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				resourcecomponent.EnvironmentID(rcqi.Environment.ID))
		}
	}

	// Validate when querying under the Resource route.
	if rcqi.Resource != nil {
		if err := rcqi.Resource.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				resourcecomponent.ResourceID(rcqi.Resource.ID))
		}
	}

	if rcqi.Refer != nil {
		if rcqi.Refer.IsID() {
			q.Where(
				resourcecomponent.ID(rcqi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of resourcecomponent")
		}
	} else if rcqi.ID != "" {
		q.Where(
			resourcecomponent.ID(rcqi.ID))
	} else {
		return errors.New("invalid identify of resourcecomponent")
	}

	q.Select(
		resourcecomponent.FieldID,
	)

	var e *ResourceComponent
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
			e = cv.(*ResourceComponent)
		}
	}

	rcqi.ID = e.ID
	return nil
}

// ResourceComponentQueryInputs holds the query input of the ResourceComponent entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type ResourceComponentQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query ResourceComponent entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to query ResourceComponent entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Resource indicates to query ResourceComponent entity MUST under the Resource route.
	Resource *ResourceQueryInput `path:",inline" query:"-" json:"-"`
}

// Validate checks the ResourceComponentQueryInputs entity.
func (rcqi *ResourceComponentQueryInputs) Validate() error {
	if rcqi == nil {
		return errors.New("nil receiver")
	}

	return rcqi.ValidateWith(rcqi.inputConfig.Context, rcqi.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceComponentQueryInputs entity with the given context and client set.
func (rcqi *ResourceComponentQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when querying under the Project route.
	if rcqi.Project != nil {
		if err := rcqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	// Validate when querying under the Environment route.
	if rcqi.Environment != nil {
		if err := rcqi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	// Validate when querying under the Resource route.
	if rcqi.Resource != nil {
		if err := rcqi.Resource.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ResourceComponentUpdateInput holds the modification input of the ResourceComponent entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceComponentUpdateInput struct {
	ResourceComponentQueryInput `path:",inline" query:"-" json:"-"`

	// Status of the component.
	Status types.ServiceResourceStatus `path:"-" query:"-" json:"status,omitempty"`

	// Components indicates replacing the stale ResourceComponent entities.
	Components []*ResourceComponentCreateInput `uri:"-" query:"-" json:"components,omitempty"`
	// Instances indicates replacing the stale ResourceComponent entities.
	Instances []*ResourceComponentCreateInput `uri:"-" query:"-" json:"instances,omitempty"`
	// Dependencies indicates replacing the stale ResourceComponentRelationship entities.
	Dependencies []*ResourceComponentRelationshipCreateInput `uri:"-" query:"-" json:"dependencies,omitempty"`
}

// Model returns the ResourceComponent entity for modifying,
// after validating.
func (rcui *ResourceComponentUpdateInput) Model() *ResourceComponent {
	if rcui == nil {
		return nil
	}

	_rc := &ResourceComponent{
		ID:     rcui.ID,
		Status: rcui.Status,
	}

	if rcui.Components != nil {
		// Empty slice is used for clearing the edge.
		_rc.Edges.Components = make([]*ResourceComponent, 0, len(rcui.Components))
	}
	for j := range rcui.Components {
		if rcui.Components[j] == nil {
			continue
		}
		_rc.Edges.Components = append(_rc.Edges.Components,
			rcui.Components[j].Model())
	}
	if rcui.Instances != nil {
		// Empty slice is used for clearing the edge.
		_rc.Edges.Instances = make([]*ResourceComponent, 0, len(rcui.Instances))
	}
	for j := range rcui.Instances {
		if rcui.Instances[j] == nil {
			continue
		}
		_rc.Edges.Instances = append(_rc.Edges.Instances,
			rcui.Instances[j].Model())
	}
	if rcui.Dependencies != nil {
		// Empty slice is used for clearing the edge.
		_rc.Edges.Dependencies = make([]*ResourceComponentRelationship, 0, len(rcui.Dependencies))
	}
	for j := range rcui.Dependencies {
		if rcui.Dependencies[j] == nil {
			continue
		}
		_rc.Edges.Dependencies = append(_rc.Edges.Dependencies,
			rcui.Dependencies[j].Model())
	}
	return _rc
}

// Validate checks the ResourceComponentUpdateInput entity.
func (rcui *ResourceComponentUpdateInput) Validate() error {
	if rcui == nil {
		return errors.New("nil receiver")
	}

	return rcui.ValidateWith(rcui.inputConfig.Context, rcui.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceComponentUpdateInput entity with the given context and client set.
func (rcui *ResourceComponentUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := rcui.ResourceComponentQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	for i := range rcui.Components {
		if rcui.Components[i] == nil {
			continue
		}

		if err := rcui.Components[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcui.Components[i] = nil
			}
		}
	}

	for i := range rcui.Instances {
		if rcui.Instances[i] == nil {
			continue
		}

		if err := rcui.Instances[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcui.Instances[i] = nil
			}
		}
	}

	for i := range rcui.Dependencies {
		if rcui.Dependencies[i] == nil {
			continue
		}

		if err := rcui.Dependencies[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcui.Dependencies[i] = nil
			}
		}
	}

	return nil
}

// ResourceComponentUpdateInputs holds the modification input item of the ResourceComponent entities.
type ResourceComponentUpdateInputsItem struct {
	// ID of the ResourceComponent entity.
	ID object.ID `path:"-" query:"-" json:"id"`

	// Status of the component.
	Status types.ServiceResourceStatus `path:"-" query:"-" json:"status,omitempty"`

	// Components indicates replacing the stale ResourceComponent entities.
	Components []*ResourceComponentCreateInput `uri:"-" query:"-" json:"components,omitempty"`
	// Instances indicates replacing the stale ResourceComponent entities.
	Instances []*ResourceComponentCreateInput `uri:"-" query:"-" json:"instances,omitempty"`
	// Dependencies indicates replacing the stale ResourceComponentRelationship entities.
	Dependencies []*ResourceComponentRelationshipCreateInput `uri:"-" query:"-" json:"dependencies,omitempty"`
}

// ValidateWith checks the ResourceComponentUpdateInputsItem entity with the given context and client set.
func (rcui *ResourceComponentUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range rcui.Components {
		if rcui.Components[i] == nil {
			continue
		}

		if err := rcui.Components[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcui.Components[i] = nil
			}
		}
	}

	for i := range rcui.Instances {
		if rcui.Instances[i] == nil {
			continue
		}

		if err := rcui.Instances[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcui.Instances[i] = nil
			}
		}
	}

	for i := range rcui.Dependencies {
		if rcui.Dependencies[i] == nil {
			continue
		}

		if err := rcui.Dependencies[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcui.Dependencies[i] = nil
			}
		}
	}

	return nil
}

// ResourceComponentUpdateInputs holds the modification input of the ResourceComponent entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceComponentUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to update ResourceComponent entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to update ResourceComponent entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Resource indicates to update ResourceComponent entity MUST under the Resource route.
	Resource *ResourceQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ResourceComponentUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ResourceComponent entities for modifying,
// after validating.
func (rcui *ResourceComponentUpdateInputs) Model() []*ResourceComponent {
	if rcui == nil || len(rcui.Items) == 0 {
		return nil
	}

	_rcs := make([]*ResourceComponent, len(rcui.Items))

	for i := range rcui.Items {
		_rc := &ResourceComponent{
			ID:     rcui.Items[i].ID,
			Status: rcui.Items[i].Status,
		}

		if rcui.Items[i].Components != nil {
			// Empty slice is used for clearing the edge.
			_rc.Edges.Components = make([]*ResourceComponent, 0, len(rcui.Items[i].Components))
		}
		for j := range rcui.Items[i].Components {
			if rcui.Items[i].Components[j] == nil {
				continue
			}
			_rc.Edges.Components = append(_rc.Edges.Components,
				rcui.Items[i].Components[j].Model())
		}
		if rcui.Items[i].Instances != nil {
			// Empty slice is used for clearing the edge.
			_rc.Edges.Instances = make([]*ResourceComponent, 0, len(rcui.Items[i].Instances))
		}
		for j := range rcui.Items[i].Instances {
			if rcui.Items[i].Instances[j] == nil {
				continue
			}
			_rc.Edges.Instances = append(_rc.Edges.Instances,
				rcui.Items[i].Instances[j].Model())
		}
		if rcui.Items[i].Dependencies != nil {
			// Empty slice is used for clearing the edge.
			_rc.Edges.Dependencies = make([]*ResourceComponentRelationship, 0, len(rcui.Items[i].Dependencies))
		}
		for j := range rcui.Items[i].Dependencies {
			if rcui.Items[i].Dependencies[j] == nil {
				continue
			}
			_rc.Edges.Dependencies = append(_rc.Edges.Dependencies,
				rcui.Items[i].Dependencies[j].Model())
		}

		_rcs[i] = _rc
	}

	return _rcs
}

// IDs returns the ID list of the ResourceComponent entities for modifying,
// after validating.
func (rcui *ResourceComponentUpdateInputs) IDs() []object.ID {
	if rcui == nil || len(rcui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(rcui.Items))
	for i := range rcui.Items {
		ids[i] = rcui.Items[i].ID
	}
	return ids
}

// Validate checks the ResourceComponentUpdateInputs entity.
func (rcui *ResourceComponentUpdateInputs) Validate() error {
	if rcui == nil {
		return errors.New("nil receiver")
	}

	return rcui.ValidateWith(rcui.inputConfig.Context, rcui.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceComponentUpdateInputs entity with the given context and client set.
func (rcui *ResourceComponentUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcui == nil {
		return errors.New("nil receiver")
	}

	if len(rcui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ResourceComponents().Query()

	// Validate when updating under the Project route.
	if rcui.Project != nil {
		if err := rcui.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				resourcecomponent.ProjectID(rcui.Project.ID))
		}
	}

	// Validate when updating under the Environment route.
	if rcui.Environment != nil {
		if err := rcui.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				resourcecomponent.EnvironmentID(rcui.Environment.ID))
		}
	}

	// Validate when updating under the Resource route.
	if rcui.Resource != nil {
		if err := rcui.Resource.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				resourcecomponent.ResourceID(rcui.Resource.ID))
		}
	}

	ids := make([]object.ID, 0, len(rcui.Items))

	for i := range rcui.Items {
		if rcui.Items[i] == nil {
			return errors.New("nil item")
		}

		if rcui.Items[i].ID != "" {
			ids = append(ids, rcui.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(resourcecomponent.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range rcui.Items {
		if err := rcui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ResourceComponentOutput holds the output of the ResourceComponent entity.
type ResourceComponentOutput struct {
	ID           object.ID                           `json:"id,omitempty"`
	CreateTime   *time.Time                          `json:"createTime,omitempty"`
	UpdateTime   *time.Time                          `json:"updateTime,omitempty"`
	Mode         string                              `json:"mode,omitempty"`
	Type         string                              `json:"type,omitempty"`
	Name         string                              `json:"name,omitempty"`
	DeployerType string                              `json:"deployerType,omitempty"`
	Shape        string                              `json:"shape,omitempty"`
	Status       types.ServiceResourceStatus         `json:"status,omitempty"`
	Keys         *types.ServiceResourceOperationKeys `json:"keys,omitempty"`

	Project      *ProjectOutput                         `json:"project,omitempty"`
	Environment  *EnvironmentOutput                     `json:"environment,omitempty"`
	Resource     *ResourceOutput                        `json:"resource,omitempty"`
	Connector    *ConnectorOutput                       `json:"connector,omitempty"`
	Composition  *ResourceComponentOutput               `json:"composition,omitempty"`
	Components   []*ResourceComponentOutput             `json:"components,omitempty"`
	Class        *ResourceComponentOutput               `json:"class,omitempty"`
	Instances    []*ResourceComponentOutput             `json:"instances,omitempty"`
	Dependencies []*ResourceComponentRelationshipOutput `json:"dependencies,omitempty"`
}

// View returns the output of ResourceComponent entity.
func (_rc *ResourceComponent) View() *ResourceComponentOutput {
	return ExposeResourceComponent(_rc)
}

// View returns the output of ResourceComponent entities.
func (_rcs ResourceComponents) View() []*ResourceComponentOutput {
	return ExposeResourceComponents(_rcs)
}

// ExposeResourceComponent converts the ResourceComponent to ResourceComponentOutput.
func ExposeResourceComponent(_rc *ResourceComponent) *ResourceComponentOutput {
	if _rc == nil {
		return nil
	}

	rco := &ResourceComponentOutput{
		ID:           _rc.ID,
		CreateTime:   _rc.CreateTime,
		UpdateTime:   _rc.UpdateTime,
		Mode:         _rc.Mode,
		Type:         _rc.Type,
		Name:         _rc.Name,
		DeployerType: _rc.DeployerType,
		Shape:        _rc.Shape,
		Status:       _rc.Status,
		Keys:         _rc.Keys,
	}

	if _rc.Edges.Project != nil {
		rco.Project = ExposeProject(_rc.Edges.Project)
	} else if _rc.ProjectID != "" {
		rco.Project = &ProjectOutput{
			ID: _rc.ProjectID,
		}
	}
	if _rc.Edges.Environment != nil {
		rco.Environment = ExposeEnvironment(_rc.Edges.Environment)
	} else if _rc.EnvironmentID != "" {
		rco.Environment = &EnvironmentOutput{
			ID: _rc.EnvironmentID,
		}
	}
	if _rc.Edges.Resource != nil {
		rco.Resource = ExposeResource(_rc.Edges.Resource)
	} else if _rc.ResourceID != "" {
		rco.Resource = &ResourceOutput{
			ID: _rc.ResourceID,
		}
	}
	if _rc.Edges.Connector != nil {
		rco.Connector = ExposeConnector(_rc.Edges.Connector)
	} else if _rc.ConnectorID != "" {
		rco.Connector = &ConnectorOutput{
			ID: _rc.ConnectorID,
		}
	}
	if _rc.Edges.Composition != nil {
		rco.Composition = ExposeResourceComponent(_rc.Edges.Composition)
	} else if _rc.CompositionID != "" {
		rco.Composition = &ResourceComponentOutput{
			ID: _rc.CompositionID,
		}
	}
	if _rc.Edges.Components != nil {
		rco.Components = ExposeResourceComponents(_rc.Edges.Components)
	}
	if _rc.Edges.Class != nil {
		rco.Class = ExposeResourceComponent(_rc.Edges.Class)
	} else if _rc.ClassID != "" {
		rco.Class = &ResourceComponentOutput{
			ID: _rc.ClassID,
		}
	}
	if _rc.Edges.Instances != nil {
		rco.Instances = ExposeResourceComponents(_rc.Edges.Instances)
	}
	if _rc.Edges.Dependencies != nil {
		rco.Dependencies = ExposeResourceComponentRelationships(_rc.Edges.Dependencies)
	}
	return rco
}

// ExposeResourceComponents converts the ResourceComponent slice to ResourceComponentOutput pointer slice.
func ExposeResourceComponents(_rcs []*ResourceComponent) []*ResourceComponentOutput {
	if len(_rcs) == 0 {
		return nil
	}

	rcos := make([]*ResourceComponentOutput, len(_rcs))
	for i := range _rcs {
		rcos[i] = ExposeResourceComponent(_rcs[i])
	}
	return rcos
}
