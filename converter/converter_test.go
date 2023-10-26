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

package converter

import (
	"testing"

	"github.com/chenmingyong0423/go-mongox/internal/types"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestStructToSetBson(t *testing.T) {
	testCases := []struct {
		name  string
		data  any
		wantD bson.D
	}{
		{
			name: "nil data",
			data: nil,

			wantD: nil,
		},
		{
			name: "struct with zero-value",
			data: types.UpdatedUser{},
			wantD: bson.D{
				bson.E{Key: types.Set, Value: bson.D{
					bson.E{Key: "name", Value: ""},
					bson.E{Key: "age", Value: int64(0)},
				}},
			},
		},
		{
			name: "struct with no zero-value",
			data: types.UpdatedUser{Name: "cmy", Age: int64(24)},
			wantD: bson.D{
				bson.E{Key: types.Set, Value: bson.D{
					bson.E{Key: "name", Value: "cmy"},
					bson.E{Key: "age", Value: int64(24)},
				}},
			},
		},
		{
			name: "struct pointer with empty-value",
			data: &types.UpdatedUser{},
			wantD: bson.D{
				bson.E{Key: types.Set, Value: bson.D{
					bson.E{Key: "name", Value: ""},
					bson.E{Key: "age", Value: int64(0)},
				}},
			},
		},
		{
			name: "struct pointer with no empty-value",
			data: &types.UpdatedUser{Name: "cmy", Age: int64(24)},
			wantD: bson.D{
				bson.E{Key: types.Set, Value: bson.D{
					bson.E{Key: "name", Value: "cmy"},
					bson.E{Key: "age", Value: int64(24)},
				}},
			},
		},
		{
			name:  "not struct",
			data:  1,
			wantD: nil,
		},
		{
			name: "empty struct",
			data: struct{}{},

			wantD: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := StructToSetBson(tc.data)
			if len(d) > 0 {
				assert.Equal(t, tc.wantD[0].Value, d[0].Value)
			} else {
				assert.Equal(t, tc.wantD, d)
			}
		})
	}
}

func Test_toBson(t *testing.T) {

	testCases := []struct {
		name string
		data any
		want bson.D
	}{
		{
			name: "nil data",
			data: nil,
			want: nil,
		},
		{
			name: "bson.D",
			data: bson.D{
				bson.E{Key: "k1", Value: "v1"},
			},
			want: bson.D{
				bson.E{Key: "k1", Value: "v1"},
			},
		},
		{
			name: "map pointer",
			data: func() *map[string]any {
				data := map[string]any{
					"k1": "v1",
				}
				return &data
			}(),
			want: bson.D{
				bson.E{Key: "k1", Value: "v1"},
			},
		},
		{
			name: "map",
			data: map[string]any{
				"k1": "v1",
			},
			want: bson.D{
				bson.E{Key: "k1", Value: "v1"},
			},
		},
		{
			name: "struct pointer",
			data: func() *types.TestUser {
				data := types.TestUser{Id: "123", Name: "cmy", Age: int64(24)}
				return &data
			}(),
			want: bson.D{
				bson.E{Key: "_id", Value: "123"},
				bson.E{Key: "name", Value: "cmy"},
				bson.E{Key: "age", Value: int64(24)},
			},
		},
		{
			name: "struct",
			data: types.TestUser{Id: "123", Name: "cmy", Age: int64(24)},
			want: bson.D{
				bson.E{Key: "_id", Value: "123"},
				bson.E{Key: "name", Value: "cmy"},
				bson.E{Key: "age", Value: int64(24)},
			},
		},
		{
			name: "struct pointer with empty-value",
			data: &types.TestUser{},
			want: bson.D{
				bson.E{Key: "_id", Value: ""},
				bson.E{Key: "name", Value: ""},
				bson.E{Key: "age", Value: int64(0)},
			},
		},
		{
			name: "struct pointer with no empty-value",
			data: &types.TestUser{Id: "123", Name: "cmy", Age: int64(24)},
			want: bson.D{
				bson.E{Key: "_id", Value: "123"},
				bson.E{Key: "name", Value: "cmy"},
				bson.E{Key: "age", Value: int64(24)},
			},
		},
		{
			name: "not struct",
			data: 1,
			want: nil,
		},
		{
			name: "empty struct",
			data: struct{}{},
			want: bson.D{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, ToBson(tc.data))
		})
	}
}

