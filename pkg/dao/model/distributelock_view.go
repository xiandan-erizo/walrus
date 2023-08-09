// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"context"
	"errors"

	"github.com/seal-io/seal/pkg/dao/model/distributelock"
	"github.com/seal-io/seal/pkg/dao/types/object"
)

// DistributeLockCreateInput holds the creation input of the DistributeLock entity.
type DistributeLockCreateInput struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	ExpireAt int64 `uri:"-" query:"-" json:"expireAt"`
}

// Model returns the DistributeLock entity for creating,
// after validating.
func (dlci *DistributeLockCreateInput) Model() *DistributeLock {
	if dlci == nil {
		return nil
	}

	_dl := &DistributeLock{
		ExpireAt: dlci.ExpireAt,
	}

	return _dl
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (dlci *DistributeLockCreateInput) Load() error {
	if dlci == nil {
		return errors.New("nil receiver")
	}

	return dlci.LoadWith(dlci.inputConfig.Context, dlci.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (dlci *DistributeLockCreateInput) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if dlci == nil {
		return errors.New("nil receiver")
	}

	return nil
}

// DistributeLockCreateInputs holds the creation input item of the DistributeLock entities.
type DistributeLockCreateInputsItem struct {
	ExpireAt int64 `uri:"-" query:"-" json:"expireAt"`
}

// DistributeLockCreateInputs holds the creation input of the DistributeLock entities.
type DistributeLockCreateInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Items []*DistributeLockCreateInputsItem `uri:"-" query:"-" json:"items"`
}

// Model returns the DistributeLock entities for creating,
// after validating.
func (dlci *DistributeLockCreateInputs) Model() []*DistributeLock {
	if dlci == nil || len(dlci.Items) == 0 {
		return nil
	}

	_dls := make([]*DistributeLock, len(dlci.Items))

	for i := range dlci.Items {
		_dl := &DistributeLock{
			ExpireAt: dlci.Items[i].ExpireAt,
		}

		_dls[i] = _dl
	}

	return _dls
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (dlci *DistributeLockCreateInputs) Load() error {
	if dlci == nil {
		return errors.New("nil receiver")
	}

	return dlci.LoadWith(dlci.inputConfig.Context, dlci.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (dlci *DistributeLockCreateInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if dlci == nil {
		return errors.New("nil receiver")
	}

	if len(dlci.Items) == 0 {
		return errors.New("empty items")
	}

	return nil
}

// DistributeLockDeleteInput holds the deletion input of the DistributeLock entity.
type DistributeLockDeleteInput = DistributeLockQueryInput

// DistributeLockDeleteInputs holds the deletion input item of the DistributeLock entities.
type DistributeLockDeleteInputsItem struct {
	ID string `uri:"-" query:"-" json:"id"`
}

// DistributeLockDeleteInputs holds the deletion input of the DistributeLock entities.
type DistributeLockDeleteInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Items []*DistributeLockDeleteInputsItem `uri:"-" query:"-" json:"items"`
}

// Model returns the DistributeLock entities for deleting,
// after validating.
func (dldi *DistributeLockDeleteInputs) Model() []*DistributeLock {
	if dldi == nil || len(dldi.Items) == 0 {
		return nil
	}

	_dls := make([]*DistributeLock, len(dldi.Items))
	for i := range dldi.Items {
		_dls[i] = &DistributeLock{
			ID: dldi.Items[i].ID,
		}
	}
	return _dls
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (dldi *DistributeLockDeleteInputs) Load() error {
	if dldi == nil {
		return errors.New("nil receiver")
	}

	return dldi.LoadWith(dldi.inputConfig.Context, dldi.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (dldi *DistributeLockDeleteInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if dldi == nil {
		return errors.New("nil receiver")
	}

	if len(dldi.Items) == 0 {
		return errors.New("empty items")
	}

	q := cs.DistributeLocks().Query()

	ids := make([]string, 0, len(dldi.Items))

	for i := range dldi.Items {
		if dldi.Items[i] == nil {
			return errors.New("nil item")
		}

		if dldi.Items[i].ID != "" {
			ids = append(ids, dldi.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	idsLen := len(ids)

	idsCnt, err := q.Where(distributelock.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != idsLen {
		return errors.New("found unrecognized item")
	}

	return nil
}

// DistributeLockQueryInput holds the query input of the DistributeLock entity.
type DistributeLockQueryInput struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Refer *object.Refer `uri:"distributelock,default=\"\"" query:"-" json:"-"`
	ID    string        `uri:"id" query:"-" json:"id"` // TODO(thxCode): remove the uri:"id" after supporting hierarchical routes.
}

// Model returns the DistributeLock entity for querying,
// after validating.
func (dlqi *DistributeLockQueryInput) Model() *DistributeLock {
	if dlqi == nil {
		return nil
	}

	return &DistributeLock{
		ID: dlqi.ID,
	}
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (dlqi *DistributeLockQueryInput) Load() error {
	if dlqi == nil {
		return errors.New("nil receiver")
	}

	return dlqi.LoadWith(dlqi.inputConfig.Context, dlqi.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (dlqi *DistributeLockQueryInput) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if dlqi == nil {
		return errors.New("nil receiver")
	}

	if dlqi.Refer != nil && *dlqi.Refer == "" {
		return nil
	}

	q := cs.DistributeLocks().Query()

	if dlqi.Refer != nil {
		if dlqi.Refer.IsString() {
			q.Where(
				distributelock.ID(dlqi.Refer.String()))
		} else {
			return errors.New("invalid identify refer of distributelock")
		}
	} else if dlqi.ID != "" {
		q.Where(
			distributelock.ID(dlqi.ID))
	} else {
		return errors.New("invalid identify of distributelock")
	}

	dlqi.ID, err = q.OnlyID(ctx)
	return err
}

// DistributeLockQueryInputs holds the query input of the DistributeLock entities.
type DistributeLockQueryInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (dlqi *DistributeLockQueryInputs) Load() error {
	if dlqi == nil {
		return errors.New("nil receiver")
	}

	return dlqi.LoadWith(dlqi.inputConfig.Context, dlqi.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (dlqi *DistributeLockQueryInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if dlqi == nil {
		return errors.New("nil receiver")
	}

	return err
}

// DistributeLockUpdateInput holds the modification input of the DistributeLock entity.
type DistributeLockUpdateInput struct {
	DistributeLockQueryInput `uri:",inline" query:"-" json:",inline"`

	ExpireAt int64 `uri:"-" query:"-" json:"expireAt,omitempty"`
}

// Model returns the DistributeLock entity for modifying,
// after validating.
func (dlui *DistributeLockUpdateInput) Model() *DistributeLock {
	if dlui == nil {
		return nil
	}

	_dl := &DistributeLock{
		ID:       dlui.ID,
		ExpireAt: dlui.ExpireAt,
	}

	return _dl
}

// DistributeLockUpdateInputs holds the modification input item of the DistributeLock entities.
type DistributeLockUpdateInputsItem struct {
	ID string `uri:"-" query:"-" json:"id"`

	ExpireAt int64 `uri:"-" query:"-" json:"expireAt,omitempty"`
}

// DistributeLockUpdateInputs holds the modification input of the DistributeLock entities.
type DistributeLockUpdateInputs struct {
	inputConfig `uri:"-" query:"-" json:"-"`

	Items []*DistributeLockUpdateInputsItem `uri:"-" query:"-" json:"items"`
}

// Model returns the DistributeLock entities for modifying,
// after validating.
func (dlui *DistributeLockUpdateInputs) Model() []*DistributeLock {
	if dlui == nil || len(dlui.Items) == 0 {
		return nil
	}

	_dls := make([]*DistributeLock, len(dlui.Items))

	for i := range dlui.Items {
		_dl := &DistributeLock{
			ID:       dlui.Items[i].ID,
			ExpireAt: dlui.Items[i].ExpireAt,
		}

		_dls[i] = _dl
	}

	return _dls
}

// Load checks the input.
// TODO(thxCode): rename to Validate after supporting hierarchical routes.
func (dlui *DistributeLockUpdateInputs) Load() error {
	if dlui == nil {
		return errors.New("nil receiver")
	}

	return dlui.LoadWith(dlui.inputConfig.Context, dlui.inputConfig.ClientSet)
}

// LoadWith checks the input with the given context and client set.
// TODO(thxCode): rename to ValidateWith after supporting hierarchical routes.
func (dlui *DistributeLockUpdateInputs) LoadWith(ctx context.Context, cs ClientSet) (err error) {
	if dlui == nil {
		return errors.New("nil receiver")
	}

	if len(dlui.Items) == 0 {
		return errors.New("empty items")
	}

	q := cs.DistributeLocks().Query()

	ids := make([]string, 0, len(dlui.Items))

	for i := range dlui.Items {
		if dlui.Items[i] == nil {
			return errors.New("nil item")
		}

		if dlui.Items[i].ID != "" {
			ids = append(ids, dlui.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	idsLen := len(ids)

	idsCnt, err := q.Where(distributelock.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != idsLen {
		return errors.New("found unrecognized item")
	}

	return nil
}

// DistributeLockOutput holds the output of the DistributeLock entity.
type DistributeLockOutput struct {
	ID       string `json:"id,omitempty"`
	ExpireAt int64  `json:"expireAt,omitempty"`
}

// View returns the output of DistributeLock.
func (_dl *DistributeLock) View() *DistributeLockOutput {
	return ExposeDistributeLock(_dl)
}

// View returns the output of DistributeLocks.
func (_dls DistributeLocks) View() []*DistributeLockOutput {
	return ExposeDistributeLocks(_dls)
}

// ExposeDistributeLock converts the DistributeLock to DistributeLockOutput.
func ExposeDistributeLock(_dl *DistributeLock) *DistributeLockOutput {
	if _dl == nil {
		return nil
	}

	dlo := &DistributeLockOutput{
		ID:       _dl.ID,
		ExpireAt: _dl.ExpireAt,
	}

	return dlo
}

// ExposeDistributeLocks converts the DistributeLock slice to DistributeLockOutput pointer slice.
func ExposeDistributeLocks(_dls []*DistributeLock) []*DistributeLockOutput {
	if len(_dls) == 0 {
		return nil
	}

	dlos := make([]*DistributeLockOutput, len(_dls))
	for i := range _dls {
		dlos[i] = ExposeDistributeLock(_dls[i])
	}
	return dlos
}
