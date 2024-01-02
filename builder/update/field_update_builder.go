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

func (b *fieldUpdateBuilder) Set(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Set, Value: value})
	return b.parent
}

// SetSimple sets the value of a simple key in a document.
// pay attention to the following example:
// - don't use it with Set, if you want to set the multiple key at once, you can use Set directly.
func (b *fieldUpdateBuilder) SetSimple(key string, value any) *Builder {
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

func (b *fieldUpdateBuilder) SetOnInsert(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.SetOnInsert, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) CurrentDate(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.CurrentDate, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) Inc(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Inc, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) Min(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Min, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) Max(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Max, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) Mul(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Mul, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) Rename(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Rename, Value: value})
	return b.parent
}
