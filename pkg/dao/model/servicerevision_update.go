// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "ent". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/model/servicerevision"
	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/dao/types/crypto"
	"github.com/seal-io/seal/pkg/dao/types/property"
)

// ServiceRevisionUpdate is the builder for updating ServiceRevision entities.
type ServiceRevisionUpdate struct {
	config
	hooks     []Hook
	mutation  *ServiceRevisionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ServiceRevisionUpdate builder.
func (sru *ServiceRevisionUpdate) Where(ps ...predicate.ServiceRevision) *ServiceRevisionUpdate {
	sru.mutation.Where(ps...)
	return sru
}

// SetStatus sets the "status" field.
func (sru *ServiceRevisionUpdate) SetStatus(s string) *ServiceRevisionUpdate {
	sru.mutation.SetStatus(s)
	return sru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sru *ServiceRevisionUpdate) SetNillableStatus(s *string) *ServiceRevisionUpdate {
	if s != nil {
		sru.SetStatus(*s)
	}
	return sru
}

// ClearStatus clears the value of the "status" field.
func (sru *ServiceRevisionUpdate) ClearStatus() *ServiceRevisionUpdate {
	sru.mutation.ClearStatus()
	return sru
}

// SetStatusMessage sets the "statusMessage" field.
func (sru *ServiceRevisionUpdate) SetStatusMessage(s string) *ServiceRevisionUpdate {
	sru.mutation.SetStatusMessage(s)
	return sru
}

// SetNillableStatusMessage sets the "statusMessage" field if the given value is not nil.
func (sru *ServiceRevisionUpdate) SetNillableStatusMessage(s *string) *ServiceRevisionUpdate {
	if s != nil {
		sru.SetStatusMessage(*s)
	}
	return sru
}

// ClearStatusMessage clears the value of the "statusMessage" field.
func (sru *ServiceRevisionUpdate) ClearStatusMessage() *ServiceRevisionUpdate {
	sru.mutation.ClearStatusMessage()
	return sru
}

// SetTemplateVersion sets the "templateVersion" field.
func (sru *ServiceRevisionUpdate) SetTemplateVersion(s string) *ServiceRevisionUpdate {
	sru.mutation.SetTemplateVersion(s)
	return sru
}

// SetAttributes sets the "attributes" field.
func (sru *ServiceRevisionUpdate) SetAttributes(pr property.Values) *ServiceRevisionUpdate {
	sru.mutation.SetAttributes(pr)
	return sru
}

// ClearAttributes clears the value of the "attributes" field.
func (sru *ServiceRevisionUpdate) ClearAttributes() *ServiceRevisionUpdate {
	sru.mutation.ClearAttributes()
	return sru
}

// SetSecrets sets the "secrets" field.
func (sru *ServiceRevisionUpdate) SetSecrets(c crypto.Map[string, string]) *ServiceRevisionUpdate {
	sru.mutation.SetSecrets(c)
	return sru
}

// SetInputPlan sets the "inputPlan" field.
func (sru *ServiceRevisionUpdate) SetInputPlan(s string) *ServiceRevisionUpdate {
	sru.mutation.SetInputPlan(s)
	return sru
}

// SetOutput sets the "output" field.
func (sru *ServiceRevisionUpdate) SetOutput(s string) *ServiceRevisionUpdate {
	sru.mutation.SetOutput(s)
	return sru
}

// SetDeployerType sets the "deployerType" field.
func (sru *ServiceRevisionUpdate) SetDeployerType(s string) *ServiceRevisionUpdate {
	sru.mutation.SetDeployerType(s)
	return sru
}

// SetNillableDeployerType sets the "deployerType" field if the given value is not nil.
func (sru *ServiceRevisionUpdate) SetNillableDeployerType(s *string) *ServiceRevisionUpdate {
	if s != nil {
		sru.SetDeployerType(*s)
	}
	return sru
}

// SetDuration sets the "duration" field.
func (sru *ServiceRevisionUpdate) SetDuration(i int) *ServiceRevisionUpdate {
	sru.mutation.ResetDuration()
	sru.mutation.SetDuration(i)
	return sru
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (sru *ServiceRevisionUpdate) SetNillableDuration(i *int) *ServiceRevisionUpdate {
	if i != nil {
		sru.SetDuration(*i)
	}
	return sru
}

// AddDuration adds i to the "duration" field.
func (sru *ServiceRevisionUpdate) AddDuration(i int) *ServiceRevisionUpdate {
	sru.mutation.AddDuration(i)
	return sru
}

// SetPreviousRequiredProviders sets the "previousRequiredProviders" field.
func (sru *ServiceRevisionUpdate) SetPreviousRequiredProviders(tr []types.ProviderRequirement) *ServiceRevisionUpdate {
	sru.mutation.SetPreviousRequiredProviders(tr)
	return sru
}

// AppendPreviousRequiredProviders appends tr to the "previousRequiredProviders" field.
func (sru *ServiceRevisionUpdate) AppendPreviousRequiredProviders(tr []types.ProviderRequirement) *ServiceRevisionUpdate {
	sru.mutation.AppendPreviousRequiredProviders(tr)
	return sru
}

// SetTags sets the "tags" field.
func (sru *ServiceRevisionUpdate) SetTags(s []string) *ServiceRevisionUpdate {
	sru.mutation.SetTags(s)
	return sru
}

// AppendTags appends s to the "tags" field.
func (sru *ServiceRevisionUpdate) AppendTags(s []string) *ServiceRevisionUpdate {
	sru.mutation.AppendTags(s)
	return sru
}

// Mutation returns the ServiceRevisionMutation object of the builder.
func (sru *ServiceRevisionUpdate) Mutation() *ServiceRevisionMutation {
	return sru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sru *ServiceRevisionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ServiceRevisionMutation](ctx, sru.sqlSave, sru.mutation, sru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sru *ServiceRevisionUpdate) SaveX(ctx context.Context) int {
	affected, err := sru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sru *ServiceRevisionUpdate) Exec(ctx context.Context) error {
	_, err := sru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sru *ServiceRevisionUpdate) ExecX(ctx context.Context) {
	if err := sru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sru *ServiceRevisionUpdate) check() error {
	if v, ok := sru.mutation.TemplateVersion(); ok {
		if err := servicerevision.TemplateVersionValidator(v); err != nil {
			return &ValidationError{Name: "templateVersion", err: fmt.Errorf(`model: validator failed for field "ServiceRevision.templateVersion": %w`, err)}
		}
	}
	if _, ok := sru.mutation.ServiceID(); sru.mutation.ServiceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRevision.service"`)
	}
	if _, ok := sru.mutation.EnvironmentID(); sru.mutation.EnvironmentCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRevision.environment"`)
	}
	if _, ok := sru.mutation.ProjectID(); sru.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRevision.project"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sru *ServiceRevisionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ServiceRevisionUpdate {
	sru.modifiers = append(sru.modifiers, modifiers...)
	return sru
}

