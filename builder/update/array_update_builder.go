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

func (b *arrayUpdateBuilder) AddToSet(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.AddToSet, Value: value})
	return b.parent
}

func (b *arrayUpdateBuilder) Pop(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Pop, Value: value})
	return b.parent
}

func (b *arrayUpdateBuilder) Pull(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Pull, Value: value})
	return b.parent
}

func (b *arrayUpdateBuilder) Push(value any) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: types.Push, Value: value})
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

func (b *arrayUpdateBuilder) Each(key string, values ...any) *Builder {
	e := bson.E{Key: types.Each, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) EachInt(key string, values ...int) *Builder {
	e := bson.E{Key: types.Each, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) EachInt8(key string, values ...int8) *Builder {
	e := bson.E{Key: types.Each, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) EachInt16(key string, values ...int16) *Builder {
	e := bson.E{Key: types.Each, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) EachInt32(key string, values ...int32) *Builder {
	e := bson.E{Key: types.Each, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) EachInt64(key string, values ...int64) *Builder {
	e := bson.E{Key: types.Each, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) EachUint(key string, values ...uint) *Builder {
	e := bson.E{Key: types.Each, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) EachUint8(key string, values ...uint8) *Builder {
	e := bson.E{Key: types.Each, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) EachUint16(key string, values ...uint16) *Builder {
	e := bson.E{Key: types.Each, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) EachUint32(key string, values ...uint32) *Builder {
	e := bson.E{Key: types.Each, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) EachUint64(key string, values ...uint64) *Builder {
	e := bson.E{Key: types.Each, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) EachFloat32(key string, values ...float32) *Builder {
	e := bson.E{Key: types.Each, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) EachFloat64(key string, values ...float64) *Builder {
	e := bson.E{Key: types.Each, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) EachString(key string, values ...string) *Builder {
	e := bson.E{Key: types.Each, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) Position(key string, value int) *Builder {
	e := bson.E{Key: types.Position, Value: value}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) Slice(key string, num int) *Builder {
	e := bson.E{Key: types.SliceForUpdate, Value: num}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arrayUpdateBuilder) Sort(key string, value any) *Builder {
	e := bson.E{Key: types.Sort, Value: value}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}
