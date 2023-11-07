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

func (b *fieldUpdateBuilder) Set(key string, value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Set, Value: bson.D{bson.E{Key: key, Value: value}}})
	return b.parent
}

func (b *fieldUpdateBuilder) SetForMap(data map[string]any) *Builder {
	if d := converter.MapToBson(data); len(d) != 0 {
		b.parent.data = append(b.parent.data, bson.E{Key: types.Set, Value: d})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) SetForStruct(data any) *Builder {
	if d := converter.StructToBson(data); len(d) != 0 {
		b.parent.data = append(b.parent.data, bson.E{Key: types.Set, Value: d})
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

func (b *fieldUpdateBuilder) SetOnInsert(bsonElements ...types.KeyValue) *Builder {
	value := bson.D{}
	for _, element := range bsonElements {
		value = append(value, bson.E{Key: element.Key, Value: element.Value})
	}

	b.parent.data = append(b.parent.data, bson.E{Key: types.SetOnInsert, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) SetOnInsertForMap(data map[string]any) *Builder {
	if data != nil {
		b.parent.data = append(b.parent.data, bson.E{Key: types.SetOnInsert, Value: converter.MapToBson(data)})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) CurrentDate(bsonElements ...types.KeyValue) *Builder {
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
	if data != nil {
		d := bson.D{}
		for k, v := range data {
			if val, ok := v.(bool); ok {
				d = append(d, bson.E{Key: k, Value: val})
			} else {
				d = append(d, bson.E{Key: k, Value: bson.M{types.Type: v}})
			}
		}
		b.parent.data = append(b.parent.data, bson.E{Key: types.CurrentDate, Value: d})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Inc(bsonElements ...types.KeyValue) *Builder {
	value := bson.D{}
	for _, element := range bsonElements {
		if val, ok := element.Value.(int); ok {
			value = append(value, bson.E{Key: element.Key, Value: val})
		}
	}
	b.parent.data = append(b.parent.data, bson.E{Key: types.Inc, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) IncForMap(data map[string]int) *Builder {
	if data != nil {
		d := bson.D{}
		for k, v := range data {
			d = append(d, bson.E{Key: k, Value: v})
		}
		b.parent.data = append(b.parent.data, bson.E{Key: types.Inc, Value: d})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Min(bsonElements ...types.KeyValue) *Builder {
	value := bson.D{}
	for _, element := range bsonElements {
		value = append(value, bson.E{Key: element.Key, Value: element.Value})
	}
	b.parent.data = append(b.parent.data, bson.E{Key: types.Min, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) MinForMap(data map[string]any) *Builder {
	if data != nil {
		b.parent.data = append(b.parent.data, bson.E{Key: types.Min, Value: converter.MapToBson(data)})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Max(bsonElements ...types.KeyValue) *Builder {
	value := bson.D{}
	for _, element := range bsonElements {
		value = append(value, bson.E{Key: element.Key, Value: element.Value})
	}
	b.parent.data = append(b.parent.data, bson.E{Key: types.Max, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) MaxForMap(data map[string]any) *Builder {
	if data != nil {
		b.parent.data = append(b.parent.data, bson.E{Key: types.Max, Value: converter.MapToBson(data)})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Mul(bsonElements ...types.KeyValue) *Builder {
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

func (b *fieldUpdateBuilder) MulForMap(data map[string]any) *Builder {
	if data != nil {
		d := bson.D{}
		for k, v := range data {
			if utils.IsNumeric(v) {
				d = append(d, bson.E{Key: k, Value: v})
			}
		}
		b.parent.data = append(b.parent.data, bson.E{Key: types.Mul, Value: d})
	}
	return b.parent
}

func (b *fieldUpdateBuilder) Rename(keyValues ...string) *Builder {
	value := bson.D{}
	if len(keyValues)%2 == 0 {
		for i := 0; i < len(keyValues); i += 2 {
			value = append(value, bson.E{Key: keyValues[i], Value: keyValues[i+1]})
		}
	}
	b.parent.data = append(b.parent.data, bson.E{Key: types.Rename, Value: value})
	return b.parent
}

func (b *fieldUpdateBuilder) RenameForMap(data map[string]string) *Builder {
	if data != nil {
		d := bson.D{}
		for k, v := range data {
			d = append(d, bson.E{Key: k, Value: v})
		}
		b.parent.data = append(b.parent.data, bson.E{Key: types.Rename, Value: d})
	}
	return b.parent
}
