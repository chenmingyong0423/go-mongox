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

	"github.com/chenmingyong0423/go-mongox/bsonx"
	"github.com/chenmingyong0423/go-mongox/types"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_condBuilder_Cond(t *testing.T) {
	t.Run("test Cond", func(t *testing.T) {
		assert.Equal(t,
			bson.D{bson.E{Key: "discount", Value: bson.D{{Key: "$cond", Value: []any{bson.D{{Key: "$gte", Value: []any{"$qty", 250}}}, 30, 20}}}}},
			BsonBuilder().Cond("discount", bson.D{{Key: "$gte", Value: []any{"$qty", 250}}}, 30, 20).Build(),
		)
	})
}

func Test_condBuilder_IfNull(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "discount", Value: bson.D{{Key: "$ifNull", Value: []any{"$coupon", int64(0)}}}}}, BsonBuilder().IfNull("discount", "$coupon", int64(0)).Build())
}

func Test_condBuilder_Switch(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "summary", Value: bson.D{
		{Key: "$switch", Value: bson.D{
			{Key: "branches", Value: bson.A{
				bson.D{{Key: "case", Value: bson.D{{Key: "$eq", Value: []any{0, 5}}}}, {Key: "then", Value: "equals"}},
				bson.D{{Key: "case", Value: bson.D{{Key: "$gt", Value: []any{0, 5}}}}, {Key: "then", Value: "greater than"}},
			}},
			{Key: "default", Value: "Did not match"},
		}},
	}}},
		BsonBuilder().Switch("summary", []types.CaseThen{
			{
				Case: bsonx.D(bsonx.E("$eq", []any{0, 5})), Then: "equals",
			},
			{
				Case: bsonx.D(bsonx.E("$gt", []any{0, 5})), Then: "greater than",
			},
		}, "Did not match").Build(),
	)
}
