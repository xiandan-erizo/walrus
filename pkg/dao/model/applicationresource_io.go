// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/dao/types/oid"
)

// ApplicationResourceQueryInput is the input for the ApplicationResource query.
type ApplicationResourceQueryInput struct {
	// ID holds the value of the "id" field.
	ID oid.ID `uri:"id,omitempty" json:"id,omitempty"`
}

// Model converts the ApplicationResourceQueryInput to ApplicationResource.
func (in ApplicationResourceQueryInput) Model() *ApplicationResource {
	return &ApplicationResource{
		ID: in.ID,
	}
}

// ApplicationResourceCreateInput is the input for the ApplicationResource creation.
type ApplicationResourceCreateInput struct {
	// Name of the module that generates the resource.
	Module string `json:"module"`
	// Mode that manages the generated resource, it is the management way of the deployer to the resource, which provides by deployer.
	Mode string `json:"mode"`
	// Type of the generated resource, it is the type of the resource which the deployer observes, which provides by deployer.
	Type string `json:"type"`
	// Name of the generated resource, it is the real identifier of the resource, which provides by deployer.
	Name string `json:"name"`
	// Type of deployer.
	DeployerType string `json:"deployerType"`
	// Status of the resource.
	Status types.ApplicationResourceStatus `json:"status,omitempty"`
	// Application instance to which the resource belongs.
	Instance ApplicationInstanceQueryInput `json:"instance"`
	// Connector to which the resource deploys.
	Connector ConnectorQueryInput `json:"connector"`
	// Application resource to which the resource makes up.
	Composition *ApplicationResourceQueryInput `json:"composition,omitempty"`
}

// Model converts the ApplicationResourceCreateInput to ApplicationResource.
func (in ApplicationResourceCreateInput) Model() *ApplicationResource {
	var entity = &ApplicationResource{
		Module:       in.Module,
		Mode:         in.Mode,
		Type:         in.Type,
		Name:         in.Name,
		DeployerType: in.DeployerType,
		Status:       in.Status,
	}
	entity.InstanceID = in.Instance.ID
	entity.ConnectorID = in.Connector.ID
	if in.Composition != nil {
		entity.CompositionID = in.Composition.ID
	}
	return entity
}

// ApplicationResourceUpdateInput is the input for the ApplicationResource modification.
type ApplicationResourceUpdateInput struct {
	// ID holds the value of the "id" field.
	ID oid.ID `uri:"id" json:"-"`
	// Status of the resource.
	Status types.ApplicationResourceStatus `json:"status,omitempty"`
}

// Model converts the ApplicationResourceUpdateInput to ApplicationResource.
func (in ApplicationResourceUpdateInput) Model() *ApplicationResource {
	var entity = &ApplicationResource{
		ID:     in.ID,
		Status: in.Status,
	}
	return entity
}

// ApplicationResourceOutput is the output for the ApplicationResource.
type ApplicationResourceOutput struct {
	// ID holds the value of the "id" field.
	ID oid.ID `json:"id,omitempty"`
	// Describe creation time.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Describe modification time.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	// Name of the module that generates the resource.
	Module string `json:"module,omitempty"`
	// Mode that manages the generated resource, it is the management way of the deployer to the resource, which provides by deployer.
	Mode string `json:"mode,omitempty"`
	// Type of the generated resource, it is the type of the resource which the deployer observes, which provides by deployer.
	Type string `json:"type,omitempty"`
	// Name of the generated resource, it is the real identifier of the resource, which provides by deployer.
	Name string `json:"name,omitempty"`
	// Type of deployer.
	DeployerType string `json:"deployerType,omitempty"`
	// Status of the resource.
	Status types.ApplicationResourceStatus `json:"status,omitempty"`
	// Application instance to which the resource belongs.
	Instance *ApplicationInstanceOutput `json:"instance,omitempty"`
	// Connector to which the resource deploys.
	Connector *ConnectorOutput `json:"connector,omitempty"`
	// Application resource to which the resource makes up.
	Composition *ApplicationResourceOutput `json:"composition,omitempty"`
	// Application resources that make up this resource.
	Components []*ApplicationResourceOutput `json:"components,omitempty"`
}

// Normalize normalizes the fields,
// and impacts the output.
func (in *ApplicationResourceOutput) Normalize() *ApplicationResourceOutput {
	if in.Module != "" {
		// show application module name only.
		in.Module = strings.Split(in.Module, "/")[0]
	}
	return in
}

// MarshalJSON implements the json.Marshaler interface.
func (in *ApplicationResourceOutput) MarshalJSON() ([]byte, error) {
	type Alias ApplicationResourceOutput
	return json.Marshal((*Alias)(in.Normalize()))
}

// ExposeApplicationResource converts the ApplicationResource to ApplicationResourceOutput.
func ExposeApplicationResource(in *ApplicationResource) *ApplicationResourceOutput {
	if in == nil {
		return nil
	}
	var entity = &ApplicationResourceOutput{
		ID:           in.ID,
		CreateTime:   in.CreateTime,
		UpdateTime:   in.UpdateTime,
		Module:       in.Module,
		Mode:         in.Mode,
		Type:         in.Type,
		Name:         in.Name,
		DeployerType: in.DeployerType,
		Status:       in.Status,
		Instance:     ExposeApplicationInstance(in.Edges.Instance),
		Connector:    ExposeConnector(in.Edges.Connector),
		Composition:  ExposeApplicationResource(in.Edges.Composition),
		Components:   ExposeApplicationResources(in.Edges.Components),
	}
	if in.InstanceID != "" {
		if entity.Instance == nil {
			entity.Instance = &ApplicationInstanceOutput{}
		}
		entity.Instance.ID = in.InstanceID
	}
	if in.ConnectorID != "" {
		if entity.Connector == nil {
			entity.Connector = &ConnectorOutput{}
		}
		entity.Connector.ID = in.ConnectorID
	}
	if in.CompositionID != "" {
		if entity.Composition == nil {
			entity.Composition = &ApplicationResourceOutput{}
		}
		entity.Composition.ID = in.CompositionID
	}
	return entity
}

// ExposeApplicationResources converts the ApplicationResource slice to ApplicationResourceOutput pointer slice.
func ExposeApplicationResources(in []*ApplicationResource) []*ApplicationResourceOutput {
	var out = make([]*ApplicationResourceOutput, 0, len(in))
	for i := 0; i < len(in); i++ {
		var o = ExposeApplicationResource(in[i])
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
