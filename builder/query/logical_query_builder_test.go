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
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Test_logicalQueryBuilder_And(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "$and", Value: []any{bson.D{bson.E{Key: "name", Value: "cmy"}}}}}, NewBuilder().And(bson.D{{Key: "name", Value: "cmy"}}).Build())
}

func Test_logicalQueryBuilder_Not(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "$not", Value: bson.D{{Key: "name", Value: "cmy"}}}}, NewBuilder().Not(bson.D{{Key: "name", Value: "cmy"}}).Build())
}

func Test_logicalQueryBuilder_Nor(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "$nor", Value: []any{bson.D{bson.E{Key: "name", Value: "cmy"}}}}}, NewBuilder().Nor(bson.D{{Key: "name", Value: "cmy"}}).Build())
}

func Test_logicalQueryBuilder_Or(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "$or", Value: []any{bson.D{{Key: "name", Value: "cmy"}}}}}, NewBuilder().Or(bson.D{{Key: "name", Value: "cmy"}}).Build())
}
