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
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/stretchr/testify/require"

	"github.com/chenmingyong0423/go-mongox/v2/field"

	"github.com/stretchr/testify/assert"
)

func TestBeforeInsert(t *testing.T) {
	testCases := []struct {
		name        string
		doc         any
		currentTime time.Time
		fields      []*field.Filed

		wantErr      error
		validateFunc func(*testing.T, any)
	}{
		{
			name:        "not reflect.Value",
			doc:         "",
			currentTime: time.Now(),
		},
		{
			name:        "empty struct",
			doc:         reflect.ValueOf(struct{}{}),
			currentTime: time.Now(),
			fields:      []*field.Filed{},
			wantErr:     nil,
		},
		{
			name:        "not inlined struct",
			doc:         reflect.ValueOf(&user{}).Elem(),
			currentTime: time.Now(),
			fields:      field.ParseFields(&user{}),
			wantErr:     nil,
			validateFunc: func(t *testing.T, v any) {
				u, ok := v.(user)
				require.True(t, ok)
				require.NotZero(t, u.ID)
				require.NotZero(t, u.CreatedAt)
				require.NotZero(t, u.UpdatedAt)
				require.NotZero(t, u.CreateSecondTime)
				require.NotZero(t, u.CreateMilliTime)
				require.NotZero(t, u.CreateNanoTime)
				require.NotZero(t, u.UpdateSecondTime)
				require.NotZero(t, u.UpdateMilliTime)
				require.NotZero(t, u.UpdateNanoTime)
			},
		},
		{
			name: "not inlined struct with CreatedAt and UpdateSecondTime",
			doc: reflect.ValueOf(&user{
				CreatedAt:        time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdateSecondTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
			}).Elem(),
			currentTime: time.Now(),
			fields:      field.ParseFields(&user{}),
			wantErr:     nil,
			validateFunc: func(t *testing.T, v any) {
				u, ok := v.(user)
				require.True(t, ok)
				require.NotZero(t, u.ID)
				require.NotZero(t, u.CreatedAt)
				require.NotZero(t, u.UpdatedAt)
				require.NotEqual(t, u.CreatedAt, u.UpdatedAt)
				require.NotZero(t, u.CreateSecondTime)
				require.NotZero(t, u.CreateMilliTime)
				require.NotZero(t, u.CreateNanoTime)
				require.NotZero(t, u.UpdateSecondTime)
				require.NotZero(t, u.UpdateMilliTime)
				require.NotZero(t, u.UpdateNanoTime)
				require.NotEqual(t, u.CreateSecondTime, u.UpdateSecondTime)
			},
		},
		{
			name:        "not inlined struct in mongoxActOnCreatedAtUser",
			doc:         reflect.ValueOf(&mongoxActOnCreatedAtUser{}).Elem(),
			currentTime: time.Now(),
			fields:      field.ParseFields(&mongoxActOnCreatedAtUser{}),
			wantErr:     nil,
			validateFunc: func(t *testing.T, v any) {
				u, ok := v.(mongoxActOnCreatedAtUser)
				require.True(t, ok)
				require.NotZero(t, u.ID)
				require.NotZero(t, u.CreatedAt)
				require.NotZero(t, u.UpdatedAt)
				require.NotZero(t, u.CreateSecondTime)
				require.NotZero(t, u.CreateMilliTime)
				require.NotZero(t, u.CreateNanoTime)
				require.NotZero(t, u.UpdateSecondTime)
				require.NotZero(t, u.UpdateMilliTime)
				require.NotZero(t, u.UpdateNanoTime)
			},
		},
		{
			name: "not inlined struct with CreatedAt and UpdateSecondTime in mongoxActOnCreatedAtUser",
			doc: reflect.ValueOf(&mongoxActOnCreatedAtUser{
				CreatedAt:        time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(),
				UpdateSecondTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
			}).Elem(),
			currentTime: time.Now(),
			fields:      field.ParseFields(&mongoxActOnCreatedAtUser{}),
			wantErr:     nil,
			validateFunc: func(t *testing.T, v any) {
				u, ok := v.(mongoxActOnCreatedAtUser)
				require.True(t, ok)
				require.NotZero(t, u.ID)
				require.NotZero(t, u.CreatedAt)
				require.NotZero(t, u.UpdatedAt)
				require.NotEqual(t, u.CreatedAt, u.UpdatedAt)
				require.NotZero(t, u.CreateSecondTime)
				require.NotZero(t, u.CreateMilliTime)
				require.NotZero(t, u.CreateNanoTime)
				require.NotZero(t, u.UpdateSecondTime)
				require.NotZero(t, u.UpdateMilliTime)
				require.NotZero(t, u.UpdateNanoTime)
				require.NotEqual(t, u.CreateSecondTime, u.UpdateSecondTime)
			},
		},
		{
			name:        "inlined struct",
			doc:         reflect.ValueOf(&inlinedUser{}).Elem(),
			currentTime: time.Now(),
			fields:      field.ParseFields(&inlinedUser{}),
			wantErr:     nil,
			validateFunc: func(t *testing.T, v any) {
				u, ok := v.(inlinedUser)
				require.True(t, ok)
				require.NotZero(t, u.ID)
				require.NotZero(t, u.CreatedAt)
				require.NotZero(t, u.UpdatedAt)
				require.NotZero(t, u.CreateSecondTime)
				require.NotZero(t, u.CreateMilliTime)
				require.NotZero(t, u.CreateNanoTime)
				require.NotZero(t, u.UpdateSecondTime)
				require.NotZero(t, u.UpdateMilliTime)
				require.NotZero(t, u.UpdateNanoTime)
			},
		},
		{
			name:        "inlined pointer struct",
			doc:         reflect.ValueOf(&inlinedUser{}),
			currentTime: time.Now(),
			fields:      field.ParseFields(&inlinedUser{}),
			wantErr:     nil,
			validateFunc: func(t *testing.T, v any) {
				u, ok := v.(*inlinedUser)
				require.True(t, ok)
				require.NotZero(t, u.ID)
				require.NotZero(t, u.CreatedAt)
				require.NotZero(t, u.UpdatedAt)
				require.NotZero(t, u.CreateSecondTime)
				require.NotZero(t, u.CreateMilliTime)
				require.NotZero(t, u.CreateNanoTime)
				require.NotZero(t, u.UpdateSecondTime)
				require.NotZero(t, u.UpdateMilliTime)
				require.NotZero(t, u.UpdateNanoTime)
			},
		},
		{
			name: "int64 / int type for default time field",
			doc: reflect.ValueOf(&struct {
				ID        bson.ObjectID `bson:"_id,omitempty" mongox:"autoID"`
				CreatedAt int64         `bson:"created_at"`
				UpdatedAt int           `bson:"updated_at"`
			}{}),
			currentTime: time.Now(),
			fields: field.ParseFields(struct {
				ID        bson.ObjectID `bson:"_id,omitempty" mongox:"autoID"`
				CreatedAt int64         `bson:"created_at"`
				UpdatedAt int           `bson:"updated_at"`
			}{}),
			wantErr: nil,
			validateFunc: func(t *testing.T, v any) {
				u, ok := v.(*struct {
					ID        bson.ObjectID `bson:"_id,omitempty" mongox:"autoID"`
					CreatedAt int64         `bson:"created_at"`
					UpdatedAt int           `bson:"updated_at"`
				})
				require.True(t, ok)
				require.NotZero(t, u.ID)
				require.NotZero(t, u.CreatedAt)
				require.NotZero(t, u.UpdatedAt)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := beforeInsert(tc.doc, tc.currentTime, tc.fields)
			assert.Equal(t, tc.wantErr, err)
			if tc.validateFunc != nil {
				tc.validateFunc(t, tc.doc.(reflect.Value).Interface())
			}
		})
	}
}

