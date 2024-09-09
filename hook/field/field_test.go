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

func (m *model) DefaultId() primitive.ObjectID {
	if m.ID.IsZero() {
		m.ID = primitive.NewObjectID()
	}
	return m.ID
}

func (m *model) DefaultCreatedAt() time.Time {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now().Local()
	}
	return m.CreatedAt
}

func (m *model) DefaultUpdatedAt() time.Time {
	m.UpdatedAt = time.Now().Local()
	return m.UpdatedAt
}

type customModel struct {
	ID        string `bson:"_id"`
	CreatedAt int64  `bson:"createdAt"`
	UpdatedAt int64  `bson:"updatedAt"`
}

func (c *customModel) CustomID() (string, any) {
	if c.ID == "" {
		c.ID = primitive.NewObjectID().Hex()
	}
	return "_id", c.ID
}

func (c *customModel) CustomCreatedAt() (string, any) {
	if c.CreatedAt == 0 {
		c.CreatedAt = time.Now().Unix()
	}
	return "createdAt", c.CreatedAt
}

func (c *customModel) CustomUpdatedAt() (string, any) {
	c.UpdatedAt = time.Now().Unix()
	return "updatedAt", c.UpdatedAt
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
