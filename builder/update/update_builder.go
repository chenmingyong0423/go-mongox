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

package update

import (
	"github.com/chenmingyong0423/go-mongox/types"
	"go.mongodb.org/mongo-driver/bson"
)

func BsonBuilder() *Builder {
	b := &Builder{data: bson.D{}}
	b.fieldUpdateBuilder = fieldUpdateBuilder{parent: b}
	b.arrayUpdateBuilder = arrayUpdateBuilder{parent: b}
	return b
}

type Builder struct {
	data bson.D
	fieldUpdateBuilder
	arrayUpdateBuilder
}

func (b *Builder) Add(bsonElements ...types.KeyValue[any]) *Builder {
	if len(bsonElements) != 0 {
		for _, element := range bsonElements {
			b.data = append(b.data, bson.E{Key: element.Key, Value: element.Value})
		}
	}
	return b
}

func (b *Builder) Build() bson.D {
	return b.data
}