func (sru *ServiceRevisionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := sru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(servicerevision.Table, servicerevision.Columns, sqlgraph.NewFieldSpec(servicerevision.FieldID, field.TypeString))
	if ps := sru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sru.mutation.Status(); ok {
		_spec.SetField(servicerevision.FieldStatus, field.TypeString, value)
	}
	if sru.mutation.StatusCleared() {
		_spec.ClearField(servicerevision.FieldStatus, field.TypeString)
	}
	if value, ok := sru.mutation.StatusMessage(); ok {
		_spec.SetField(servicerevision.FieldStatusMessage, field.TypeString, value)
	}
	if sru.mutation.StatusMessageCleared() {
		_spec.ClearField(servicerevision.FieldStatusMessage, field.TypeString)
	}
	if value, ok := sru.mutation.TemplateVersion(); ok {
		_spec.SetField(servicerevision.FieldTemplateVersion, field.TypeString, value)
	}
	if value, ok := sru.mutation.Attributes(); ok {
		_spec.SetField(servicerevision.FieldAttributes, field.TypeOther, value)
	}
	if sru.mutation.AttributesCleared() {
		_spec.ClearField(servicerevision.FieldAttributes, field.TypeOther)
	}
	if value, ok := sru.mutation.Secrets(); ok {
		_spec.SetField(servicerevision.FieldSecrets, field.TypeOther, value)
	}
	if value, ok := sru.mutation.InputPlan(); ok {
		_spec.SetField(servicerevision.FieldInputPlan, field.TypeString, value)
	}
	if value, ok := sru.mutation.Output(); ok {
		_spec.SetField(servicerevision.FieldOutput, field.TypeString, value)
	}
	if value, ok := sru.mutation.DeployerType(); ok {
		_spec.SetField(servicerevision.FieldDeployerType, field.TypeString, value)
	}
	if value, ok := sru.mutation.Duration(); ok {
		_spec.SetField(servicerevision.FieldDuration, field.TypeInt, value)
	}
	if value, ok := sru.mutation.AddedDuration(); ok {
		_spec.AddField(servicerevision.FieldDuration, field.TypeInt, value)
	}
	if value, ok := sru.mutation.PreviousRequiredProviders(); ok {
		_spec.SetField(servicerevision.FieldPreviousRequiredProviders, field.TypeJSON, value)
	}
	if value, ok := sru.mutation.AppendedPreviousRequiredProviders(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, servicerevision.FieldPreviousRequiredProviders, value)
		})
	}
	if value, ok := sru.mutation.Tags(); ok {
		_spec.SetField(servicerevision.FieldTags, field.TypeJSON, value)
	}
	if value, ok := sru.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, servicerevision.FieldTags, value)
		})
	}
	_spec.Node.Schema = sru.schemaConfig.ServiceRevision
	ctx = internal.NewSchemaConfigContext(ctx, sru.schemaConfig)
	_spec.AddModifiers(sru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, sru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servicerevision.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sru.mutation.done = true
	return n, nil
}

