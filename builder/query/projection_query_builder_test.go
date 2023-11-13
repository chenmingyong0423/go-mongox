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

package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_projectionQueryBuilder_Slice(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "key", Value: bson.D{bson.E{Key: "$slice", Value: 1}}}}, BsonBuilder().Slice("key", 1).Build())
}

func Test_projectionQueryBuilder_SliceRanger(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "key", Value: bson.D{bson.E{Key: "$slice", Value: []int{1, 2}}}}}, BsonBuilder().SliceRanger("key", 1, 2).Build())
}
