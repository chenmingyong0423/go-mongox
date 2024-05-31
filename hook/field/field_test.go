// Copyright 2024 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package field

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/chenmingyong0423/go-mongox/operation"
)

type model struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (m *model) DefaultId() {
	if m.ID.IsZero() {
		m.ID = primitive.NewObjectID()
	}
}

func (m *model) DefaultCreatedAt() {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now().Local()
	}
}

func (m *model) DefaultUpdatedAt() {
	m.UpdatedAt = time.Now().Local()
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
			opType:  operation.OpTypeAfterInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "nil payload",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithDoc((*model)(nil))),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "not pointer",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithDoc(model{})),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "not pointer and not slice",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithDoc(model{})),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "nil pointer",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithDoc((*model)(nil))),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "nil slice",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithDoc(([]*model)(nil))),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "pointer",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithDoc(&model{})),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "slice",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithDoc([]model{{}, {}})),
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

func Test_getPayload(t *testing.T) {

	testCases := []struct {
		name   string
		opCtx  *operation.OpContext
		opType operation.OpType

		want any
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
			opType: operation.OpTypeAfterInsert,
			want:   nil,
		},
		{
			name:   "before insert",
			opCtx:  operation.NewOpContext(nil, operation.WithDoc(&model{})),
			opType: operation.OpTypeBeforeInsert,
			want:   &model{},
		},
		{
			name:   "before update",
			opCtx:  operation.NewOpContext(nil, operation.WithUpdate(&model{})),
			opType: operation.OpTypeBeforeUpdate,
			want:   &model{},
		},
		{
			name:   "before upsert",
			opCtx:  operation.NewOpContext(nil, operation.WithReplacement(&model{})),
			opType: operation.OpTypeBeforeUpsert,
			want:   &model{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := getPayload(tc.opCtx, tc.opType)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_execute(t *testing.T) {
	testCases := []struct {
		name    string
		doc     any
		opType  operation.OpType
		wantErr error
	}{
		{
			name:    "unexpect op type",
			doc:     &model{},
			opType:  operation.OpTypeAfterInsert,
			wantErr: nil,
		},
		{
			name:    "before insert",
			doc:     &model{},
			opType:  operation.OpTypeBeforeInsert,
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := execute(context.TODO(), tc.doc, tc.opType)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
