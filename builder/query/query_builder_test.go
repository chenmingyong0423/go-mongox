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

package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestQuery(t *testing.T) {
	query := Query()
	assert.NotNil(t, query)
	assert.Equal(t, bson.D{}, query.Build())
}

func TestQueryBuilder_Id(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "_id", Value: "123"}}, Query().Id("123").Build())
}

func TestQueryBuilder_Add(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "name", Value: "cmy"}}, Query().Add("name", "cmy").Build())
}
