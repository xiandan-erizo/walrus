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

	"github.com/seal-io/walrus/pkg/dao/model/workflowstageexecution"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/utils/json"
)

// WorkflowStageExecutionCreateInput holds the creation input of the WorkflowStageExecution entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStageExecutionCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// WorkflowExecution indicates to create WorkflowStageExecution entity MUST under the WorkflowExecution route.
	WorkflowExecution *WorkflowExecutionQueryInput `path:",inline" query:"-" json:"-"`

	// ID of the workflow stage that this workflow stage execution belongs to.
	WorkflowStageID object.ID `path:"-" query:"-" json:"workflowStageID"`
	// ID of the workflow that this workflow execution belongs to.
	WorkflowID object.ID `path:"-" query:"-" json:"workflowID"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`

	// Steps specifies full inserting the new WorkflowStepExecution entities of the WorkflowStageExecution entity.
	Steps []*WorkflowStepExecutionCreateInput `uri:"-" query:"-" json:"steps,omitempty"`
}

// Model returns the WorkflowStageExecution entity for creating,
// after validating.
func (wseci *WorkflowStageExecutionCreateInput) Model() *WorkflowStageExecution {
	if wseci == nil {
		return nil
	}

	_wse := &WorkflowStageExecution{
		WorkflowStageID: wseci.WorkflowStageID,
		WorkflowID:      wseci.WorkflowID,
		Name:            wseci.Name,
		Description:     wseci.Description,
		Labels:          wseci.Labels,
	}

	if wseci.WorkflowExecution != nil {
		_wse.WorkflowExecutionID = wseci.WorkflowExecution.ID
	}

	if wseci.Steps != nil {
		// Empty slice is used for clearing the edge.
		_wse.Edges.Steps = make([]*WorkflowStepExecution, 0, len(wseci.Steps))
	}
	for j := range wseci.Steps {
		if wseci.Steps[j] == nil {
			continue
		}
		_wse.Edges.Steps = append(_wse.Edges.Steps,
			wseci.Steps[j].Model())
	}
	return _wse
}