func Test_beforeUpdate(t *testing.T) {

	tests := []struct {
		name        string
		updates     any
		currentTime time.Time
		fields      []*field.Filed
		want        any
		wantErr     error
	}{
		{
			name:    "nil updates",
			updates: nil,
			fields:  []*field.Filed{},
			want:    nil,
		},
		{
			name:    "a bson.M updates without $set",
			updates: bson.M{},
			fields:  []*field.Filed{},
			want:    bson.M{},
		},
		{
			name:    "a bson.M updates but not additional fields",
			updates: bson.M{"$set": bson.M{"name": "Mingyong Chen"}},
			fields:  []*field.Filed{},
			want:    bson.M{"$set": bson.M{"name": "Mingyong Chen"}},
		},
		{
			name:        "a bson.M updates and additional fields",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen"}},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields:      field.ParseFields(&user{}),
			want:        bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), "update_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "update_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}},
		},
		{
			name:        "a bson.M updates with updateTime and additional fields",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)}},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields:      field.ParseFields(&user{}),
			want:        bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC), "update_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "update_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}},
		},
		{
			name:        "a bson.M updates and additional fields",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen"}},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields:      field.ParseFields(&mongoxActOnCreatedAtUser{}),
			want:        bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "update_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}},
		},
		{
			name:        "a bson.M updates with updateTime and additional fields",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)}},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields:      field.ParseFields(&mongoxActOnCreatedAtUser{}),
			want:        bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC), "update_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "update_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}},
		},
		{
			name:        "a bson.M updates and additional-inlined fields",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen"}},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields:      field.ParseFields(&inlinedUser{}),
			want:        bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), "update_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "update_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}},
		},
		{
			name:        "a bson.M updates and not time.Time type for default time field",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen"}},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields: field.ParseFields(struct {
				CreatedAt int64 `bson:"created_at"`
				UpdatedAt int   `bson:"updated_at"`
			}{}),
			want: bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix()}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := beforeUpdate(tt.updates, tt.currentTime, tt.fields)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, tt.updates)
		})
	}
}

