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

package aggregation

import (
	"testing"

	"github.com/chenmingyong0423/go-mongox/converter"
	"github.com/chenmingyong0423/go-mongox/types"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestBuilder_AddKeyValues(t *testing.T) {
	testCases := []struct {
		name      string
		keyValues []types.KeyValue[any]
		expected  bson.D
	}{
		{
			name:      "nil keyValues",
			keyValues: nil,
			expected:  bson.D{},
		},
		{
			name:      "normal",
			keyValues: []types.KeyValue[any]{converter.KeyValue[any]("name", "cmy")},
			expected:  bson.D{{Key: "name", Value: "cmy"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().AddKeyValues(tc.keyValues...).Build())
		})
	}
}
