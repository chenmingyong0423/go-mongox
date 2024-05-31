// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"context"
	"reflect"
	"testing"

	"github.com/chenmingyong0423/go-mongox/operation"
	"github.com/stretchr/testify/assert"
)

type entity struct {
	beforeInsert int
	afterInsert  int
	beforeUpdate int
	afterUpdate  int
	beforeUpsert int
	afterUpsert  int
	afterFind    int
}

func (m *entity) BeforeInsert(_ context.Context) error {
	m.beforeInsert++
	return nil
}

func (m *entity) AfterInsert(_ context.Context) error {
	m.afterInsert++
	return nil
}

func (m *entity) BeforeUpdate(_ context.Context) error {
	m.beforeUpdate++
	return nil
}

func (m *entity) AfterUpdate(_ context.Context) error {
	m.afterUpdate++
	return nil
}

func (m *entity) BeforeUpsert(_ context.Context) error {
	m.beforeUpsert++
	return nil
}

func (m *entity) AfterUpsert(_ context.Context) error {
	m.afterUpsert++
	return nil
}

func (m *entity) AfterFind(_ context.Context) error {
	m.afterFind++
	return nil
}

func Test_getPayload(t *testing.T) {
	testCases := []struct {
		name   string
		opCtx  *operation.OpContext
		opType operation.OpType
		want   any
	}{
		{
			name:   "nil opCtx",
			opCtx:  nil,
			opType: operation.OpTypeBeforeInsert,
			want:   nil,
		},
		{
			name:   "empty opCtx",
			opCtx:  operation.NewOpContext(nil),
			opType: operation.OpTypeBeforeInsert,
			want:   nil,
		},
		{
			name:   "unexpect op type",
			opCtx:  operation.NewOpContext(nil),
			opType: operation.OpTypeBeforeFind,
			want:   nil,
		},
		{
			name:   "before insert",
			opCtx:  operation.NewOpContext(nil, operation.WithDoc(&entity{})),
			opType: operation.OpTypeBeforeInsert,
			want:   &entity{},
		},
		{
			name:   "after insert",
			opCtx:  operation.NewOpContext(nil, operation.WithDoc(&entity{})),
			opType: operation.OpTypeAfterInsert,
			want:   &entity{},
		},
		{
			name:   "before upsert",
			opCtx:  operation.NewOpContext(nil, operation.WithReplacement(&entity{})),
			opType: operation.OpTypeBeforeUpsert,
			want:   &entity{},
		},
		{
			name:   "after upsert",
			opCtx:  operation.NewOpContext(nil, operation.WithReplacement(&entity{})),
			opType: operation.OpTypeAfterUpsert,
			want:   &entity{},
		},
		{
			name:   "after find",
			opCtx:  operation.NewOpContext(nil, operation.WithDoc(&entity{})),
			opType: operation.OpTypeAfterFind,
			want:   &entity{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := getPayload(tc.opCtx, tc.opType)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestExecute(t *testing.T) {
	testCases := []struct {
		name   string
		ctx    context.Context
		opCtx  *operation.OpContext
		opType operation.OpType
		opts   []any

		wantErr error
	}{
		{
			name:    "unexpect op type",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil),
			opType:  operation.OpTypeBeforeFind,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "nil payload",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithDoc((*entity)(nil))),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "not pointer",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithDoc(entity{})),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "nil pointer",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithDoc((*entity)(nil))),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "nil slice",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithDoc(([]*entity)(nil))),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "pointer",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithDoc(&entity{})),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "slice",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithDoc([]*entity{{}, {}})),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Execute(tc.ctx, tc.opCtx, tc.opType, tc.opts...)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func Test_execute(t *testing.T) {

	testCases := []struct {
		name   string
		ctx    context.Context
		doc    any
		opType operation.OpType

		want    any
		wantErr error
	}{
		{
			name:   "not implement BeforeAfter interface",
			ctx:    context.Background(),
			doc:    "",
			opType: operation.OpTypeBeforeInsert,

			want:    "",
			wantErr: nil,
		},
		{
			name:   "nil doc",
			ctx:    context.Background(),
			doc:    nil,
			opType: operation.OpTypeBeforeInsert,

			want:    nil,
			wantErr: nil,
		},
		{
			name:   "before insert",
			ctx:    context.Background(),
			doc:    &entity{},
			opType: operation.OpTypeBeforeInsert,

			want:    &entity{beforeInsert: 1},
			wantErr: nil,
		},
		{
			name:   "after insert",
			ctx:    context.Background(),
			doc:    &entity{},
			opType: operation.OpTypeAfterInsert,

			want:    &entity{afterInsert: 1},
			wantErr: nil,
		},
		{
			name:   "before upsert",
			ctx:    context.Background(),
			doc:    &entity{},
			opType: operation.OpTypeBeforeUpsert,

			want:    &entity{beforeUpsert: 1},
			wantErr: nil,
		},
		{
			name:   "after upsert",
			ctx:    context.Background(),
			doc:    &entity{},
			opType: operation.OpTypeAfterUpsert,

			want:    &entity{afterUpsert: 1},
			wantErr: nil,
		},
		{
			name:   "after find",
			ctx:    context.Background(),
			doc:    &entity{},
			opType: operation.OpTypeAfterFind,

			want:    &entity{afterFind: 1},
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := execute(tc.ctx, tc.doc, tc.opType)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, tc.doc)
		})
	}
}

func Test_executeSlice(t *testing.T) {
	testCases := []struct {
		name   string
		ctx    context.Context
		docs   reflect.Value
		opType operation.OpType
		opts   []any

		want    any
		wantErr error
	}{
		{
			name:   "not implement BeforeAfter interface",
			ctx:    context.Background(),
			docs:   reflect.ValueOf([]string{"", ""}),
			opType: operation.OpTypeBeforeInsert,
			opts:   nil,

			want:    []string{"", ""},
			wantErr: nil,
		},
		{
			name:   "before insert",
			ctx:    context.Background(),
			docs:   reflect.ValueOf([]*entity{{}, {}}),
			opType: operation.OpTypeBeforeInsert,
			opts:   nil,

			want:    []*entity{{beforeInsert: 1}, {beforeInsert: 1}},
			wantErr: nil,
		},
		{
			name:    "after insert",
			ctx:     context.Background(),
			docs:    reflect.ValueOf([]*entity{{}, {}}),
			opType:  operation.OpTypeAfterInsert,
			opts:    nil,
			want:    []*entity{{afterInsert: 1}, {afterInsert: 1}},
			wantErr: nil,
		},
		{
			name:   "before upsert",
			ctx:    context.Background(),
			docs:   reflect.ValueOf([]*entity{{}, {}}),
			opType: operation.OpTypeBeforeUpsert,
			opts:   nil,

			want:    []*entity{{beforeUpsert: 1}, {beforeUpsert: 1}},
			wantErr: nil,
		},
		{
			name:    "after upsert",
			ctx:     context.Background(),
			docs:    reflect.ValueOf([]*entity{{}, {}}),
			opType:  operation.OpTypeAfterUpsert,
			opts:    nil,
			want:    []*entity{{afterUpsert: 1}, {afterUpsert: 1}},
			wantErr: nil,
		},
		{
			name:   "after find",
			ctx:    context.Background(),
			docs:   reflect.ValueOf([]*entity{{}, {}}),
			opType: operation.OpTypeAfterFind,
			opts:   nil,

			want:    []*entity{{afterFind: 1}, {afterFind: 1}},
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := executeSlice(tc.ctx, tc.docs, tc.opType, tc.opts...)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.want, tc.docs.Interface())
		})
	}
}