func Test_toSetBson(t *testing.T) {

	testCases := []struct {
		name    string
		updates any
		want    bson.D
	}{
		{
			name:    "nil data",
			updates: nil,
			want:    nil,
		},
		{
			name: "bson.D",
			updates: bson.D{
				bson.E{Key: types.Set, Value: bson.D{bson.E{Key: "k1", Value: "v1"}}},
			},
			want: bson.D{bson.E{Key: types.Set, Value: bson.D{bson.E{Key: "k1", Value: "v1"}}}},
		},
		{
			name: "map",
			updates: map[string]any{
				"k1": "v1",
			},
			want: bson.D{bson.E{Key: types.Set, Value: bson.D{bson.E{Key: "k1", Value: "v1"}}}},
		},
		{
			name: "map pointer",
			updates: func() *map[string]any {
				data := map[string]any{
					"k1": "v1",
				}
				return &data
			}(),

			want: bson.D{bson.E{Key: types.Set, Value: bson.D{bson.E{Key: "k1", Value: "v1"}}}},
		},
		{
			name:    "empty struct",
			updates: struct{}{},

			want: bson.D{},
		},
		{
			name:    "struct",
			updates: types.UpdatedUser{Name: "cmy", Age: int64(24)},
			want: bson.D{
				bson.E{Key: types.Set, Value: bson.D{
					bson.E{Key: "name", Value: "cmy"},
					bson.E{Key: "age", Value: int64(24)},
				},
				},
			},
		},
		{
			name:    "struct pointer",
			updates: &types.UpdatedUser{Name: "cmy", Age: int64(24)},
			want: bson.D{
				bson.E{Key: types.Set, Value: bson.D{
					bson.E{Key: "name", Value: "cmy"},
					bson.E{Key: "age", Value: int64(24)},
				},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			setBson := ToSetBson(tc.updates)
			if len(setBson) > 0 {
				assert.Equal(t, tc.want[0].Value, setBson[0].Value)
			} else {
				assert.Equal(t, tc.want, setBson)
			}
		})
	}
}

func TestMapToSetBson(t *testing.T) {
	testCases := []struct {
		name  string
		data  map[string]any
		wantD bson.D
	}{
		{
			name:  "nil data",
			data:  nil,
			wantD: nil,
		},
		{
			name: "map with zero-value",
			data: map[string]any{
				"name": "",
			},
			wantD: bson.D{
				bson.E{Key: types.Set, Value: bson.D{
					bson.E{Key: "name", Value: ""},
				}},
			},
		},
		{
			name: "map with no zero-value",
			data: map[string]any{
				"name": "cmy",
			},
			wantD: bson.D{
				bson.E{Key: types.Set, Value: bson.D{
					bson.E{Key: "name", Value: "cmy"},
				},
				}},
		},
		{
			name: "empty map",
			data: map[string]any{},
			wantD: bson.D{
				bson.E{Key: types.Set, Value: bson.D{}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.data != nil {
				assert.Equal(t, tc.wantD[0].Value, MapToSetBson(tc.data)[0].Value)
			}
			assert.Equal(t, tc.wantD, MapToSetBson(tc.data))
		})
	}
}

func Test_structToBson(t *testing.T) {
	testCases := []struct {
		name string
		val  any
		want bson.D
	}{
		{
			name: "nil",
			val:  nil,
			want: nil,
		},
		{
			name: "struct with zero-value",
			val:  types.TestUser{},
			want: bson.D{
				bson.E{Key: "_id", Value: ""},
				bson.E{Key: "name", Value: ""},
				bson.E{Key: "age", Value: int64(0)},
			},
		},
		{
			name: "struct without zero-value",
			val: types.TestUser{
				Id:   "123",
				Name: "cmy",
				Age:  24,
			},
			want: bson.D{
				bson.E{Key: "_id", Value: "123"},
				bson.E{Key: "name", Value: "cmy"},
				bson.E{Key: "age", Value: int64(24)},
			},
		},
		{
			name: "struct pointer",
			val: &types.TestUser{
				Id:   "123",
				Name: "cmy",
				Age:  24,
			},
			want: bson.D{
				bson.E{Key: "_id", Value: "123"},
				bson.E{Key: "name", Value: "cmy"},
				bson.E{Key: "age", Value: int64(24)},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, structToBson(tc.val))
		})
	}
}

func TestStructToBson(t *testing.T) {
	testCases := []struct {
		name  string
		data  any
		wantD bson.D
	}{
		{
			name:  "nil",
			data:  nil,
			wantD: nil,
		},
		{
			name:  "not struct",
			data:  "123",
			wantD: nil,
		},
		{
			name: "struct with zero-value",
			data: types.TestUser{},
			wantD: bson.D{
				bson.E{Key: "_id", Value: ""},
				bson.E{Key: "name", Value: ""},
				bson.E{Key: "age", Value: int64(0)},
			},
		},
		{
			name: "struct without zero-value",
			data: types.TestUser{
				Id:   "123",
				Name: "cmy",
				Age:  24,
			},
			wantD: bson.D{
				bson.E{Key: "_id", Value: "123"},
				bson.E{Key: "name", Value: "cmy"},
				bson.E{Key: "age", Value: int64(24)},
			},
		},
		{
			name: "struct pointer",
			data: &types.TestUser{
				Id:   "123",
				Name: "cmy",
				Age:  24,
			},
			wantD: bson.D{
				bson.E{Key: "_id", Value: "123"},
				bson.E{Key: "name", Value: "cmy"},
				bson.E{Key: "age", Value: int64(24)},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantD, StructToBson(tc.data))
		})
	}
}
