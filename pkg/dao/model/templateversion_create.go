// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/resource"
	"github.com/seal-io/walrus/pkg/dao/model/template"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// TemplateVersionCreate is the builder for creating a TemplateVersion entity.
type TemplateVersionCreate struct {
	config
	mutation   *TemplateVersionMutation
	hooks      []Hook
	conflict   []sql.ConflictOption
	object     *TemplateVersion
	fromUpsert bool
}

// SetCreateTime sets the "create_time" field.
func (tvc *TemplateVersionCreate) SetCreateTime(t time.Time) *TemplateVersionCreate {
	tvc.mutation.SetCreateTime(t)
	return tvc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (tvc *TemplateVersionCreate) SetNillableCreateTime(t *time.Time) *TemplateVersionCreate {
	if t != nil {
		tvc.SetCreateTime(*t)
	}
	return tvc
}

// SetUpdateTime sets the "update_time" field.
func (tvc *TemplateVersionCreate) SetUpdateTime(t time.Time) *TemplateVersionCreate {
	tvc.mutation.SetUpdateTime(t)
	return tvc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (tvc *TemplateVersionCreate) SetNillableUpdateTime(t *time.Time) *TemplateVersionCreate {
	if t != nil {
		tvc.SetUpdateTime(*t)
	}
	return tvc
}

// SetTemplateID sets the "template_id" field.
func (tvc *TemplateVersionCreate) SetTemplateID(o object.ID) *TemplateVersionCreate {
	tvc.mutation.SetTemplateID(o)
	return tvc
}

// SetName sets the "name" field.
func (tvc *TemplateVersionCreate) SetName(s string) *TemplateVersionCreate {
	tvc.mutation.SetName(s)
	return tvc
}

// SetVersion sets the "version" field.
func (tvc *TemplateVersionCreate) SetVersion(s string) *TemplateVersionCreate {
	tvc.mutation.SetVersion(s)
	return tvc
}

// SetSource sets the "source" field.
func (tvc *TemplateVersionCreate) SetSource(s string) *TemplateVersionCreate {
	tvc.mutation.SetSource(s)
	return tvc
}

// SetSchema sets the "schema" field.
func (tvc *TemplateVersionCreate) SetSchema(t types.Schema) *TemplateVersionCreate {
	tvc.mutation.SetSchema(t)
	return tvc
}

// SetNillableSchema sets the "schema" field if the given value is not nil.
func (tvc *TemplateVersionCreate) SetNillableSchema(t *types.Schema) *TemplateVersionCreate {
	if t != nil {
		tvc.SetSchema(*t)
	}
	return tvc
}

// SetUiSchema sets the "uiSchema" field.
func (tvc *TemplateVersionCreate) SetUiSchema(ts types.UISchema) *TemplateVersionCreate {
	tvc.mutation.SetUiSchema(ts)
	return tvc
}

// SetNillableUiSchema sets the "uiSchema" field if the given value is not nil.
func (tvc *TemplateVersionCreate) SetNillableUiSchema(ts *types.UISchema) *TemplateVersionCreate {
	if ts != nil {
		tvc.SetUiSchema(*ts)
	}
	return tvc
}

// SetProjectID sets the "project_id" field.
func (tvc *TemplateVersionCreate) SetProjectID(o object.ID) *TemplateVersionCreate {
	tvc.mutation.SetProjectID(o)
	return tvc
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (tvc *TemplateVersionCreate) SetNillableProjectID(o *object.ID) *TemplateVersionCreate {
	if o != nil {
		tvc.SetProjectID(*o)
	}
	return tvc
}

// SetID sets the "id" field.
func (tvc *TemplateVersionCreate) SetID(o object.ID) *TemplateVersionCreate {
	tvc.mutation.SetID(o)
	return tvc
}

// SetTemplate sets the "template" edge to the Template entity.
func (tvc *TemplateVersionCreate) SetTemplate(t *Template) *TemplateVersionCreate {
	return tvc.SetTemplateID(t.ID)
}

// AddResourceIDs adds the "resources" edge to the Resource entity by IDs.
func (tvc *TemplateVersionCreate) AddResourceIDs(ids ...object.ID) *TemplateVersionCreate {
	tvc.mutation.AddResourceIDs(ids...)
	return tvc
}

// AddResources adds the "resources" edges to the Resource entity.
func (tvc *TemplateVersionCreate) AddResources(r ...*Resource) *TemplateVersionCreate {
	ids := make([]object.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tvc.AddResourceIDs(ids...)
}

// SetProject sets the "project" edge to the Project entity.
func (tvc *TemplateVersionCreate) SetProject(p *Project) *TemplateVersionCreate {
	return tvc.SetProjectID(p.ID)
}

// Mutation returns the TemplateVersionMutation object of the builder.
func (tvc *TemplateVersionCreate) Mutation() *TemplateVersionMutation {
	return tvc.mutation
}

// Save creates the TemplateVersion in the database.
func (tvc *TemplateVersionCreate) Save(ctx context.Context) (*TemplateVersion, error) {
	if err := tvc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tvc.sqlSave, tvc.mutation, tvc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tvc *TemplateVersionCreate) SaveX(ctx context.Context) *TemplateVersion {
	v, err := tvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tvc *TemplateVersionCreate) Exec(ctx context.Context) error {
	_, err := tvc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tvc *TemplateVersionCreate) ExecX(ctx context.Context) {
	if err := tvc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tvc *TemplateVersionCreate) defaults() error {
	if _, ok := tvc.mutation.CreateTime(); !ok {
		if templateversion.DefaultCreateTime == nil {
			return fmt.Errorf("model: uninitialized templateversion.DefaultCreateTime (forgotten import model/runtime?)")
		}
		v := templateversion.DefaultCreateTime()
		tvc.mutation.SetCreateTime(v)
	}
	if _, ok := tvc.mutation.UpdateTime(); !ok {
		if templateversion.DefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized templateversion.DefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := templateversion.DefaultUpdateTime()
		tvc.mutation.SetUpdateTime(v)
	}
	if _, ok := tvc.mutation.Schema(); !ok {
		v := templateversion.DefaultSchema
		tvc.mutation.SetSchema(v)
	}
	if _, ok := tvc.mutation.UiSchema(); !ok {
		v := templateversion.DefaultUiSchema
		tvc.mutation.SetUiSchema(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tvc *TemplateVersionCreate) check() error {
	if _, ok := tvc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "TemplateVersion.create_time"`)}
	}
	if _, ok := tvc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "TemplateVersion.update_time"`)}
	}
	if _, ok := tvc.mutation.TemplateID(); !ok {
		return &ValidationError{Name: "template_id", err: errors.New(`model: missing required field "TemplateVersion.template_id"`)}
	}
	if v, ok := tvc.mutation.TemplateID(); ok {
		if err := templateversion.TemplateIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "template_id", err: fmt.Errorf(`model: validator failed for field "TemplateVersion.template_id": %w`, err)}
		}
	}
	if _, ok := tvc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model: missing required field "TemplateVersion.name"`)}
	}
	if v, ok := tvc.mutation.Name(); ok {
		if err := templateversion.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "TemplateVersion.name": %w`, err)}
		}
	}
	if _, ok := tvc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`model: missing required field "TemplateVersion.version"`)}
	}
	if v, ok := tvc.mutation.Version(); ok {
		if err := templateversion.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`model: validator failed for field "TemplateVersion.version": %w`, err)}
		}
	}
	if _, ok := tvc.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New(`model: missing required field "TemplateVersion.source"`)}
	}
	if v, ok := tvc.mutation.Source(); ok {
		if err := templateversion.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`model: validator failed for field "TemplateVersion.source": %w`, err)}
		}
	}
	if _, ok := tvc.mutation.Schema(); !ok {
		return &ValidationError{Name: "schema", err: errors.New(`model: missing required field "TemplateVersion.schema"`)}
	}
	if v, ok := tvc.mutation.Schema(); ok {
		if err := v.Validate(); err != nil {
			return &ValidationError{Name: "schema", err: fmt.Errorf(`model: validator failed for field "TemplateVersion.schema": %w`, err)}
		}
	}
	if _, ok := tvc.mutation.UiSchema(); !ok {
		return &ValidationError{Name: "uiSchema", err: errors.New(`model: missing required field "TemplateVersion.uiSchema"`)}
	}
	if v, ok := tvc.mutation.UiSchema(); ok {
		if err := v.Validate(); err != nil {
			return &ValidationError{Name: "uiSchema", err: fmt.Errorf(`model: validator failed for field "TemplateVersion.uiSchema": %w`, err)}
		}
	}
	if _, ok := tvc.mutation.TemplateID(); !ok {
		return &ValidationError{Name: "template", err: errors.New(`model: missing required edge "TemplateVersion.template"`)}
	}
	return nil
}

