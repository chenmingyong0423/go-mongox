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

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/stretchr/testify/require"

	"github.com/chenmingyong0423/go-mongox/v2/field"
	"github.com/chenmingyong0423/go-mongox/v2/operation"

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
			name:        "not inlined struct in user2",
			doc:         reflect.ValueOf(&user2{}).Elem(),
			currentTime: time.Now(),
			fields:      field.ParseFields(&user2{}),
			wantErr:     nil,
			validateFunc: func(t *testing.T, v any) {
				u, ok := v.(user2)
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
			name: "not inlined struct with CreatedAt and UpdateSecondTime in user2",
			doc: reflect.ValueOf(&user2{
				CreatedAt:        time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(),
				UpdateSecondTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
			}).Elem(),
			currentTime: time.Now(),
			fields:      field.ParseFields(&user2{}),
			wantErr:     nil,
			validateFunc: func(t *testing.T, v any) {
				u, ok := v.(user2)
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
			name:        "time.Time default fields with explicit mongox tag",
			doc:         reflect.ValueOf(&userWithTaggedDefaultTimeFields{}).Elem(),
			currentTime: time.Now(),
			fields:      field.ParseFields(&userWithTaggedDefaultTimeFields{}),
			wantErr:     nil,
			validateFunc: func(t *testing.T, v any) {
				u, ok := v.(userWithTaggedDefaultTimeFields)
				require.True(t, ok)
				require.NotZero(t, u.CreatedAt)
				require.NotZero(t, u.UpdatedAt)
				require.Zero(t, u.CreateSecondTime)
				require.Zero(t, u.UpdateMilliTime)
				require.NotZero(t, u.CreateMilliTime)
				require.NotZero(t, u.UpdateNanoTime)
				require.Zero(t, u.InvalidTime)
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
			name:    "a bson.M updates without $set and not additional fields",
			updates: bson.M{},
			fields:  []*field.Filed{},
			want:    bson.M{},
		},
		{
			name:        "a bson.M updates without $set",
			updates:     bson.M{},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields:      field.ParseFields(&user{}),
			want:        bson.M{"$set": bson.M{"updated_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), "update_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "update_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}},
		},
		{
			name:    "a bson.M updates with invalid $set",
			updates: bson.M{"$set": "invalid"},
			fields:  field.ParseFields(&user{}),
			want:    bson.M{"$set": "invalid"},
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
			fields:      field.ParseFields(&user2{}),
			want:        bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "update_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}},
		},
		{
			name:        "a bson.M updates with updateTime and additional fields",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)}},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields:      field.ParseFields(&user2{}),
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
			name:        "a bson.M updates ignore precision tags on time.Time fields",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen"}},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields:      field.ParseFields(&userWithTaggedDefaultTimeFields{}),
			want:        bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), "update_nano_time": int(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano())}},
		},
		{
			name:        "a bson.M updates and not time.Time type for default time field",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen"}},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields: field.ParseFields(struct {
				CreatedAt int64 `bson:"created_at"`
				UpdatedAt int   `bson:"updated_at"`
			}{}),
			want: bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": int(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix())}},
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
			name:    "a bson.M updates without $set and not additional fields",
			updates: bson.M{},
			fields:  []*field.Filed{},
			want:    bson.M{},
		},
		{
			name:        "a bson.M updates without $set",
			updates:     bson.M{},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields:      field.ParseFields(&updatedUser{}),
			want:        bson.M{"$set": bson.M{"updated_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), "update_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "update_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "update_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}, "$setOnInsert": bson.M{"created_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), "create_second_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "create_milli_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), "create_nano_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()}},
		},
		{
			name:    "a bson.M updates with invalid $set",
			updates: bson.M{"$set": "invalid"},
			fields:  []*field.Filed{},
			want:    bson.M{"$set": "invalid"},
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
			name:        "a bson.M upserts ignore precision tags on time.Time fields",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen"}},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields:      field.ParseFields(&userWithTaggedDefaultTimeFields{}),
			want: bson.M{
				"$set": bson.M{
					"name":             "Mingyong Chen",
					"updated_at":       time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
					"update_nano_time": int(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()),
				},
				"$setOnInsert": bson.M{
					"created_at":        time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
					"create_milli_time": int(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli()),
				},
			},
		},
		{
			name:        "a bson.M updates and not time.Time type for default time field",
			updates:     bson.M{"$set": bson.M{"name": "Mingyong Chen"}},
			currentTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			fields: field.ParseFields(struct {
				CreatedAt int64 `bson:"created_at"`
				UpdatedAt int   `bson:"updated_at"`
			}{}),
			want: bson.M{"$set": bson.M{"name": "Mingyong Chen", "updated_at": int(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix())}, "$setOnInsert": bson.M{"created_at": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix()}},
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

func Test_setTimeField(t *testing.T) {
	currentTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	type unixSecond int64

	testCases := []struct {
		name        string
		doc         any
		timeType    field.TimeType
		currentTime time.Time
		fieldType   reflect.Type

		validateFunc func(*testing.T, any)
	}{
		{
			name: "skip invalid field type",
			doc: &struct {
				Value string
			}{},
			timeType:    field.UnixMillisecond,
			currentTime: currentTime,
			fieldType:   reflect.TypeFor[string](),
			validateFunc: func(t *testing.T, v any) {
				doc, ok := v.(*struct {
					Value string
				})
				require.True(t, ok)
				require.Zero(t, doc.Value)
			},
		},
		{
			name: "convertible time value",
			doc: &struct {
				Value unixSecond
			}{},
			timeType:    field.UnixSecond,
			currentTime: currentTime,
			fieldType:   reflect.TypeFor[int64](),
			validateFunc: func(t *testing.T, v any) {
				doc, ok := v.(*struct {
					Value unixSecond
				})
				require.True(t, ok)
				require.Equal(t, unixSecond(currentTime.Unix()), doc.Value)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			setTimeField(reflect.ValueOf(tc.doc).Elem().Field(0), tc.timeType, tc.currentTime, tc.fieldType)
			if tc.validateFunc != nil {
				tc.validateFunc(t, tc.doc)
			}
		})
	}
}

func Test_getTimeValue(t *testing.T) {
	currentTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name        string
		timeType    field.TimeType
		currentTime time.Time
		fieldType   reflect.Type
		want        any
		wantOK      bool
	}{
		{
			name:        "time.Time",
			timeType:    field.UnixTime,
			currentTime: currentTime,
			fieldType:   reflect.TypeFor[time.Time](),
			want:        currentTime,
			wantOK:      true,
		},
		{
			name:        "time.Time with timestamp precision",
			timeType:    field.UnixMillisecond,
			currentTime: currentTime,
			fieldType:   reflect.TypeFor[time.Time](),
			wantOK:      false,
		},
		{
			name:        "int milli",
			timeType:    field.UnixMillisecond,
			currentTime: currentTime,
			fieldType:   reflect.TypeFor[int](),
			want:        int(currentTime.UnixMilli()),
			wantOK:      true,
		},
		{
			name:        "int64 nano",
			timeType:    field.UnixNanosecond,
			currentTime: currentTime,
			fieldType:   reflect.TypeFor[int64](),
			want:        currentTime.UnixNano(),
			wantOK:      true,
		},
		{
			name:        "invalid time type",
			timeType:    0,
			currentTime: currentTime,
			fieldType:   reflect.TypeFor[time.Time](),
			wantOK:      false,
		},
		{
			name:        "invalid field type",
			timeType:    field.UnixMillisecond,
			currentTime: currentTime,
			fieldType:   reflect.TypeFor[string](),
			wantOK:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := getTimeValue(tt.timeType, tt.currentTime, tt.fieldType)
			require.Equal(t, tt.wantOK, ok)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestSoftDeleteScope(t *testing.T) {
	type Model struct {
		DeletedAt time.Time `bson:"deleted_at,omitempty"`
	}
	type Doc struct {
		Model `bson:",inline"`
	}
	fields := field.ParseFields(new(Doc))
	guard := bson.D{{Key: "deleted_at", Value: bson.D{{Key: "$in", Value: bson.A{nil, time.Time{}}}}}}
	for _, original := range []any{
		bson.M{"name": "alice", "$or": bson.A{bson.M{"age": 20}, bson.M{"age": 30}}},
		map[string]any{"name": "alice"},
		bson.M{"$and": bson.A{bson.M{"name": "alice"}, bson.M{"age": 18}}},
		map[string]any{"$and": bson.A{bson.M{"name": "alice"}, bson.M{"age": 18}}},
		bson.D{{Key: "deleted_at", Value: bson.M{"$exists": true}}},
		struct {
			Name string `bson:"name"`
		}{"alice"},
		bson.M{},
		bson.D{},
		map[string]any{},
	} {
		before, err := bson.Marshal(original)
		require.NoError(t, err)
		for _, op := range []operation.OpType{operation.OpTypeBeforeFind, operation.OpTypeBeforeUpdate, operation.OpTypeBeforeUpsert, operation.OpTypeBeforeDelete} {
			opCtx := operation.NewOpContext(nil, operation.WithFields(fields), operation.WithFilter(original), operation.WithStartTime(time.Now()))
			require.NoError(t, Execute(context.Background(), opCtx, op))
			require.Equal(t, bson.D{{Key: "$and", Value: bson.A{original, guard}}}, opCtx.Filter)
			// Another phase adds its scope while preserving the existing filter.
			previous := opCtx.Filter
			require.NoError(t, Execute(context.Background(), opCtx, operation.OpTypeBeforeUpdate))
			require.Equal(t, bson.D{{Key: "$and", Value: bson.A{previous, guard}}}, opCtx.Filter)
		}
		after, err := bson.Marshal(original)
		require.NoError(t, err)
		var beforeDoc, afterDoc bson.M
		require.NoError(t, bson.Unmarshal(before, &beforeDoc))
		require.NoError(t, bson.Unmarshal(after, &afterDoc))
		require.Equal(t, beforeDoc, afterDoc)
	}
	require.Equal(t, bson.M{"x": 1}, FilterNotDeleted(bson.M{"x": 1}, nil))
}

func TestSoftDeleteRawFilter(t *testing.T) {
	fields := field.ParseFields(struct {
		DeletedAt time.Time `bson:"deleted_at"`
	}{})
	raw, err := bson.Marshal(bson.M{"name": "alice"})
	require.NoError(t, err)
	for _, filter := range []any{raw, bson.Raw(raw)} {
		encoded, err := bson.Marshal(FilterNotDeleted(filter, fields))
		require.NoError(t, err)
		original := bson.Raw(encoded).Lookup("$and").Array().Index(0).Document()
		require.Equal(t, "alice", original.Lookup("name").StringValue())
	}
}

func TestSoftDeleteTimes(t *testing.T) {
	now := time.Date(2026, 9, 15, 12, 0, 0, 123000000, time.UTC)
	for _, tt := range []struct {
		doc  any
		want any
	}{
		{&struct {
			DeletedAt time.Time `bson:"deleted_at"`
			UpdatedAt time.Time `bson:"updated_at"`
		}{}, now},
		{&struct {
			DeletedAt int       `bson:"deleted_at"`
			UpdatedAt time.Time `bson:"updated_at"`
		}{}, int(now.Unix())},
		{&struct {
			DeletedAt int64     `bson:"deleted_at" mongox:"autoDeleteTime:milli"`
			UpdatedAt time.Time `bson:"updated_at"`
		}{}, now.UnixMilli()},
		{&struct {
			Removed   int64     `bson:"deleted_at" mongox:"autoDeleteTime:nano"`
			UpdatedAt time.Time `bson:"updated_at"`
		}{}, now.UnixNano()},
	} {
		fields := field.ParseFields(tt.doc)
		opCtx := operation.NewOpContext(nil, operation.WithFields(fields), operation.WithFilter(bson.M{}), operation.WithStartTime(now), operation.WithReflectValue(reflect.ValueOf(tt.doc)))
		require.NoError(t, Execute(context.Background(), opCtx, operation.OpTypeBeforeInsert))
		require.True(t, reflect.ValueOf(tt.doc).Elem().Field(0).IsZero(), "insert must not mark a document deleted")
		require.NoError(t, Execute(context.Background(), opCtx, operation.OpTypeBeforeDelete))
		require.Equal(t, bson.M{"$set": bson.M{"deleted_at": tt.want}}, opCtx.Updates)
		require.True(t, opCtx.SoftDelete)
	}
}

func TestDeleteWithoutDeletedAt(t *testing.T) {
	for _, doc := range []any{struct{ Name string }{}, struct {
		UpdatedAt time.Time `bson:"updated_at"`
	}{}, struct{ DeletedAt string }{}, struct {
		DeletedAt time.Time `bson:"-"`
	}{}} {
		opCtx := operation.NewOpContext(nil, operation.WithFields(field.ParseFields(doc)), operation.WithFilter(bson.M{"x": 1}))
		require.NoError(t, Execute(context.Background(), opCtx, operation.OpTypeBeforeDelete))
		require.Nil(t, opCtx.Updates)
		require.False(t, opCtx.SoftDelete)
		require.Equal(t, bson.M{"x": 1}, opCtx.Filter)
	}
	require.NoError(t, Execute(context.Background(), nil, operation.OpTypeBeforeDelete))
}