// ServiceRevisionUpdateOne is the builder for updating a single ServiceRevision entity.
type ServiceRevisionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ServiceRevisionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetStatus sets the "status" field.
func (sruo *ServiceRevisionUpdateOne) SetStatus(s string) *ServiceRevisionUpdateOne {
	sruo.mutation.SetStatus(s)
	return sruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sruo *ServiceRevisionUpdateOne) SetNillableStatus(s *string) *ServiceRevisionUpdateOne {
	if s != nil {
		sruo.SetStatus(*s)
	}
	return sruo
}

// ClearStatus clears the value of the "status" field.
func (sruo *ServiceRevisionUpdateOne) ClearStatus() *ServiceRevisionUpdateOne {
	sruo.mutation.ClearStatus()
	return sruo
}

// SetStatusMessage sets the "statusMessage" field.
func (sruo *ServiceRevisionUpdateOne) SetStatusMessage(s string) *ServiceRevisionUpdateOne {
	sruo.mutation.SetStatusMessage(s)
	return sruo
}

// SetNillableStatusMessage sets the "statusMessage" field if the given value is not nil.
func (sruo *ServiceRevisionUpdateOne) SetNillableStatusMessage(s *string) *ServiceRevisionUpdateOne {
	if s != nil {
		sruo.SetStatusMessage(*s)
	}
	return sruo
}

// ClearStatusMessage clears the value of the "statusMessage" field.
func (sruo *ServiceRevisionUpdateOne) ClearStatusMessage() *ServiceRevisionUpdateOne {
	sruo.mutation.ClearStatusMessage()
	return sruo
}

// SetTemplateVersion sets the "templateVersion" field.
func (sruo *ServiceRevisionUpdateOne) SetTemplateVersion(s string) *ServiceRevisionUpdateOne {
	sruo.mutation.SetTemplateVersion(s)
	return sruo
}

// SetAttributes sets the "attributes" field.
func (sruo *ServiceRevisionUpdateOne) SetAttributes(pr property.Values) *ServiceRevisionUpdateOne {
	sruo.mutation.SetAttributes(pr)
	return sruo
}

// ClearAttributes clears the value of the "attributes" field.
func (sruo *ServiceRevisionUpdateOne) ClearAttributes() *ServiceRevisionUpdateOne {
	sruo.mutation.ClearAttributes()
	return sruo
}

// SetSecrets sets the "secrets" field.
func (sruo *ServiceRevisionUpdateOne) SetSecrets(c crypto.Map[string, string]) *ServiceRevisionUpdateOne {
	sruo.mutation.SetSecrets(c)
	return sruo
}

