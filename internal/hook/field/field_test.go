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
	"reflect"
	"testing"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/operation"
	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type model struct {
	ID        bson.ObjectID `bson:"_id,omitempty" mongox:"autoID"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
	DeletedAt time.Time     `bson:"deleted_at,omitempty"`
}

type inlinedUser struct {
	model            `bson:",inline"`
	Name             string `bson:"name"`
	CreateSecondTime int64  `bson:"create_second_time" mongox:"autoCreateTime:second"`
	UpdateSecondTime int64  `bson:"update_second_time" mongox:"autoUpdateTime:second"`
	CreateMilliTime  int64  `bson:"create_milli_time" mongox:"autoCreateTime:milli"`
	UpdateMilliTime  int64  `bson:"update_milli_time" mongox:"autoUpdateTime:milli"`
	CreateNanoTime   int64  `bson:"create_nano_time" mongox:"autoCreateTime:nano"`
	UpdateNanoTime   int64  `bson:"update_nano_time" mongox:"autoUpdateTime:nano"`
}

type user struct {
	ID                  bson.ObjectID `bson:"_id,omitempty" mongox:"autoID"`
	CreatedAt           time.Time     `bson:"created_at"`
	UpdatedAt           time.Time     `bson:"updated_at"`
	DeletedAt           time.Time     `bson:"deleted_at,omitempty"`
	Name                string        `bson:"name"`
	CreateSecondTime    int64         `bson:"create_second_time" mongox:"autoCreateTime:second"`
	UpdateSecondTime    int64         `bson:"update_second_time" mongox:"autoUpdateTime:second"`
	CreateMilliTime     int64         `bson:"create_milli_time" mongox:"autoCreateTime:milli"`
	UpdateMilliTime     int64         `bson:"update_milli_time" mongox:"autoUpdateTime:milli"`
	CreateNanoTime      int64         `bson:"create_nano_time" mongox:"autoCreateTime:nano"`
	UpdateNanoTime      int64         `bson:"update_nano_time" mongox:"autoUpdateTime:nano"`
	InvalidTimeTagField time.Time     `bson:"invalid_time_tag_field" mongox:"autoCreateTime:time"`
}

type updatedUser struct {
	CreatedAt           time.Time `bson:"created_at"`
	UpdatedAt           time.Time `bson:"updated_at"`
	DeletedAt           time.Time `bson:"deleted_at,omitempty"`
	Name                string    `bson:"name"`
	CreateSecondTime    int64     `bson:"create_second_time" mongox:"autoCreateTime:second"`
	UpdateSecondTime    int64     `bson:"update_second_time" mongox:"autoUpdateTime:second"`
	CreateMilliTime     int64     `bson:"create_milli_time" mongox:"autoCreateTime:milli"`
	UpdateMilliTime     int64     `bson:"update_milli_time" mongox:"autoUpdateTime:milli"`
	CreateNanoTime      int64     `bson:"create_nano_time" mongox:"autoCreateTime:nano"`
	UpdateNanoTime      int64     `bson:"update_nano_time" mongox:"autoUpdateTime:nano"`
	InvalidTimeTagField time.Time `bson:"invalid_time_tag_field" mongox:"autoCreateTime:time"`
}

type inlinedUpdatedUser struct {
	updatedModel     `bson:",inline"`
	Name             string `bson:"name"`
	CreateSecondTime int64  `bson:"create_second_time" mongox:"autoCreateTime:second"`
	UpdateSecondTime int64  `bson:"update_second_time" mongox:"autoUpdateTime:second"`
	CreateMilliTime  int64  `bson:"create_milli_time" mongox:"autoCreateTime:milli"`
	UpdateMilliTime  int64  `bson:"update_milli_time" mongox:"autoUpdateTime:milli"`
	CreateNanoTime   int64  `bson:"create_nano_time" mongox:"autoCreateTime:nano"`
	UpdateNanoTime   int64  `bson:"update_nano_time" mongox:"autoUpdateTime:nano"`
}

type updatedModel struct {
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	DeletedAt time.Time `bson:"deleted_at,omitempty"`
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
			name:    "nil reflect value",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithReflectValue(reflect.ValueOf(nil))),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "not pointer",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithReflectValue(reflect.ValueOf(model{}))),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "nil pointer",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithReflectValue(reflect.ValueOf((*model)(nil)))),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "nil slice",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithReflectValue(reflect.ValueOf(([]*model)(nil)))),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "pointer - beforeInsert",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithReflectValue(reflect.ValueOf(&model{}))),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "slice - beforeInsert",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithReflectValue(reflect.ValueOf([]model{{}, {}}))),
			opType:  operation.OpTypeBeforeInsert,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "beforeUpdate",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithUpdates(bson.M{})),
			opType:  operation.OpTypeBeforeUpdate,
			opts:    nil,
			wantErr: nil,
		},
		{
			name:    "beforeUpsert",
			ctx:     context.Background(),
			opCtx:   operation.NewOpContext(nil, operation.WithUpdates(bson.M{})),
			opType:  operation.OpTypeBeforeUpsert,
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
