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

type fieldUpdateBuilder struct {
	parent *Builder
}

func (b *fieldUpdateBuilder) Set(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(types.Set, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: types.Set, Value: bson.D{e}})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Unset(keys ...string) *Builder {
	value := bson.D{}
	for i := range keys {
		value = append(value, bson.E{Key: keys[i], Value: ""})
	}
	b.parent.data = append(b.parent.data, bson.E{Key: types.Unset, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) SetOnInsert(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(types.SetOnInsert, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: types.SetOnInsert, Value: bson.D{e}})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) CurrentDate(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(types.CurrentDate, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: types.CurrentDate, Value: bson.D{e}})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Inc(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(types.Inc, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: types.Inc, Value: bson.D{e}})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Min(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(types.Min, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: types.Min, Value: bson.D{e}})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Max(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(types.Max, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: types.Max, Value: bson.D{e}})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Mul(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(types.Mul, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: types.Mul, Value: bson.D{e}})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Rename(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(types.Rename, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: types.Rename, Value: bson.D{e}})
	}
	return b.parent
}
