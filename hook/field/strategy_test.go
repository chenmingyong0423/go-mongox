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
	"testing"

	"github.com/stretchr/testify/require"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/stretchr/testify/assert"
)

func TestBeforeInsert(t *testing.T) {
	testCases := []struct {
		name string
		doc  any

		wantErr      error
		validateFunc func(*testing.T, *model, *customModel)
	}{
		{
			name:    "nil document",
			doc:     nil,
			wantErr: nil,
		},
		{
			name:    "the type not implement DefaultModel and CustomModel",
			doc:     struct{}{},
			wantErr: nil,
		},
		{
			name:    "default model",
			doc:     &model{},
			wantErr: nil,
			validateFunc: func(t *testing.T, defaultModel *model, _ *customModel) {
				assert.NotZero(t, defaultModel.ID)
				assert.NotZero(t, defaultModel.CreatedAt)
				assert.Zero(t, defaultModel.UpdatedAt)
			},
		},
		{
			name:    "custom model",
			doc:     &customModel{},
			wantErr: nil,
			validateFunc: func(t *testing.T, _ *model, customModel *customModel) {
				assert.NotZero(t, customModel.ID)
				assert.NotZero(t, customModel.CreatedAt)
				assert.Zero(t, customModel.UpdatedAt)

			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := BeforeInsert(tc.doc)
			assert.Equal(t, tc.wantErr, err)
			if tc.validateFunc != nil {
				m, _ := tc.doc.(*model)
				cm, _ := tc.doc.(*customModel)
				tc.validateFunc(t, m, cm)
			}
		})
	}
}

func TestBeforeUpdate(t *testing.T) {
	testCases := []struct {
		name string
		doc  any
		opt  []any

		wantErr     error
		wantUpdates func(*model, *customModel) any
	}{
		{
			name:    "nil document",
			doc:     nil,
			wantErr: nil,
		},
		{
			name:    "the type not implement DefaultModel and CustomModel",
			doc:     struct{}{},
			wantErr: nil,
		},
		{
			name:    "default model: nil options",
			doc:     &model{},
			wantErr: nil,
		},
		{
			name:    "default model: empty options",
			doc:     &model{},
			opt:     []any{},
			wantErr: nil,
		},
		{
			name:    "default model: the first option is not a bson.M",
			doc:     &model{},
			opt:     []any{1},
			wantErr: nil,
			wantUpdates: func(_ *model, _ *customModel) any {
				return 1
			},
		},
		{
			name:    "default model: the first option is a bson.M without $set",
			doc:     &model{},
			opt:     []any{bson.M{}},
			wantErr: nil,
			wantUpdates: func(defaultModel *model, _ *customModel) any {
				return bson.M{
					"$set": bson.M{
						"updated_at": defaultModel.UpdatedAt,
					},
				}
			},
		},
		{
			name:    "default model: the first option is a bson.M with invalid $set",
			doc:     &model{},
			opt:     []any{bson.M{"$set": 1}},
			wantErr: nil,
			wantUpdates: func(defaultModel *model, _ *customModel) any {
				return bson.M{
					"$set": 1,
				}
			},
		},
		{
			name:    "default model: the first option is a bson.M with empty $set",
			doc:     &model{},
			opt:     []any{bson.M{"$set": bson.M{}}},
			wantErr: nil,
			wantUpdates: func(defaultModel *model, _ *customModel) any {
				return bson.M{
					"$set": bson.M{
						"updated_at": defaultModel.UpdatedAt,
					},
				}
			},
		},
		{
			name:    "default model: the first option is a bson.M with $set",
			doc:     &model{},
			opt:     []any{bson.M{"$set": bson.M{"name": "Mingyong Chen"}}},
			wantErr: nil,
			wantUpdates: func(defaultModel *model, _ *customModel) any {
				return bson.M{
					"$set": bson.M{
						"name":       "Mingyong Chen",
						"updated_at": defaultModel.UpdatedAt,
					},
				}
			},
		},
		{
			name:    "custom model: nil options",
			doc:     &customModel{},
			wantErr: nil,
		},
		{
			name:    "custom model: empty options",
			doc:     &customModel{},
			opt:     []any{},
			wantErr: nil,
		},
		{
			name:    "custom model: the first option is not a bson.M",
			doc:     &customModel{},
			opt:     []any{1},
			wantErr: nil,
			wantUpdates: func(_ *model, _ *customModel) any {
				return 1
			},
		},
		{
			name:    "custom model: the first option is a bson.M without $set",
			doc:     &customModel{},
			opt:     []any{bson.M{}},
			wantErr: nil,
			wantUpdates: func(_ *model, customModel *customModel) any {
				return bson.M{
					"$set": bson.M{
						"updatedAt": customModel.UpdatedAt,
					},
				}
			},
		},
		{
			name:    "custom model: the first option is a bson.M with invalid $set",
			doc:     &customModel{},
			opt:     []any{bson.M{"$set": 1}},
			wantErr: nil,
			wantUpdates: func(_ *model, customModel *customModel) any {
				return bson.M{
					"$set": 1,
				}
			},
		},
		{
			name:    "custom model: the first option is a bson.M with empty $set",
			doc:     &customModel{},
			opt:     []any{bson.M{"$set": bson.M{}}},
			wantErr: nil,
			wantUpdates: func(_ *model, customModel *customModel) any {
				return bson.M{
					"$set": bson.M{
						"updatedAt": customModel.UpdatedAt,
					},
				}
			},
		},
		{
			name:    "custom model: the first option is a bson.M with $set",
			doc:     &customModel{},
			opt:     []any{bson.M{"$set": bson.M{"name": "Mingyong Chen"}}},
			wantErr: nil,
			wantUpdates: func(_ *model, customModel *customModel) any {
				return bson.M{
					"$set": bson.M{
						"name":      "Mingyong Chen",
						"updatedAt": customModel.UpdatedAt,
					},
				}
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := BeforeUpdate(tc.doc, tc.opt...)
			assert.Equal(t, tc.wantErr, err)
			if len(tc.opt) > 0 {
				m, _ := tc.doc.(*model)
				cm, _ := tc.doc.(*customModel)
				assert.Equal(t, tc.wantUpdates(m, cm), tc.opt[0])
			}
		})
	}
}