// SetInputPlan sets the "inputPlan" field.
func (sruo *ServiceRevisionUpdateOne) SetInputPlan(s string) *ServiceRevisionUpdateOne {
	sruo.mutation.SetInputPlan(s)
	return sruo
}

// SetOutput sets the "output" field.
func (sruo *ServiceRevisionUpdateOne) SetOutput(s string) *ServiceRevisionUpdateOne {
	sruo.mutation.SetOutput(s)
	return sruo
}

// SetDeployerType sets the "deployerType" field.
func (sruo *ServiceRevisionUpdateOne) SetDeployerType(s string) *ServiceRevisionUpdateOne {
	sruo.mutation.SetDeployerType(s)
	return sruo
}

// SetNillableDeployerType sets the "deployerType" field if the given value is not nil.
func (sruo *ServiceRevisionUpdateOne) SetNillableDeployerType(s *string) *ServiceRevisionUpdateOne {
	if s != nil {
		sruo.SetDeployerType(*s)
	}
	return sruo
}

// SetDuration sets the "duration" field.
func (sruo *ServiceRevisionUpdateOne) SetDuration(i int) *ServiceRevisionUpdateOne {
	sruo.mutation.ResetDuration()
	sruo.mutation.SetDuration(i)
	return sruo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (sruo *ServiceRevisionUpdateOne) SetNillableDuration(i *int) *ServiceRevisionUpdateOne {
	if i != nil {
		sruo.SetDuration(*i)
	}
	return sruo
}

// AddDuration adds i to the "duration" field.
func (sruo *ServiceRevisionUpdateOne) AddDuration(i int) *ServiceRevisionUpdateOne {
	sruo.mutation.AddDuration(i)
	return sruo
}

// SetPreviousRequiredProviders sets the "previousRequiredProviders" field.
func (sruo *ServiceRevisionUpdateOne) SetPreviousRequiredProviders(tr []types.ProviderRequirement) *ServiceRevisionUpdateOne {
	sruo.mutation.SetPreviousRequiredProviders(tr)
	return sruo
}

// AppendPreviousRequiredProviders appends tr to the "previousRequiredProviders" field.
func (sruo *ServiceRevisionUpdateOne) AppendPreviousRequiredProviders(tr []types.ProviderRequirement) *ServiceRevisionUpdateOne {
	sruo.mutation.AppendPreviousRequiredProviders(tr)
	return sruo
}

// SetTags sets the "tags" field.
func (sruo *ServiceRevisionUpdateOne) SetTags(s []string) *ServiceRevisionUpdateOne {
	sruo.mutation.SetTags(s)
	return sruo
}

// AppendTags appends s to the "tags" field.
func (sruo *ServiceRevisionUpdateOne) AppendTags(s []string) *ServiceRevisionUpdateOne {
	sruo.mutation.AppendTags(s)
	return sruo
}

// Mutation returns the ServiceRevisionMutation object of the builder.
func (sruo *ServiceRevisionUpdateOne) Mutation() *ServiceRevisionMutation {
	return sruo.mutation
}

// Where appends a list predicates to the ServiceRevisionUpdate builder.
func (sruo *ServiceRevisionUpdateOne) Where(ps ...predicate.ServiceRevision) *ServiceRevisionUpdateOne {
	sruo.mutation.Where(ps...)
	return sruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sruo *ServiceRevisionUpdateOne) Select(field string, fields ...string) *ServiceRevisionUpdateOne {
	sruo.fields = append([]string{field}, fields...)
	return sruo
}

// Save executes the query and returns the updated ServiceRevision entity.
func (sruo *ServiceRevisionUpdateOne) Save(ctx context.Context) (*ServiceRevision, error) {
	return withHooks[*ServiceRevision, ServiceRevisionMutation](ctx, sruo.sqlSave, sruo.mutation, sruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sruo *ServiceRevisionUpdateOne) SaveX(ctx context.Context) *ServiceRevision {
	node, err := sruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sruo *ServiceRevisionUpdateOne) Exec(ctx context.Context) error {
	_, err := sruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sruo *ServiceRevisionUpdateOne) ExecX(ctx context.Context) {
	if err := sruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sruo *ServiceRevisionUpdateOne) check() error {
	if v, ok := sruo.mutation.TemplateVersion(); ok {
		if err := servicerevision.TemplateVersionValidator(v); err != nil {
			return &ValidationError{Name: "templateVersion", err: fmt.Errorf(`model: validator failed for field "ServiceRevision.templateVersion": %w`, err)}
		}
	}
	if _, ok := sruo.mutation.ServiceID(); sruo.mutation.ServiceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRevision.service"`)
	}
	if _, ok := sruo.mutation.EnvironmentID(); sruo.mutation.EnvironmentCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRevision.environment"`)
	}
	if _, ok := sruo.mutation.ProjectID(); sruo.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRevision.project"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sruo *ServiceRevisionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ServiceRevisionUpdateOne {
	sruo.modifiers = append(sruo.modifiers, modifiers...)
	return sruo
}

