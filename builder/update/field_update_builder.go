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
	"github.com/chenmingyong0423/go-mongox/converter"
	"github.com/chenmingyong0423/go-mongox/pkg/utils"
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

func (b *fieldUpdateBuilder) SetKeyValues(bsonElements ...types.KeyValue) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Set, Value: converter.KeyValuesToBson(bsonElements...)})
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

func (b *fieldUpdateBuilder) SetOnInsertKeyValues(bsonElements ...types.KeyValue) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.SetOnInsert, Value: converter.KeyValuesToBson(bsonElements...)})
	return b.parent
}

func (b *fieldUpdateBuilder) CurrentDate(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.CurrentDate, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) CurrentDateKeyValues(bsonElements ...types.KeyValue) *Builder {
	value := bson.D{}
	for _, element := range bsonElements {
		if v, ok := element.Value.(bool); ok {
			value = append(value, bson.E{Key: element.Key, Value: v})
		} else {
			value = append(value, bson.E{Key: element.Key, Value: bson.M{types.Type: element.Value}})
		}
	}
	b.parent.data = append(b.parent.data, bson.E{Key: types.CurrentDate, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) CurrentDateForMap(data map[string]any) *Builder {
	d := bson.D{}
	for k, v := range data {
		if val, ok := v.(bool); ok {
			d = append(d, bson.E{Key: k, Value: val})
		} else {
			d = append(d, bson.E{Key: k, Value: bson.M{types.Type: v}})
		}
	}
	b.parent.data = append(b.parent.data, bson.E{Key: types.CurrentDate, Value: d})
	return b.parent
}

func (b *fieldUpdateBuilder) Inc(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Inc, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) IncKeyValues(bsonElements ...types.KeyValue) *Builder {
	value := bson.D{}
	for _, element := range bsonElements {
		if val, ok := element.Value.(int); ok {
			value = append(value, bson.E{Key: element.Key, Value: val})
		}
	}
	b.parent.data = append(b.parent.data, bson.E{Key: types.Inc, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) Min(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Min, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) MinKeyValues(bsonElements ...types.KeyValue) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Min, Value: converter.KeyValuesToBson(bsonElements...)})
	return b.parent
}

func (b *fieldUpdateBuilder) Max(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Max, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) MaxKeyValues(bsonElements ...types.KeyValue) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Max, Value: converter.KeyValuesToBson(bsonElements...)})
	return b.parent
}

func (b *fieldUpdateBuilder) Mul(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Mul, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) MulKeyValues(bsonElements ...types.KeyValue) *Builder {
	value := bson.D{}
	for _, element := range bsonElements {
		v := element.Value
		if utils.IsNumeric(v) {
			value = append(value, bson.E{Key: element.Key, Value: v})
		}
	}
	b.parent.data = append(b.parent.data, bson.E{Key: types.Mul, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) Rename(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Rename, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) RenameKeyValues(bsonElements ...types.KeyValue) *Builder {
	value := bson.D{}
	for _, element := range bsonElements {
		if v, ok := element.Value.(string); ok {
			value = append(value, bson.E{Key: element.Key, Value: v})
		}
	}
	b.parent.data = append(b.parent.data, bson.E{Key: types.Rename, Value: value})
	return b.parent
}
