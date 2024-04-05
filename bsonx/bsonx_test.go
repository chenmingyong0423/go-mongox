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

package bsonx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestD(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "name", Value: "Mingyong Chen"}}, D("name", "Mingyong Chen"))
}

func TestId(t *testing.T) {
	assert.Equal(t, bson.M{"_id": "1"}, Id("1"))
}

func TestE(t *testing.T) {
	assert.Equal(t, bson.E{Key: "name", Value: "chenmingyong"}, E("name", "chenmingyong"))
}

func TestA(t *testing.T) {
	testCases := []struct {
		name   string
		values []any
		want   bson.A
	}{
		{
			name:   "nil values",
			values: nil,
			want:   bson.A{},
		},
		{
			name:   "empty values",
			values: []any{},
			want:   bson.A{},
		},
		{
			name:   "multiple values",
			values: []any{"1", "2"},
			want:   bson.A{"1", "2"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, A(tc.values...))
		})
	}

}
