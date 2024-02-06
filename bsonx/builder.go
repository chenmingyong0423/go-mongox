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

import "go.mongodb.org/mongo-driver/bson"

// DBuilder is a builder for bson.D
type DBuilder struct {
	d bson.D
}

func NewD() *DBuilder {
	return &DBuilder{d: bson.D{}}
}

func (b *DBuilder) Add(key string, value any) *DBuilder {
	b.d = append(b.d, bson.E{Key: key, Value: value})
	return b
}

func (b *DBuilder) Build() bson.D {
	return b.d
}
