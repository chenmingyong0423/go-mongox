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
	"github.com/chenmingyong0423/go-mongox/types"
	"go.mongodb.org/mongo-driver/bson"
)

type projectionQueryBuilder struct {
	parent *Builder
}

func (b *projectionQueryBuilder) Slice(key string, number int) *Builder {
	e := bson.E{Key: types.Slice, Value: number}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *projectionQueryBuilder) SliceRanger(key string, start, end int) *Builder {
	e := bson.E{Key: types.Slice, Value: []int{start, end}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}