func TestBeforeUpsert(t *testing.T) {
	testCases := []struct {
		name string
		doc  any
		opt  []any

		wantErr     error
		wantUpdates func(*model, *customModel) any
	}{
		{
			name:    "nil document",
			doc:     nil,
			wantErr: nil,
		},
		{
			name:    "the type not implement DefaultModel and CustomModel",
			doc:     struct{}{},
			wantErr: nil,
		},
		{
			name:    "default model: nil options",
			doc:     &model{},
			wantErr: nil,
		},
		{
			name:    "default model: empty options",
			doc:     &model{},
			opt:     []any{},
			wantErr: nil,
		},
		{
			name:    "default model: the first option is not a bson.M",
			doc:     &model{},
			opt:     []any{1},
			wantErr: nil,
			wantUpdates: func(_ *model, _ *customModel) any {
				return 1
			},
		},
		{
			name:    "default model: the first option is a bson.M without $set",
			doc:     &model{},
			opt:     []any{bson.M{}},
			wantErr: nil,
			wantUpdates: func(defaultModel *model, _ *customModel) any {
				return bson.M{
					"$set": bson.M{
						"updated_at": defaultModel.UpdatedAt,
					},
					"$setOnInsert": bson.M{
						"created_at": defaultModel.CreatedAt,
						"_id":        defaultModel.ID,
					},
				}
			},
		},
		{
			name:    "default model: the first option is a bson.M with invalid $set",
			doc:     &model{},
			opt:     []any{bson.M{"$set": 1}},
			wantErr: nil,
			wantUpdates: func(_ *model, _ *customModel) any {
				return bson.M{
					"$set": 1,
				}
			},
		},
		{
			name:    "default model: the first option is a bson.M with empty $set",
			doc:     &model{},
			opt:     []any{bson.M{"$set": bson.M{}}},
			wantErr: nil,
			wantUpdates: func(defaultModel *model, _ *customModel) any {
				return bson.M{
					"$set": bson.M{
						"updated_at": defaultModel.UpdatedAt,
					},
					"$setOnInsert": bson.M{
						"created_at": defaultModel.CreatedAt,
						"_id":        defaultModel.ID,
					},
				}
			},
		},
		{
			name:    "default model: the first option is a bson.M with empty $set and invalid $setOnInsert",
			doc:     &model{},
			opt:     []any{bson.M{"$set": bson.M{}, "$setOnInsert": 1}},
			wantErr: nil,
			wantUpdates: func(defaultModel *model, _ *customModel) any {
				return bson.M{"$set": bson.M{}, "$setOnInsert": 1}
			},
		},
		{
			name:    "default model: the first option is a bson.M with $set and empty $setOnInsert",
			doc:     &model{},
			opt:     []any{bson.M{"$set": bson.M{}, "$setOnInsert": bson.M{}}},
			wantErr: nil,
			wantUpdates: func(defaultModel *model, _ *customModel) any {
				return bson.M{
					"$set": bson.M{
						"updated_at": defaultModel.UpdatedAt,
					},
					"$setOnInsert": bson.M{
						"created_at": defaultModel.CreatedAt,
						"_id":        defaultModel.ID,
					},
				}
			},
		},
		{
			name:    "default model: the first option is a bson.M with $set",
			doc:     &model{},
			opt:     []any{bson.M{"$set": bson.M{"name": "Mingyong Chen"}}},
			wantErr: nil,
			wantUpdates: func(defaultModel *model, _ *customModel) any {
				return bson.M{
					"$set": bson.M{
						"updated_at": defaultModel.UpdatedAt,
						"name":       "Mingyong Chen",
					},
					"$setOnInsert": bson.M{
						"created_at": defaultModel.CreatedAt,
						"_id":        defaultModel.ID,
					},
				}
			},
		},
		{
			name:    "custom model: nil options",
			doc:     &customModel{},
			wantErr: nil,
		},
		{
			name:    "custom model: empty options",
			doc:     &customModel{},
			opt:     []any{},
			wantErr: nil,
		},
		{
			name:    "custom model: the first option is not a bson.M",
			doc:     &customModel{},
			opt:     []any{1},
			wantErr: nil,
			wantUpdates: func(_ *model, _ *customModel) any {
				return 1
			},
		},
		{
			name:    "custom model: the first option is a bson.M without $set",
			doc:     &customModel{},
			opt:     []any{bson.M{}},
			wantErr: nil,
			wantUpdates: func(_ *model, customModel *customModel) any {
				return bson.M{
					"$set": bson.M{
						"updatedAt": customModel.UpdatedAt,
					},
					"$setOnInsert": bson.M{
						"createdAt": customModel.CreatedAt,
						"_id":       customModel.ID,
					},
				}
			},
		},
		{
			name:    "custom model: the first option is a bson.M with invalid $set",
			doc:     &customModel{},
			opt:     []any{bson.M{"$set": 1}},
			wantErr: nil,
			wantUpdates: func(_ *model, _ *customModel) any {
				return bson.M{
					"$set": 1,
				}
			},
		},
		{
			name:    "custom model: the first option is a bson.M with empty $set",
			doc:     &customModel{},
			opt:     []any{bson.M{"$set": bson.M{}}},
			wantErr: nil,
			wantUpdates: func(_ *model, customModel *customModel) any {
				return bson.M{
					"$set": bson.M{
						"updatedAt": customModel.UpdatedAt,
					},
					"$setOnInsert": bson.M{
						"createdAt": customModel.CreatedAt,
						"_id":       customModel.ID,
					},
				}
			},
		},
		{
			name:    "custom model: the first option is a bson.M with empty $set and invalid $setOnInsert",
			doc:     &customModel{},
			opt:     []any{bson.M{"$set": bson.M{}, "$setOnInsert": 1}},
			wantErr: nil,
			wantUpdates: func(_ *model, _ *customModel) any {
				return bson.M{"$set": bson.M{}, "$setOnInsert": 1}
			},
		},
		{
			name:    "custom model: the first option is a bson.M with $set and empty $setOnInsert",
			doc:     &customModel{},
			opt:     []any{bson.M{"$set": bson.M{}, "$setOnInsert": bson.M{}}},
			wantErr: nil,
			wantUpdates: func(_ *model, customModel *customModel) any {
				return bson.M{
					"$set": bson.M{
						"updatedAt": customModel.UpdatedAt,
					},
					"$setOnInsert": bson.M{
						"createdAt": customModel.CreatedAt,
						"_id":       customModel.ID,
					},
				}
			},
		},
		{
			name:    "custom model: the first option is a bson.M with $set",
			doc:     &customModel{},
			opt:     []any{bson.M{"$set": bson.M{"name": "Mingyong Chen"}}},
			wantErr: nil,
			wantUpdates: func(_ *model, customModel *customModel) any {
				return bson.M{
					"$set": bson.M{
						"updatedAt": customModel.UpdatedAt,
						"name":      "Mingyong Chen",
					},
					"$setOnInsert": bson.M{
						"createdAt": customModel.CreatedAt,
						"_id":       customModel.ID,
					},
				}
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := BeforeUpsert(tc.doc, tc.opt...)
			assert.Equal(t, tc.wantErr, err)
			if len(tc.opt) > 0 {
				m, _ := tc.doc.(*model)
				cm, _ := tc.doc.(*customModel)
				assert.Equal(t, tc.wantUpdates(m, cm), tc.opt[0])
			}
		})
	}
}

func Test_getField(t *testing.T) {
	t.Run("supplement test", func(t *testing.T) {
		require.Equal(t, field{}, getField("", nil, nil))
	})
}