func (sruo *ServiceRevisionUpdateOne) sqlSave(ctx context.Context) (_node *ServiceRevision, err error) {
	if err := sruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(servicerevision.Table, servicerevision.Columns, sqlgraph.NewFieldSpec(servicerevision.FieldID, field.TypeString))
	id, ok := sruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "ServiceRevision.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, servicerevision.FieldID)
		for _, f := range fields {
			if !servicerevision.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != servicerevision.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sruo.mutation.Status(); ok {
		_spec.SetField(servicerevision.FieldStatus, field.TypeString, value)
	}
	if sruo.mutation.StatusCleared() {
		_spec.ClearField(servicerevision.FieldStatus, field.TypeString)
	}
	if value, ok := sruo.mutation.StatusMessage(); ok {
		_spec.SetField(servicerevision.FieldStatusMessage, field.TypeString, value)
	}
	if sruo.mutation.StatusMessageCleared() {
		_spec.ClearField(servicerevision.FieldStatusMessage, field.TypeString)
	}
	if value, ok := sruo.mutation.TemplateVersion(); ok {
		_spec.SetField(servicerevision.FieldTemplateVersion, field.TypeString, value)
	}
	if value, ok := sruo.mutation.Attributes(); ok {
		_spec.SetField(servicerevision.FieldAttributes, field.TypeOther, value)
	}
	if sruo.mutation.AttributesCleared() {
		_spec.ClearField(servicerevision.FieldAttributes, field.TypeOther)
	}
	if value, ok := sruo.mutation.Secrets(); ok {
		_spec.SetField(servicerevision.FieldSecrets, field.TypeOther, value)
	}
	if value, ok := sruo.mutation.InputPlan(); ok {
		_spec.SetField(servicerevision.FieldInputPlan, field.TypeString, value)
	}
	if value, ok := sruo.mutation.Output(); ok {
		_spec.SetField(servicerevision.FieldOutput, field.TypeString, value)
	}
	if value, ok := sruo.mutation.DeployerType(); ok {
		_spec.SetField(servicerevision.FieldDeployerType, field.TypeString, value)
	}
	if value, ok := sruo.mutation.Duration(); ok {
		_spec.SetField(servicerevision.FieldDuration, field.TypeInt, value)
	}
	if value, ok := sruo.mutation.AddedDuration(); ok {
		_spec.AddField(servicerevision.FieldDuration, field.TypeInt, value)
	}
	if value, ok := sruo.mutation.PreviousRequiredProviders(); ok {
		_spec.SetField(servicerevision.FieldPreviousRequiredProviders, field.TypeJSON, value)
	}
	if value, ok := sruo.mutation.AppendedPreviousRequiredProviders(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, servicerevision.FieldPreviousRequiredProviders, value)
		})
	}
	if value, ok := sruo.mutation.Tags(); ok {
		_spec.SetField(servicerevision.FieldTags, field.TypeJSON, value)
	}
	if value, ok := sruo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, servicerevision.FieldTags, value)
		})
	}
	_spec.Node.Schema = sruo.schemaConfig.ServiceRevision
	ctx = internal.NewSchemaConfigContext(ctx, sruo.schemaConfig)
	_spec.AddModifiers(sruo.modifiers...)
	_node = &ServiceRevision{config: sruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servicerevision.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sruo.mutation.done = true
	return _node, nil
}
