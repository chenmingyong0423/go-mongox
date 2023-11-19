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

type arrayQueryBuilder struct {
	parent *Builder
}

// All appends an element with '$all' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) All(key string, values ...any) *Builder {
	e := bson.E{Key: types.All, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

// AllUint appends an element with '$all' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) AllUint(key string, values ...uint) *Builder {
	e := bson.E{Key: types.All, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

// AllUint8 appends an element with '$all' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) AllUint8(key string, values ...uint8) *Builder {
	e := bson.E{Key: types.All, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

// AllUint16 appends an element with '$all' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) AllUint16(key string, values ...uint16) *Builder {
	e := bson.E{Key: types.All, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

// AllUint32 appends an element with '$all' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) AllUint32(key string, values ...uint32) *Builder {
	e := bson.E{Key: types.All, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

// AllUint64 appends an element with '$all' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) AllUint64(key string, values ...uint64) *Builder {
	e := bson.E{Key: types.All, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

// AllInt appends an element with '$all' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) AllInt(key string, values ...int) *Builder {
	e := bson.E{Key: types.All, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

// AllInt8 appends an element with '$all' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) AllInt8(key string, values ...int8) *Builder {
	e := bson.E{Key: types.All, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

// AllInt16 appends an element with '$all' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) AllInt16(key string, values ...int16) *Builder {
	e := bson.E{Key: types.All, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

// AllInt32 appends an element with '$all' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) AllInt32(key string, values ...int32) *Builder {
	e := bson.E{Key: types.All, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

// AllInt64 appends an element with '$all' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) AllInt64(key string, values ...int64) *Builder {
	e := bson.E{Key: types.All, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

// AllString appends an element with '$all' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) AllString(key string, values ...string) *Builder {
	e := bson.E{Key: types.All, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

// AllFloat32 appends an element with '$all' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) AllFloat32(key string, values ...float32) *Builder {
	e := bson.E{Key: types.All, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

// AllFloat64 appends an element with '$all' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) AllFloat64(key string, values ...float64) *Builder {
	e := bson.E{Key: types.All, Value: values}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

// ElemMatch appends an element with '$elemMatch' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) ElemMatch(key string, condition any) *Builder {
	e := bson.E{Key: types.ElemMatch, Value: condition}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

// Size appends an element with '$size' key and given value to the builder's data slice.
func (b *arrayQueryBuilder) Size(key string, size int) *Builder {
	e := bson.E{Key: types.Size, Value: size}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}
