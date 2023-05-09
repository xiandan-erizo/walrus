// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"time"

	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/dao/types/oid"
)

// ModuleVersionQueryInput is the input for the ModuleVersion query.
type ModuleVersionQueryInput struct {
	// ID holds the value of the "id" field.
	ID oid.ID `uri:"id,omitempty" json:"id,omitempty"`
}

// Model converts the ModuleVersionQueryInput to ModuleVersion.
func (in ModuleVersionQueryInput) Model() *ModuleVersion {
	return &ModuleVersion{
		ID: in.ID,
	}
}

// ModuleVersionCreateInput is the input for the ModuleVersion creation.
type ModuleVersionCreateInput struct {
	// Module version.
	Version string `json:"version"`
	// Module version source.
	Source string `json:"source"`
	// Schema of the module.
	Schema *types.ModuleSchema `json:"schema,omitempty"`
	// Module holds the value of the module edge.
	Module ModuleQueryInput `json:"module"`
}

// Model converts the ModuleVersionCreateInput to ModuleVersion.
func (in ModuleVersionCreateInput) Model() *ModuleVersion {
	var entity = &ModuleVersion{
		Version: in.Version,
		Source:  in.Source,
		Schema:  in.Schema,
	}
	entity.ModuleID = in.Module.ID
	return entity
}

// ModuleVersionUpdateInput is the input for the ModuleVersion modification.
type ModuleVersionUpdateInput struct {
	// ID holds the value of the "id" field.
	ID oid.ID `uri:"id" json:"-"`
	// Schema of the module.
	Schema *types.ModuleSchema `json:"schema,omitempty"`
}

// Model converts the ModuleVersionUpdateInput to ModuleVersion.
func (in ModuleVersionUpdateInput) Model() *ModuleVersion {
	var entity = &ModuleVersion{
		ID:     in.ID,
		Schema: in.Schema,
	}
	return entity
}

// ModuleVersionOutput is the output for the ModuleVersion.
type ModuleVersionOutput struct {
	// ID holds the value of the "id" field.
	ID oid.ID `json:"id,omitempty"`
	// Describe creation time.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Describe modification time.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	// Module version.
	Version string `json:"version,omitempty"`
	// Module version source.
	Source string `json:"source,omitempty"`
	// Schema of the module.
	Schema *types.ModuleSchema `json:"schema,omitempty"`
	// Module holds the value of the module edge.
	Module *ModuleOutput `json:"module,omitempty"`
}

// ExposeModuleVersion converts the ModuleVersion to ModuleVersionOutput.
func ExposeModuleVersion(in *ModuleVersion) *ModuleVersionOutput {
	if in == nil {
		return nil
	}
	var entity = &ModuleVersionOutput{
		ID:         in.ID,
		CreateTime: in.CreateTime,
		UpdateTime: in.UpdateTime,
		Version:    in.Version,
		Source:     in.Source,
		Schema:     in.Schema,
		Module:     ExposeModule(in.Edges.Module),
	}
	if in.ModuleID != "" {
		if entity.Module == nil {
			entity.Module = &ModuleOutput{}
		}
		entity.Module.ID = in.ModuleID
	}
	return entity
}

// ExposeModuleVersions converts the ModuleVersion slice to ModuleVersionOutput pointer slice.
func ExposeModuleVersions(in []*ModuleVersion) []*ModuleVersionOutput {
	var out = make([]*ModuleVersionOutput, 0, len(in))
	for i := 0; i < len(in); i++ {
		var o = ExposeModuleVersion(in[i])
		if o == nil {
			continue
		}
		out = append(out, o)
	}
	if len(out) == 0 {
		return nil
	}
	return out
}