func (tvc *TemplateVersionCreate) sqlSave(ctx context.Context) (*TemplateVersion, error) {
	if err := tvc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tvc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tvc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*object.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	tvc.mutation.id = &_node.ID
	tvc.mutation.done = true
	return _node, nil
}

func (tvc *TemplateVersionCreate) createSpec() (*TemplateVersion, *sqlgraph.CreateSpec) {
	var (
		_node = &TemplateVersion{config: tvc.config}
		_spec = sqlgraph.NewCreateSpec(templateversion.Table, sqlgraph.NewFieldSpec(templateversion.FieldID, field.TypeString))
	)
	_spec.Schema = tvc.schemaConfig.TemplateVersion
	_spec.OnConflict = tvc.conflict
	if id, ok := tvc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tvc.mutation.CreateTime(); ok {
		_spec.SetField(templateversion.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := tvc.mutation.UpdateTime(); ok {
		_spec.SetField(templateversion.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = &value
	}
	if value, ok := tvc.mutation.Name(); ok {
		_spec.SetField(templateversion.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tvc.mutation.Version(); ok {
		_spec.SetField(templateversion.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := tvc.mutation.Source(); ok {
		_spec.SetField(templateversion.FieldSource, field.TypeString, value)
		_node.Source = value
	}
	if value, ok := tvc.mutation.Schema(); ok {
		_spec.SetField(templateversion.FieldSchema, field.TypeJSON, value)
		_node.Schema = value
	}
	if value, ok := tvc.mutation.UiSchema(); ok {
		_spec.SetField(templateversion.FieldUiSchema, field.TypeJSON, value)
		_node.UiSchema = value
	}
	if nodes := tvc.mutation.TemplateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   templateversion.TemplateTable,
			Columns: []string{templateversion.TemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(template.FieldID, field.TypeString),
			},
		}
		edge.Schema = tvc.schemaConfig.TemplateVersion
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TemplateID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tvc.mutation.ResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   templateversion.ResourcesTable,
			Columns: []string{templateversion.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeString),
			},
		}
		edge.Schema = tvc.schemaConfig.Resource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tvc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   templateversion.ProjectTable,
			Columns: []string{templateversion.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeString),
			},
		}
		edge.Schema = tvc.schemaConfig.TemplateVersion
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProjectID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For required fields, Set calls directly.
//
// For optional fields, Set calls if the value is not zero.
//
// For example:
//
//	## Required
//
//	db.SetX(obj.X)
//
//	## Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (tvc *TemplateVersionCreate) Set(obj *TemplateVersion) *TemplateVersionCreate {
	// Required.
	tvc.SetTemplateID(obj.TemplateID)
	tvc.SetName(obj.Name)
	tvc.SetVersion(obj.Version)
	tvc.SetSource(obj.Source)
	tvc.SetSchema(obj.Schema)
	tvc.SetUiSchema(obj.UiSchema)

	// Optional.
	if obj.CreateTime != nil {
		tvc.SetCreateTime(*obj.CreateTime)
	}
	if obj.UpdateTime != nil {
		tvc.SetUpdateTime(*obj.UpdateTime)
	}
	if obj.ProjectID != "" {
		tvc.SetProjectID(obj.ProjectID)
	}

	// Record the given object.
	tvc.object = obj

	return tvc
}

// getClientSet returns the ClientSet for the given builder.
func (tvc *TemplateVersionCreate) getClientSet() (mc ClientSet) {
	if _, ok := tvc.config.driver.(*txDriver); ok {
		tx := &Tx{config: tvc.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: tvc.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the TemplateVersion entity,
// which is always good for cascading create operations.
func (tvc *TemplateVersionCreate) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *TemplateVersion) error) (*TemplateVersion, error) {
	obj, err := tvc.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := tvc.getClientSet()
	if tvc.fromUpsert {
		q := mc.TemplateVersions().Query().
			Where(
				templateversion.Name(obj.Name),
				templateversion.Version(obj.Version),
			)
		if obj.ProjectID != "" {
			q.Where(templateversion.ProjectID(obj.ProjectID))
		} else {
			q.Where(templateversion.ProjectIDIsNil())
		}
		obj.ID, err = q.OnlyID(ctx)
		if err != nil {
			return nil, fmt.Errorf("model: failed to query id of TemplateVersion entity: %w", err)
		}
	}

	if x := tvc.object; x != nil {
		if _, set := tvc.mutation.Field(templateversion.FieldTemplateID); set {
			obj.TemplateID = x.TemplateID
		}
		if _, set := tvc.mutation.Field(templateversion.FieldName); set {
			obj.Name = x.Name
		}
		if _, set := tvc.mutation.Field(templateversion.FieldVersion); set {
			obj.Version = x.Version
		}
		if _, set := tvc.mutation.Field(templateversion.FieldSource); set {
			obj.Source = x.Source
		}
		if _, set := tvc.mutation.Field(templateversion.FieldProjectID); set {
			obj.ProjectID = x.ProjectID
		}
		obj.Edges = x.Edges
	}

	for i := range cbs {
		if err = cbs[i](ctx, mc, obj); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (tvc *TemplateVersionCreate) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *TemplateVersion) error) *TemplateVersion {
	obj, err := tvc.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (tvc *TemplateVersionCreate) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *TemplateVersion) error) error {
	_, err := tvc.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (tvc *TemplateVersionCreate) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *TemplateVersion) error) {
	if err := tvc.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Set leverages the TemplateVersionCreate Set method,
// it sets the value by judging the definition of each field within the entire item of the given list.
//
// For required fields, Set calls directly.
//
// For optional fields, Set calls if the value is not zero.
//
// For example:
//
//	## Required
//
//	db.SetX(obj.X)
//
//	## Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (tvcb *TemplateVersionCreateBulk) Set(objs ...*TemplateVersion) *TemplateVersionCreateBulk {
	if len(objs) != 0 {
		client := NewTemplateVersionClient(tvcb.config)

		tvcb.builders = make([]*TemplateVersionCreate, len(objs))
		for i := range objs {
			tvcb.builders[i] = client.Create().Set(objs[i])
		}

		// Record the given objects.
		tvcb.objects = objs
	}

	return tvcb
}

// getClientSet returns the ClientSet for the given builder.
func (tvcb *TemplateVersionCreateBulk) getClientSet() (mc ClientSet) {
	if _, ok := tvcb.config.driver.(*txDriver); ok {
		tx := &Tx{config: tvcb.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: tvcb.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the TemplateVersion entities,
// which is always good for cascading create operations.
func (tvcb *TemplateVersionCreateBulk) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *TemplateVersion) error) ([]*TemplateVersion, error) {
	objs, err := tvcb.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return objs, err
	}

	mc := tvcb.getClientSet()
	if tvcb.fromUpsert {
		for i := range objs {
			obj := objs[i]
			q := mc.TemplateVersions().Query().
				Where(
					templateversion.Name(obj.Name),
					templateversion.Version(obj.Version),
				)
			if obj.ProjectID != "" {
				q.Where(templateversion.ProjectID(obj.ProjectID))
			} else {
				q.Where(templateversion.ProjectIDIsNil())
			}
			objs[i].ID, err = q.OnlyID(ctx)
			if err != nil {
				return nil, fmt.Errorf("model: failed to query id of TemplateVersion entity: %w", err)
			}
		}
	}

	if x := tvcb.objects; x != nil {
		for i := range x {
			if _, set := tvcb.builders[i].mutation.Field(templateversion.FieldTemplateID); set {
				objs[i].TemplateID = x[i].TemplateID
			}
			if _, set := tvcb.builders[i].mutation.Field(templateversion.FieldName); set {
				objs[i].Name = x[i].Name
			}
			if _, set := tvcb.builders[i].mutation.Field(templateversion.FieldVersion); set {
				objs[i].Version = x[i].Version
			}
			if _, set := tvcb.builders[i].mutation.Field(templateversion.FieldSource); set {
				objs[i].Source = x[i].Source
			}
			if _, set := tvcb.builders[i].mutation.Field(templateversion.FieldProjectID); set {
				objs[i].ProjectID = x[i].ProjectID
			}
			objs[i].Edges = x[i].Edges
		}
	}

	for i := range objs {
		for j := range cbs {
			if err = cbs[j](ctx, mc, objs[i]); err != nil {
				return nil, err
			}
		}
	}

	return objs, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (tvcb *TemplateVersionCreateBulk) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *TemplateVersion) error) []*TemplateVersion {
	objs, err := tvcb.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return objs
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (tvcb *TemplateVersionCreateBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *TemplateVersion) error) error {
	_, err := tvcb.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (tvcb *TemplateVersionCreateBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *TemplateVersion) error) {
	if err := tvcb.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *TemplateVersionUpsertOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *TemplateVersion) error) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for TemplateVersionUpsertOne.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *TemplateVersionUpsertOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *TemplateVersion) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *TemplateVersionUpsertBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *TemplateVersion) error) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the TemplateVersionUpsertBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for TemplateVersionUpsertBulk.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *TemplateVersionUpsertBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *TemplateVersion) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TemplateVersion.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TemplateVersionUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (tvc *TemplateVersionCreate) OnConflict(opts ...sql.ConflictOption) *TemplateVersionUpsertOne {
	tvc.conflict = opts
	return &TemplateVersionUpsertOne{
		create: tvc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TemplateVersion.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tvc *TemplateVersionCreate) OnConflictColumns(columns ...string) *TemplateVersionUpsertOne {
	tvc.conflict = append(tvc.conflict, sql.ConflictColumns(columns...))
	return &TemplateVersionUpsertOne{
		create: tvc,
	}
}

type (
	// TemplateVersionUpsertOne is the builder for "upsert"-ing
	//  one TemplateVersion node.
	TemplateVersionUpsertOne struct {
		create *TemplateVersionCreate
	}

	// TemplateVersionUpsert is the "OnConflict" setter.
	TemplateVersionUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *TemplateVersionUpsert) SetUpdateTime(v time.Time) *TemplateVersionUpsert {
	u.Set(templateversion.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *TemplateVersionUpsert) UpdateUpdateTime() *TemplateVersionUpsert {
	u.SetExcluded(templateversion.FieldUpdateTime)
	return u
}

// SetSchema sets the "schema" field.
func (u *TemplateVersionUpsert) SetSchema(v types.Schema) *TemplateVersionUpsert {
	u.Set(templateversion.FieldSchema, v)
	return u
}

// UpdateSchema sets the "schema" field to the value that was provided on create.
func (u *TemplateVersionUpsert) UpdateSchema() *TemplateVersionUpsert {
	u.SetExcluded(templateversion.FieldSchema)
	return u
}

// SetUiSchema sets the "uiSchema" field.
func (u *TemplateVersionUpsert) SetUiSchema(v types.UISchema) *TemplateVersionUpsert {
	u.Set(templateversion.FieldUiSchema, v)
	return u
}

// UpdateUiSchema sets the "uiSchema" field to the value that was provided on create.
func (u *TemplateVersionUpsert) UpdateUiSchema() *TemplateVersionUpsert {
	u.SetExcluded(templateversion.FieldUiSchema)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TemplateVersion.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(templateversion.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TemplateVersionUpsertOne) UpdateNewValues() *TemplateVersionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(templateversion.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(templateversion.FieldCreateTime)
		}
		if _, exists := u.create.mutation.TemplateID(); exists {
			s.SetIgnore(templateversion.FieldTemplateID)
		}
		if _, exists := u.create.mutation.Name(); exists {
			s.SetIgnore(templateversion.FieldName)
		}
		if _, exists := u.create.mutation.Version(); exists {
			s.SetIgnore(templateversion.FieldVersion)
		}
		if _, exists := u.create.mutation.Source(); exists {
			s.SetIgnore(templateversion.FieldSource)
		}
		if _, exists := u.create.mutation.ProjectID(); exists {
			s.SetIgnore(templateversion.FieldProjectID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TemplateVersion.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TemplateVersionUpsertOne) Ignore() *TemplateVersionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TemplateVersionUpsertOne) DoNothing() *TemplateVersionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TemplateVersionCreate.OnConflict
// documentation for more info.
func (u *TemplateVersionUpsertOne) Update(set func(*TemplateVersionUpsert)) *TemplateVersionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TemplateVersionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *TemplateVersionUpsertOne) SetUpdateTime(v time.Time) *TemplateVersionUpsertOne {
	return u.Update(func(s *TemplateVersionUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *TemplateVersionUpsertOne) UpdateUpdateTime() *TemplateVersionUpsertOne {
	return u.Update(func(s *TemplateVersionUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetSchema sets the "schema" field.
func (u *TemplateVersionUpsertOne) SetSchema(v types.Schema) *TemplateVersionUpsertOne {
	return u.Update(func(s *TemplateVersionUpsert) {
		s.SetSchema(v)
	})
}

// UpdateSchema sets the "schema" field to the value that was provided on create.
func (u *TemplateVersionUpsertOne) UpdateSchema() *TemplateVersionUpsertOne {
	return u.Update(func(s *TemplateVersionUpsert) {
		s.UpdateSchema()
	})
}

// SetUiSchema sets the "uiSchema" field.
func (u *TemplateVersionUpsertOne) SetUiSchema(v types.UISchema) *TemplateVersionUpsertOne {
	return u.Update(func(s *TemplateVersionUpsert) {
		s.SetUiSchema(v)
	})
}

// UpdateUiSchema sets the "uiSchema" field to the value that was provided on create.
func (u *TemplateVersionUpsertOne) UpdateUiSchema() *TemplateVersionUpsertOne {
	return u.Update(func(s *TemplateVersionUpsert) {
		s.UpdateUiSchema()
	})
}

// Exec executes the query.
func (u *TemplateVersionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for TemplateVersionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TemplateVersionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TemplateVersionUpsertOne) ID(ctx context.Context) (id object.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: TemplateVersionUpsertOne.ID is not supported by MySQL driver. Use TemplateVersionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TemplateVersionUpsertOne) IDX(ctx context.Context) object.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TemplateVersionCreateBulk is the builder for creating many TemplateVersion entities in bulk.
type TemplateVersionCreateBulk struct {
	config
	builders   []*TemplateVersionCreate
	conflict   []sql.ConflictOption
	objects    []*TemplateVersion
	fromUpsert bool
}

// Save creates the TemplateVersion entities in the database.
func (tvcb *TemplateVersionCreateBulk) Save(ctx context.Context) ([]*TemplateVersion, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tvcb.builders))
	nodes := make([]*TemplateVersion, len(tvcb.builders))
	mutators := make([]Mutator, len(tvcb.builders))
	for i := range tvcb.builders {
		func(i int, root context.Context) {
			builder := tvcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TemplateVersionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tvcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tvcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tvcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tvcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tvcb *TemplateVersionCreateBulk) SaveX(ctx context.Context) []*TemplateVersion {
	v, err := tvcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tvcb *TemplateVersionCreateBulk) Exec(ctx context.Context) error {
	_, err := tvcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tvcb *TemplateVersionCreateBulk) ExecX(ctx context.Context) {
	if err := tvcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TemplateVersion.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TemplateVersionUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (tvcb *TemplateVersionCreateBulk) OnConflict(opts ...sql.ConflictOption) *TemplateVersionUpsertBulk {
	tvcb.conflict = opts
	return &TemplateVersionUpsertBulk{
		create: tvcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TemplateVersion.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tvcb *TemplateVersionCreateBulk) OnConflictColumns(columns ...string) *TemplateVersionUpsertBulk {
	tvcb.conflict = append(tvcb.conflict, sql.ConflictColumns(columns...))
	return &TemplateVersionUpsertBulk{
		create: tvcb,
	}
}

// TemplateVersionUpsertBulk is the builder for "upsert"-ing
// a bulk of TemplateVersion nodes.
type TemplateVersionUpsertBulk struct {
	create *TemplateVersionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TemplateVersion.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(templateversion.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TemplateVersionUpsertBulk) UpdateNewValues() *TemplateVersionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(templateversion.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(templateversion.FieldCreateTime)
			}
			if _, exists := b.mutation.TemplateID(); exists {
				s.SetIgnore(templateversion.FieldTemplateID)
			}
			if _, exists := b.mutation.Name(); exists {
				s.SetIgnore(templateversion.FieldName)
			}
			if _, exists := b.mutation.Version(); exists {
				s.SetIgnore(templateversion.FieldVersion)
			}
			if _, exists := b.mutation.Source(); exists {
				s.SetIgnore(templateversion.FieldSource)
			}
			if _, exists := b.mutation.ProjectID(); exists {
				s.SetIgnore(templateversion.FieldProjectID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TemplateVersion.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TemplateVersionUpsertBulk) Ignore() *TemplateVersionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TemplateVersionUpsertBulk) DoNothing() *TemplateVersionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TemplateVersionCreateBulk.OnConflict
// documentation for more info.
func (u *TemplateVersionUpsertBulk) Update(set func(*TemplateVersionUpsert)) *TemplateVersionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TemplateVersionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *TemplateVersionUpsertBulk) SetUpdateTime(v time.Time) *TemplateVersionUpsertBulk {
	return u.Update(func(s *TemplateVersionUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *TemplateVersionUpsertBulk) UpdateUpdateTime() *TemplateVersionUpsertBulk {
	return u.Update(func(s *TemplateVersionUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetSchema sets the "schema" field.
func (u *TemplateVersionUpsertBulk) SetSchema(v types.Schema) *TemplateVersionUpsertBulk {
	return u.Update(func(s *TemplateVersionUpsert) {
		s.SetSchema(v)
	})
}

// UpdateSchema sets the "schema" field to the value that was provided on create.
func (u *TemplateVersionUpsertBulk) UpdateSchema() *TemplateVersionUpsertBulk {
	return u.Update(func(s *TemplateVersionUpsert) {
		s.UpdateSchema()
	})
}

// SetUiSchema sets the "uiSchema" field.
func (u *TemplateVersionUpsertBulk) SetUiSchema(v types.UISchema) *TemplateVersionUpsertBulk {
	return u.Update(func(s *TemplateVersionUpsert) {
		s.SetUiSchema(v)
	})
}

// UpdateUiSchema sets the "uiSchema" field to the value that was provided on create.
func (u *TemplateVersionUpsertBulk) UpdateUiSchema() *TemplateVersionUpsertBulk {
	return u.Update(func(s *TemplateVersionUpsert) {
		s.UpdateUiSchema()
	})
}

// Exec executes the query.
func (u *TemplateVersionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the TemplateVersionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for TemplateVersionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TemplateVersionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