// Validate checks the WorkflowStageExecutionCreateInput entity.
func (wseci *WorkflowStageExecutionCreateInput) Validate() error {
	if wseci == nil {
		return errors.New("nil receiver")
	}

	return wseci.ValidateWith(wseci.inputConfig.Context, wseci.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStageExecutionCreateInput entity with the given context and client set.
func (wseci *WorkflowStageExecutionCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wseci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the WorkflowExecution route.
	if wseci.WorkflowExecution != nil {
		if err := wseci.WorkflowExecution.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	for i := range wseci.Steps {
		if wseci.Steps[i] == nil {
			continue
		}

		if err := wseci.Steps[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				wseci.Steps[i] = nil
			}
		}
	}

	return nil
}

// WorkflowStageExecutionCreateInputs holds the creation input item of the WorkflowStageExecution entities.
type WorkflowStageExecutionCreateInputsItem struct {
	// ID of the workflow stage that this workflow stage execution belongs to.
	WorkflowStageID object.ID `path:"-" query:"-" json:"workflowStageID"`
	// ID of the workflow that this workflow execution belongs to.
	WorkflowID object.ID `path:"-" query:"-" json:"workflowID"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`

	// Steps specifies full inserting the new WorkflowStepExecution entities.
	Steps []*WorkflowStepExecutionCreateInput `uri:"-" query:"-" json:"steps,omitempty"`
}

// ValidateWith checks the WorkflowStageExecutionCreateInputsItem entity with the given context and client set.
func (wseci *WorkflowStageExecutionCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wseci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range wseci.Steps {
		if wseci.Steps[i] == nil {
			continue
		}

		if err := wseci.Steps[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				wseci.Steps[i] = nil
			}
		}
	}

	return nil
}

// WorkflowStageExecutionCreateInputs holds the creation input of the WorkflowStageExecution entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStageExecutionCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// WorkflowExecution indicates to create WorkflowStageExecution entity MUST under the WorkflowExecution route.
	WorkflowExecution *WorkflowExecutionQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowStageExecutionCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the WorkflowStageExecution entities for creating,
// after validating.
func (wseci *WorkflowStageExecutionCreateInputs) Model() []*WorkflowStageExecution {
	if wseci == nil || len(wseci.Items) == 0 {
		return nil
	}

	_wses := make([]*WorkflowStageExecution, len(wseci.Items))

	for i := range wseci.Items {
		_wse := &WorkflowStageExecution{
			WorkflowStageID: wseci.Items[i].WorkflowStageID,
			WorkflowID:      wseci.Items[i].WorkflowID,
			Name:            wseci.Items[i].Name,
			Description:     wseci.Items[i].Description,
			Labels:          wseci.Items[i].Labels,
		}

		if wseci.WorkflowExecution != nil {
			_wse.WorkflowExecutionID = wseci.WorkflowExecution.ID
		}

		if wseci.Items[i].Steps != nil {
			// Empty slice is used for clearing the edge.
			_wse.Edges.Steps = make([]*WorkflowStepExecution, 0, len(wseci.Items[i].Steps))
		}
		for j := range wseci.Items[i].Steps {
			if wseci.Items[i].Steps[j] == nil {
				continue
			}
			_wse.Edges.Steps = append(_wse.Edges.Steps,
				wseci.Items[i].Steps[j].Model())
		}

		_wses[i] = _wse
	}

	return _wses
}

// Validate checks the WorkflowStageExecutionCreateInputs entity .
func (wseci *WorkflowStageExecutionCreateInputs) Validate() error {
	if wseci == nil {
		return errors.New("nil receiver")
	}

	return wseci.ValidateWith(wseci.inputConfig.Context, wseci.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStageExecutionCreateInputs entity with the given context and client set.
func (wseci *WorkflowStageExecutionCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wseci == nil {
		return errors.New("nil receiver")
	}

	if len(wseci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the WorkflowExecution route.
	if wseci.WorkflowExecution != nil {
		if err := wseci.WorkflowExecution.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				wseci.WorkflowExecution = nil
			}
		}
	}

	for i := range wseci.Items {
		if wseci.Items[i] == nil {
			continue
		}

		if err := wseci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowStageExecutionDeleteInput holds the deletion input of the WorkflowStageExecution entity,
// please tags with `path:",inline"` if embedding.
type WorkflowStageExecutionDeleteInput struct {
	WorkflowStageExecutionQueryInput `path:",inline"`
}

// WorkflowStageExecutionDeleteInputs holds the deletion input item of the WorkflowStageExecution entities.
type WorkflowStageExecutionDeleteInputsItem struct {
	// ID of the WorkflowStageExecution entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// WorkflowStageExecutionDeleteInputs holds the deletion input of the WorkflowStageExecution entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStageExecutionDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// WorkflowExecution indicates to delete WorkflowStageExecution entity MUST under the WorkflowExecution route.
	WorkflowExecution *WorkflowExecutionQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowStageExecutionDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the WorkflowStageExecution entities for deleting,
// after validating.
func (wsedi *WorkflowStageExecutionDeleteInputs) Model() []*WorkflowStageExecution {
	if wsedi == nil || len(wsedi.Items) == 0 {
		return nil
	}

	_wses := make([]*WorkflowStageExecution, len(wsedi.Items))
	for i := range wsedi.Items {
		_wses[i] = &WorkflowStageExecution{
			ID: wsedi.Items[i].ID,
		}
	}
	return _wses
}

// IDs returns the ID list of the WorkflowStageExecution entities for deleting,
// after validating.
func (wsedi *WorkflowStageExecutionDeleteInputs) IDs() []object.ID {
	if wsedi == nil || len(wsedi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(wsedi.Items))
	for i := range wsedi.Items {
		ids[i] = wsedi.Items[i].ID
	}
	return ids
}

// Validate checks the WorkflowStageExecutionDeleteInputs entity.
func (wsedi *WorkflowStageExecutionDeleteInputs) Validate() error {
	if wsedi == nil {
		return errors.New("nil receiver")
	}

	return wsedi.ValidateWith(wsedi.inputConfig.Context, wsedi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStageExecutionDeleteInputs entity with the given context and client set.
func (wsedi *WorkflowStageExecutionDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wsedi == nil {
		return errors.New("nil receiver")
	}

	if len(wsedi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.WorkflowStageExecutions().Query()

	// Validate when deleting under the WorkflowExecution route.
	if wsedi.WorkflowExecution != nil {
		if err := wsedi.WorkflowExecution.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				workflowstageexecution.WorkflowExecutionID(wsedi.WorkflowExecution.ID))
		}
	}

	ids := make([]object.ID, 0, len(wsedi.Items))

	for i := range wsedi.Items {
		if wsedi.Items[i] == nil {
			return errors.New("nil item")
		}

		if wsedi.Items[i].ID != "" {
			ids = append(ids, wsedi.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(workflowstageexecution.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	return nil
}

// WorkflowStageExecutionPatchInput holds the patch input of the WorkflowStageExecution entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStageExecutionPatchInput struct {
	WorkflowStageExecutionQueryInput `path:",inline" query:"-" json:"-"`

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
	// ID of the workflow that this workflow execution belongs to.
	WorkflowID object.ID `path:"-" query:"-" json:"workflowID,omitempty"`
	// ID of the workflow stage that this workflow stage execution belongs to.
	WorkflowStageID object.ID `path:"-" query:"-" json:"workflowStageID,omitempty"`
	// Time of the stage execution started.
	ExecuteTime time.Time `path:"-" query:"-" json:"executeTime,omitempty"`
	// Duration of the workflow stage execution.
	Duration int `path:"-" query:"-" json:"duration,omitempty"`
	// Order of the workflow stage execution.
	Order int `path:"-" query:"-" json:"order,omitempty"`

	// Steps indicates replacing the stale WorkflowStepExecution entities.
	Steps []*WorkflowStepExecutionCreateInput `uri:"-" query:"-" json:"steps,omitempty"`

	patchedEntity *WorkflowStageExecution `path:"-" query:"-" json:"-"`
}

// PatchModel returns the WorkflowStageExecution partition entity for patching.
func (wsepi *WorkflowStageExecutionPatchInput) PatchModel() *WorkflowStageExecution {
	if wsepi == nil {
		return nil
	}

	_wse := &WorkflowStageExecution{
		Name:            wsepi.Name,
		Description:     wsepi.Description,
		Labels:          wsepi.Labels,
		Annotations:     wsepi.Annotations,
		CreateTime:      wsepi.CreateTime,
		UpdateTime:      wsepi.UpdateTime,
		Status:          wsepi.Status,
		WorkflowID:      wsepi.WorkflowID,
		WorkflowStageID: wsepi.WorkflowStageID,
		ExecuteTime:     wsepi.ExecuteTime,
		Duration:        wsepi.Duration,
		Order:           wsepi.Order,
	}

	if wsepi.WorkflowExecution != nil {
		_wse.WorkflowExecutionID = wsepi.WorkflowExecution.ID
	}

	if wsepi.Steps != nil {
		// Empty slice is used for clearing the edge.
		_wse.Edges.Steps = make([]*WorkflowStepExecution, 0, len(wsepi.Steps))
	}
	for j := range wsepi.Steps {
		if wsepi.Steps[j] == nil {
			continue
		}
		_wse.Edges.Steps = append(_wse.Edges.Steps,
			wsepi.Steps[j].Model())
	}
	return _wse
}

// Model returns the WorkflowStageExecution patched entity,
// after validating.
func (wsepi *WorkflowStageExecutionPatchInput) Model() *WorkflowStageExecution {
	if wsepi == nil {
		return nil
	}

	return wsepi.patchedEntity
}

// Validate checks the WorkflowStageExecutionPatchInput entity.
func (wsepi *WorkflowStageExecutionPatchInput) Validate() error {
	if wsepi == nil {
		return errors.New("nil receiver")
	}

	return wsepi.ValidateWith(wsepi.inputConfig.Context, wsepi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStageExecutionPatchInput entity with the given context and client set.
func (wsepi *WorkflowStageExecutionPatchInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := wsepi.WorkflowStageExecutionQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	q := cs.WorkflowStageExecutions().Query()

	// Validate when querying under the WorkflowExecution route.
	if wsepi.WorkflowExecution != nil {
		if err := wsepi.WorkflowExecution.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				workflowstageexecution.WorkflowExecutionID(wsepi.WorkflowExecution.ID))
		}
	}

	for i := range wsepi.Steps {
		if wsepi.Steps[i] == nil {
			continue
		}

		if err := wsepi.Steps[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				wsepi.Steps[i] = nil
			}
		}
	}

	if wsepi.Refer != nil {
		if wsepi.Refer.IsID() {
			q.Where(
				workflowstageexecution.ID(wsepi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of workflowstageexecution")
		}
	} else if wsepi.ID != "" {
		q.Where(
			workflowstageexecution.ID(wsepi.ID))
	} else {
		return errors.New("invalid identify of workflowstageexecution")
	}

	q.Select(
		workflowstageexecution.WithoutFields(
			workflowstageexecution.FieldAnnotations,
			workflowstageexecution.FieldCreateTime,
			workflowstageexecution.FieldUpdateTime,
			workflowstageexecution.FieldStatus,
			workflowstageexecution.FieldExecuteTime,
			workflowstageexecution.FieldDuration,
			workflowstageexecution.FieldOrder,
		)...,
	)

	var e *WorkflowStageExecution
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
			e = cv.(*WorkflowStageExecution)
		}
	}

	_pm := wsepi.PatchModel()

	_po, err := json.PatchObject(*e, *_pm)
	if err != nil {
		return err
	}

	_obj := _po.(*WorkflowStageExecution)

	if e.Name != _obj.Name {
		return errors.New("field name is immutable")
	}
	if !reflect.DeepEqual(e.CreateTime, _obj.CreateTime) {
		return errors.New("field createTime is immutable")
	}
	if e.WorkflowID != _obj.WorkflowID {
		return errors.New("field workflowID is immutable")
	}
	if e.WorkflowStageID != _obj.WorkflowStageID {
		return errors.New("field workflowStageID is immutable")
	}

	wsepi.patchedEntity = _obj
	return nil
}

// WorkflowStageExecutionQueryInput holds the query input of the WorkflowStageExecution entity,
// please tags with `path:",inline"` if embedding.
type WorkflowStageExecutionQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// WorkflowExecution indicates to query WorkflowStageExecution entity MUST under the WorkflowExecution route.
	WorkflowExecution *WorkflowExecutionQueryInput `path:",inline" query:"-" json:"workflowExecution"`

	// Refer holds the route path reference of the WorkflowStageExecution entity.
	Refer *object.Refer `path:"workflowstageexecution,default=" query:"-" json:"-"`
	// ID of the WorkflowStageExecution entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// Model returns the WorkflowStageExecution entity for querying,
// after validating.
func (wseqi *WorkflowStageExecutionQueryInput) Model() *WorkflowStageExecution {
	if wseqi == nil {
		return nil
	}

	return &WorkflowStageExecution{
		ID: wseqi.ID,
	}
}

// Validate checks the WorkflowStageExecutionQueryInput entity.
func (wseqi *WorkflowStageExecutionQueryInput) Validate() error {
	if wseqi == nil {
		return errors.New("nil receiver")
	}

	return wseqi.ValidateWith(wseqi.inputConfig.Context, wseqi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStageExecutionQueryInput entity with the given context and client set.
func (wseqi *WorkflowStageExecutionQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wseqi == nil {
		return errors.New("nil receiver")
	}

	if wseqi.Refer != nil && *wseqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", workflowstageexecution.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.WorkflowStageExecutions().Query()

	// Validate when querying under the WorkflowExecution route.
	if wseqi.WorkflowExecution != nil {
		if err := wseqi.WorkflowExecution.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				workflowstageexecution.WorkflowExecutionID(wseqi.WorkflowExecution.ID))
		}
	}

	if wseqi.Refer != nil {
		if wseqi.Refer.IsID() {
			q.Where(
				workflowstageexecution.ID(wseqi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of workflowstageexecution")
		}
	} else if wseqi.ID != "" {
		q.Where(
			workflowstageexecution.ID(wseqi.ID))
	} else {
		return errors.New("invalid identify of workflowstageexecution")
	}

	q.Select(
		workflowstageexecution.FieldID,
	)

	var e *WorkflowStageExecution
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
			e = cv.(*WorkflowStageExecution)
		}
	}

	wseqi.ID = e.ID
	return nil
}

// WorkflowStageExecutionQueryInputs holds the query input of the WorkflowStageExecution entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type WorkflowStageExecutionQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// WorkflowExecution indicates to query WorkflowStageExecution entity MUST under the WorkflowExecution route.
	WorkflowExecution *WorkflowExecutionQueryInput `path:",inline" query:"-" json:"-"`
}

// Validate checks the WorkflowStageExecutionQueryInputs entity.
func (wseqi *WorkflowStageExecutionQueryInputs) Validate() error {
	if wseqi == nil {
		return errors.New("nil receiver")
	}

	return wseqi.ValidateWith(wseqi.inputConfig.Context, wseqi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStageExecutionQueryInputs entity with the given context and client set.
func (wseqi *WorkflowStageExecutionQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wseqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when querying under the WorkflowExecution route.
	if wseqi.WorkflowExecution != nil {
		if err := wseqi.WorkflowExecution.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowStageExecutionUpdateInput holds the modification input of the WorkflowStageExecution entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStageExecutionUpdateInput struct {
	WorkflowStageExecutionQueryInput `path:",inline" query:"-" json:"-"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`

	// Steps indicates replacing the stale WorkflowStepExecution entities.
	Steps []*WorkflowStepExecutionCreateInput `uri:"-" query:"-" json:"steps,omitempty"`
}

// Model returns the WorkflowStageExecution entity for modifying,
// after validating.
func (wseui *WorkflowStageExecutionUpdateInput) Model() *WorkflowStageExecution {
	if wseui == nil {
		return nil
	}

	_wse := &WorkflowStageExecution{
		ID:          wseui.ID,
		Description: wseui.Description,
		Labels:      wseui.Labels,
	}

	if wseui.Steps != nil {
		// Empty slice is used for clearing the edge.
		_wse.Edges.Steps = make([]*WorkflowStepExecution, 0, len(wseui.Steps))
	}
	for j := range wseui.Steps {
		if wseui.Steps[j] == nil {
			continue
		}
		_wse.Edges.Steps = append(_wse.Edges.Steps,
			wseui.Steps[j].Model())
	}
	return _wse
}

// Validate checks the WorkflowStageExecutionUpdateInput entity.
func (wseui *WorkflowStageExecutionUpdateInput) Validate() error {
	if wseui == nil {
		return errors.New("nil receiver")
	}

	return wseui.ValidateWith(wseui.inputConfig.Context, wseui.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStageExecutionUpdateInput entity with the given context and client set.
func (wseui *WorkflowStageExecutionUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := wseui.WorkflowStageExecutionQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	for i := range wseui.Steps {
		if wseui.Steps[i] == nil {
			continue
		}

		if err := wseui.Steps[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				wseui.Steps[i] = nil
			}
		}
	}

	return nil
}

// WorkflowStageExecutionUpdateInputs holds the modification input item of the WorkflowStageExecution entities.
type WorkflowStageExecutionUpdateInputsItem struct {
	// ID of the WorkflowStageExecution entity.
	ID object.ID `path:"-" query:"-" json:"id"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`

	// Steps indicates replacing the stale WorkflowStepExecution entities.
	Steps []*WorkflowStepExecutionCreateInput `uri:"-" query:"-" json:"steps,omitempty"`
}

// ValidateWith checks the WorkflowStageExecutionUpdateInputsItem entity with the given context and client set.
func (wseui *WorkflowStageExecutionUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wseui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range wseui.Steps {
		if wseui.Steps[i] == nil {
			continue
		}

		if err := wseui.Steps[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				wseui.Steps[i] = nil
			}
		}
	}

	return nil
}

// WorkflowStageExecutionUpdateInputs holds the modification input of the WorkflowStageExecution entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStageExecutionUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// WorkflowExecution indicates to update WorkflowStageExecution entity MUST under the WorkflowExecution route.
	WorkflowExecution *WorkflowExecutionQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowStageExecutionUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the WorkflowStageExecution entities for modifying,
// after validating.
func (wseui *WorkflowStageExecutionUpdateInputs) Model() []*WorkflowStageExecution {
	if wseui == nil || len(wseui.Items) == 0 {
		return nil
	}

	_wses := make([]*WorkflowStageExecution, len(wseui.Items))

	for i := range wseui.Items {
		_wse := &WorkflowStageExecution{
			ID:          wseui.Items[i].ID,
			Description: wseui.Items[i].Description,
			Labels:      wseui.Items[i].Labels,
		}

		if wseui.Items[i].Steps != nil {
			// Empty slice is used for clearing the edge.
			_wse.Edges.Steps = make([]*WorkflowStepExecution, 0, len(wseui.Items[i].Steps))
		}
		for j := range wseui.Items[i].Steps {
			if wseui.Items[i].Steps[j] == nil {
				continue
			}
			_wse.Edges.Steps = append(_wse.Edges.Steps,
				wseui.Items[i].Steps[j].Model())
		}

		_wses[i] = _wse
	}

	return _wses
}

// IDs returns the ID list of the WorkflowStageExecution entities for modifying,
// after validating.
func (wseui *WorkflowStageExecutionUpdateInputs) IDs() []object.ID {
	if wseui == nil || len(wseui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(wseui.Items))
	for i := range wseui.Items {
		ids[i] = wseui.Items[i].ID
	}
	return ids
}

// Validate checks the WorkflowStageExecutionUpdateInputs entity.
func (wseui *WorkflowStageExecutionUpdateInputs) Validate() error {
	if wseui == nil {
		return errors.New("nil receiver")
	}

	return wseui.ValidateWith(wseui.inputConfig.Context, wseui.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStageExecutionUpdateInputs entity with the given context and client set.
func (wseui *WorkflowStageExecutionUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wseui == nil {
		return errors.New("nil receiver")
	}

	if len(wseui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.WorkflowStageExecutions().Query()

	// Validate when updating under the WorkflowExecution route.
	if wseui.WorkflowExecution != nil {
		if err := wseui.WorkflowExecution.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				workflowstageexecution.WorkflowExecutionID(wseui.WorkflowExecution.ID))
		}
	}

	ids := make([]object.ID, 0, len(wseui.Items))

	for i := range wseui.Items {
		if wseui.Items[i] == nil {
			return errors.New("nil item")
		}

		if wseui.Items[i].ID != "" {
			ids = append(ids, wseui.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(workflowstageexecution.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range wseui.Items {
		if err := wseui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowStageExecutionOutput holds the output of the WorkflowStageExecution entity.
type WorkflowStageExecutionOutput struct {
	ID              object.ID         `json:"id,omitempty"`
	Name            string            `json:"name,omitempty"`
	Description     string            `json:"description,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"`
	CreateTime      *time.Time        `json:"createTime,omitempty"`
	UpdateTime      *time.Time        `json:"updateTime,omitempty"`
	Status          status.Status     `json:"status,omitempty"`
	WorkflowID      object.ID         `json:"workflowID,omitempty"`
	WorkflowStageID object.ID         `json:"workflowStageID,omitempty"`
	ExecuteTime     time.Time         `json:"executeTime,omitempty"`

	Project           *ProjectOutput                 `json:"project,omitempty"`
	Steps             []*WorkflowStepExecutionOutput `json:"steps,omitempty"`
	WorkflowExecution *WorkflowExecutionOutput       `json:"workflowExecution,omitempty"`
}

// View returns the output of WorkflowStageExecution entity.
func (_wse *WorkflowStageExecution) View() *WorkflowStageExecutionOutput {
	return ExposeWorkflowStageExecution(_wse)
}

// View returns the output of WorkflowStageExecution entities.
func (_wses WorkflowStageExecutions) View() []*WorkflowStageExecutionOutput {
	return ExposeWorkflowStageExecutions(_wses)
}

// ExposeWorkflowStageExecution converts the WorkflowStageExecution to WorkflowStageExecutionOutput.
func ExposeWorkflowStageExecution(_wse *WorkflowStageExecution) *WorkflowStageExecutionOutput {
	if _wse == nil {
		return nil
	}

	wseo := &WorkflowStageExecutionOutput{
		ID:              _wse.ID,
		Name:            _wse.Name,
		Description:     _wse.Description,
		Labels:          _wse.Labels,
		CreateTime:      _wse.CreateTime,
		UpdateTime:      _wse.UpdateTime,
		Status:          _wse.Status,
		WorkflowID:      _wse.WorkflowID,
		WorkflowStageID: _wse.WorkflowStageID,
		ExecuteTime:     _wse.ExecuteTime,
	}

	if _wse.Edges.Project != nil {
		wseo.Project = ExposeProject(_wse.Edges.Project)
	} else if _wse.ProjectID != "" {
		wseo.Project = &ProjectOutput{
			ID: _wse.ProjectID,
		}
	}
	if _wse.Edges.Steps != nil {
		wseo.Steps = ExposeWorkflowStepExecutions(_wse.Edges.Steps)
	}
	if _wse.Edges.WorkflowExecution != nil {
		wseo.WorkflowExecution = ExposeWorkflowExecution(_wse.Edges.WorkflowExecution)
	} else if _wse.WorkflowExecutionID != "" {
		wseo.WorkflowExecution = &WorkflowExecutionOutput{
			ID: _wse.WorkflowExecutionID,
		}
	}
	return wseo
}

// ExposeWorkflowStageExecutions converts the WorkflowStageExecution slice to WorkflowStageExecutionOutput pointer slice.
func ExposeWorkflowStageExecutions(_wses []*WorkflowStageExecution) []*WorkflowStageExecutionOutput {
	if len(_wses) == 0 {
		return nil
	}

	wseos := make([]*WorkflowStageExecutionOutput, len(_wses))
	for i := range _wses {
		wseos[i] = ExposeWorkflowStageExecution(_wses[i])
	}
	return wseos
}
