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

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_arithmeticBuilder_Add(t *testing.T) {
	t.Run("test add", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$add", Value: []any{1, 2, 3, "$a", "$b", "$c"}}}}},
			BsonBuilder().Add("total", 1, 2, 3, "$a", "$b", "$c").Build(),
		)
	})
}

func Test_arithmeticBuilder_Multiply(t *testing.T) {
	t.Run("test multiply", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$multiply", Value: []any{1, 2, 3, "$a", "$b", "$c"}}}}},
			BsonBuilder().Multiply("total", 1, 2, 3, "$a", "$b", "$c").Build(),
		)
	})
}

func Test_arithmeticBuilder_Subtract(t *testing.T) {
	t.Run("test subtract", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$subtract", Value: []any{"$quarter", int64(0), int64(2)}}}}},
			BsonBuilder().Subtract("total", "$quarter", int64(0), int64(2)).Build(),
		)
	})

}

func Test_arithmeticBuilder_Divide(t *testing.T) {
	t.Run("test divide", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$divide", Value: []any{1, 2, 3, "$a", "$b", "$c"}}}}},
			BsonBuilder().Divide("total", 1, 2, 3, "$a", "$b", "$c").Build(),
		)
	})
}

func Test_arithmeticBuilder_Mod(t *testing.T) {
	t.Run("test mod", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "total", Value: bson.D{bson.E{Key: "$mod", Value: []any{1, 2, 3, "$a", "$b", "$c"}}}}},
			BsonBuilder().Mod("total", 1, 2, 3, "$a", "$b", "$c").Build(),
		)
	})
}
