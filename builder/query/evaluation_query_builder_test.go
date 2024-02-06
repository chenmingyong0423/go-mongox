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

func Test_evaluationQueryBuilder_Expr(t *testing.T) {
	assert.Equal(t,
		bson.D{{Key: "$expr", Value: bson.D{{Key: "$gt", Value: []any{"$spent", "$budget"}}}}},
		BsonBuilder().Expr(BsonBuilder().Add("$gt", []any{
			"$spent",
			"$budget",
		}).Build()).Build())
}

func Test_evaluationQueryBuilder_JsonSchema(t *testing.T) {
	// $jsonSchema: {
	//     required: [ "name", "major", "gpa", "address" ],
	//     properties: {
	//        name: {
	//           bsonType: "string",
	//           description: "must be a string and is required"
	//        },
	//        address: {
	//           bsonType: "object",
	//           required: [ "zipcode" ],
	//           properties: {
	//               "street": { bsonType: "string" },
	//               "zipcode": { bsonType: "string" }
	//           }
	//        }
	//     }
	//  }

	assert.Equal(t,
		bson.D{bson.E{Key: "$jsonSchema", Value: bson.D{
			bson.E{Key: "required", Value: []string{"name", "major", "gpa", "address"}},
			bson.E{Key: "properties", Value: bson.D{
				bson.E{Key: "name", Value: bson.D{
					bson.E{Key: "bsonType", Value: "string"},
					bson.E{Key: "description", Value: "must be a string and is required"},
				}},
				bson.E{Key: "address", Value: bson.D{
					bson.E{Key: "bsonType", Value: "object"},
					bson.E{Key: "required", Value: []string{"zipcode"}},
					bson.E{Key: "properties", Value: bson.D{
						bson.E{Key: "street", Value: bson.D{bson.E{Key: "bsonType", Value: "string"}}},
						bson.E{Key: "zipcode", Value: bson.D{bson.E{Key: "bsonType", Value: "string"}}},
					}},
				}},
			},
			},
		}}},
		BsonBuilder().JsonSchema(
			BsonBuilder().
				Add("required", []string{"name", "major", "gpa", "address"}).
				Add("properties",
					BsonBuilder().
						Add("name",
							BsonBuilder().
								Add("bsonType", "string").
								Add("description", "must be a string and is required").
								Build()).
						Add("address",
							BsonBuilder().Add("bsonType", "object").
								Add("required", []string{"zipcode"}).
								Add("properties",
									BsonBuilder().
										Add("street",
											BsonBuilder().
												Add("bsonType", "string").
												Build()).
										Add("zipcode",
											BsonBuilder().
												Add("bsonType", "string").
												Build()).
										Build()).
								Build()).
						Build()).
				Build()).
			Build())
}

func Test_evaluationQueryBuilder_Mod(t *testing.T) {

	testCases := []struct {
		name      string
		key       string
		divisor   any
		remainder int

		want bson.D
	}{
		{
			name:      "divisor not numeric",
			key:       "qty",
			divisor:   "4",
			remainder: 0,
			want:      bson.D{},
		},
		{
			name:      "divisor int",
			key:       "qty",
			divisor:   4,
			remainder: 0,
			want:      bson.D{{Key: "qty", Value: bson.D{{Key: "$mod", Value: bson.A{4, 0}}}}},
		},
		{
			name:      "divisor float",
			key:       "qty",
			divisor:   4.0,
			remainder: 0,
			want:      bson.D{{Key: "qty", Value: bson.D{{Key: "$mod", Value: bson.A{4.0, 0}}}}},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, BsonBuilder().Mod(tt.key, tt.divisor, tt.remainder).Build())
		})
	}
}

func Test_evaluationQueryBuilder_Regex(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "name", Value: bson.D{{Key: "$regex", Value: "acme.*corp"}}}},
		BsonBuilder().Regex("name", "acme.*corp").Build())
}

func Test_evaluationQueryBuilder_RegexOptions(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "name", Value: bson.D{{Key: "$regex", Value: "acme.*corp"}, {Key: "$options", Value: "i"}}}},
		BsonBuilder().RegexOptions("name", "acme.*corp", "i").Build())
}

func Test_evaluationQueryBuilder_Text(t *testing.T) {

	testCases := []struct {
		name               string
		value              string
		language           string
		caseSensitive      bool
		diacriticSensitive bool

		want bson.D
	}{
		{
			name:  "language, caseSensitive and diacriticSensitive are zero value",
			value: "java coffee shop",
			want:  bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: "java coffee shop"}}}},
		},
		{
			name:     "language is not zero value",
			value:    "java coffee shop",
			language: "en",
			want:     bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: "java coffee shop"}, {Key: "$language", Value: "en"}}}},
		},
		{
			name:          "caseSensitive is not zero value",
			value:         "java coffee shop",
			caseSensitive: true,
			want:          bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: "java coffee shop"}, {Key: "$caseSensitive", Value: true}}}},
		},
		{
			name:               "diacriticSensitive is not zero value",
			value:              "java coffee shop",
			diacriticSensitive: true,
			want:               bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: "java coffee shop"}, {Key: "$diacriticSensitive", Value: true}}}},
		},
		{
			name:               "language, caseSensitive and diacriticSensitive are not zero value",
			value:              "java coffee shop",
			language:           "en",
			caseSensitive:      true,
			diacriticSensitive: true,
			want:               bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: "java coffee shop"}, {Key: "$language", Value: "en"}, {Key: "$caseSensitive", Value: true}, {Key: "$diacriticSensitive", Value: true}}}},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, BsonBuilder().Text(tt.value, tt.language, tt.caseSensitive, tt.diacriticSensitive).Build())
		})
	}
}

func Test_evaluationQueryBuilder_Where(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "$where", Value: "this.credits == this.debits"}}, BsonBuilder().Where("this.credits == this.debits").Build())
}