func Test_beforeUpsert(t *testing.T) {
	tests := []struct {
		name        string
		updates     any
		currentTime time.Time
		fields      []*field.Filed
		want        any
		wantErr     error
	}{
		{
			name:    "nil updates",
			updates: nil,
			fields:  []*field.Filed{},
			want:    nil,
		},
		{
			name:    "a bson.M updates without $set",
			updates: bson.M{},
			fields:  []*field.Filed{},
			want:    bson.M{},
		},
		{
			name:    "a bson.M updates but not additional fields",
			updates: bson.M{"$set": bson.M{"name": "Mingyong Chen"}},
			fields:  []*field.Filed{},
			want:    bson.M{"$set": bson.M{"name": "Mingyong Chen"}},
		},
		{
			name:        "a bson.M updates with invalid $setOnInsert",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen"}, "$setOnInsert": "invalid"},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields:      field.ParseFields(&user{}),
			want:        bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), "update_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "update_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}, "$setOnInsert": "invalid"},
		},
		{
			name:        "a bson.M updates with invalid $setOnInsert",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen"}, "$setOnInsert": "invalid"},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields:      field.ParseFields(&mongoxActOnCreatedAtUser{}),
			want:        bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "update_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}, "$setOnInsert": "invalid"},
		},
		{
			name:        "a bson.M updates",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen"}},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields:      field.ParseFields(&updatedUser{}),
			want:        bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), "update_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "update_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}, "$setOnInsert": bson.M{"created_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), "create_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "create_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "create_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}},
		},
		{
			name:        "a bson.M updates with additional fields",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)}, "$setOnInsert": bson.M{"created_at": time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)}},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields:      field.ParseFields(&updatedUser{}),
			want:        bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC), "update_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "update_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}, "$setOnInsert": bson.M{"created_at": time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC), "create_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "create_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "create_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}},
		},
		{
			name:        "a bson.M updates and additional-inlined fields",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen"}},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields:      field.ParseFields(&inlinedUpdatedUser{}),
			want:        bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), "update_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "update_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}, "$setOnInsert": bson.M{"created_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), "create_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "create_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "create_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}},
		},
		{
			name:        "a bson.M updates and not time.Time type for default time field",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen"}},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields: field.ParseFields(struct {
				CreatedAt int64 `bson:"created_at"`
				UpdatedAt int   `bson:"updated_at"`
			}{}),
			want: bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix()}, "$setOnInsert": bson.M{"created_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix()}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := beforeUpsert(tt.updates, tt.currentTime, tt.fields)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, tt.updates)
		})
	}
}

func Test_getTimeValue(t *testing.T) {
	t.Run("invalid type", func(t *testing.T) {
		require.Nil(t, getTimeValue(0, time.Time{}))
	})
}
