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

type arrayUpdateBuilder struct {
	parent *Builder
}

func (b *arrayUpdateBuilder) AddToSet(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(types.AddToSet, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: types.AddToSet, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) Pop(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(types.Pop, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: types.Pop, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) Pull(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(types.Pull, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: types.Pull, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) Push(key string, value any) *Builder {
	e := bson.E{Key: key, Value: value}
	if !b.parent.tryMergeValue(types.Push, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: types.Push, Value: bson.D{bson.E{Key: key, Value: value}}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) PullAll(key string, values ...any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.PullAll, Value: bson.D{bson.E{Key: key, Value: values}}})
	return b.parent
}

func (b *arrayUpdateBuilder) PullAllInt(key string, values ...int) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.PullAll, Value: bson.D{bson.E{Key: key, Value: values}}})
	return b.parent
}

func (b *arrayUpdateBuilder) PullAllInt8(key string, values ...int8) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.PullAll, Value: bson.D{bson.E{Key: key, Value: values}}})

	return b.parent
}

func (b *arrayUpdateBuilder) PullAllInt16(key string, values ...int16) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.PullAll, Value: bson.D{bson.E{Key: key, Value: values}}})

	return b.parent
}

func (b *arrayUpdateBuilder) PullAllInt32(key string, values ...int32) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.PullAll, Value: bson.D{bson.E{Key: key, Value: values}}})
	return b.parent
}

func (b *arrayUpdateBuilder) PullAllInt64(key string, values ...int64) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.PullAll, Value: bson.D{bson.E{Key: key, Value: values}}})
	return b.parent
}

func (b *arrayUpdateBuilder) PullAllUint(key string, values ...uint) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.PullAll, Value: bson.D{bson.E{Key: key, Value: values}}})
	return b.parent
}

func (b *arrayUpdateBuilder) PullAllUint8(key string, values ...uint8) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.PullAll, Value: bson.D{{Key: key, Value: values}}})
	return b.parent
}

func (b *arrayUpdateBuilder) PullAllUint16(key string, values ...uint16) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.PullAll, Value: bson.D{bson.E{Key: key, Value: values}}})
	return b.parent
}

func (b *arrayUpdateBuilder) PullAllUint32(key string, values ...uint32) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.PullAll, Value: bson.D{bson.E{Key: key, Value: values}}})
	return b.parent
}

func (b *arrayUpdateBuilder) PullAllUint64(key string, values ...uint64) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.PullAll, Value: bson.D{bson.E{Key: key, Value: values}}})
	return b.parent

}

func (b *arrayUpdateBuilder) PullAllFloat32(key string, values ...float32) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.PullAll, Value: bson.D{bson.E{Key: key, Value: values}}})
	return b.parent
}

func (b *arrayUpdateBuilder) PullAllFloat64(key string, values ...float64) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.PullAll, Value: bson.D{bson.E{Key: key, Value: values}}})
	return b.parent
}

func (b *arrayUpdateBuilder) PullAllString(key string, values ...string) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.PullAll, Value: bson.D{bson.E{Key: key, Value: values}}})
	return b.parent
}

func (b *arrayUpdateBuilder) Each(values ...any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Each, Value: values})
	return b.parent
}

func (b *arrayUpdateBuilder) EachInt(values ...int) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Each, Value: values})
	return b.parent
}

func (b *arrayUpdateBuilder) EachInt8(values ...int8) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Each, Value: values})
	return b.parent
}

func (b *arrayUpdateBuilder) EachInt16(values ...int16) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Each, Value: values})

	return b.parent
}

func (b *arrayUpdateBuilder) EachInt32(values ...int32) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Each, Value: values})

	return b.parent
}

func (b *arrayUpdateBuilder) EachInt64(values ...int64) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Each, Value: values})

	return b.parent
}

func (b *arrayUpdateBuilder) EachUint(values ...uint) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Each, Value: values})

	return b.parent
}

func (b *arrayUpdateBuilder) EachUint8(values ...uint8) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Each, Value: values})

	return b.parent
}

func (b *arrayUpdateBuilder) EachUint16(values ...uint16) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Each, Value: values})

	return b.parent
}

func (b *arrayUpdateBuilder) EachUint32(values ...uint32) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Each, Value: values})

	return b.parent
}

func (b *arrayUpdateBuilder) EachUint64(values ...uint64) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Each, Value: values})

	return b.parent
}

func (b *arrayUpdateBuilder) EachFloat32(values ...float32) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Each, Value: values})

	return b.parent
}

func (b *arrayUpdateBuilder) EachFloat64(values ...float64) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Each, Value: values})

	return b.parent
}

func (b *arrayUpdateBuilder) EachString(values ...string) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Each, Value: values})

	return b.parent
}

func (b *arrayUpdateBuilder) Position(value int) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Position, Value: value})
	return b.parent
}

func (b *arrayUpdateBuilder) Slice(num int) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.SliceForUpdate, Value: num})
	return b.parent
}

func (b *arrayUpdateBuilder) Sort(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Sort, Value: value})
	return b.parent
}
