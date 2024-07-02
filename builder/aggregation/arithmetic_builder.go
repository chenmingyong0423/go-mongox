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

package aggregation

import (
	"go.mongodb.org/mongo-driver/bson"
)

type arithmeticBuilder struct {
	parent *Builder
}

func (b *arithmeticBuilder) Add(key string, expressions ...any) *Builder {
	e := bson.E{Key: AddOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) AddWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: AddOp, Value: expressions})
	return b.parent
}

func (b *arithmeticBuilder) Multiply(key string, expressions ...any) *Builder {
	e := bson.E{Key: MultiplyOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) MultiplyWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: MultiplyOp, Value: expressions})
	return b.parent
}

func (b *arithmeticBuilder) Subtract(key string, expressions ...any) *Builder {
	e := bson.E{Key: SubtractOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) SubtractWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: SubtractOp, Value: expressions})
	return b.parent
}

func (b *arithmeticBuilder) Divide(key string, expressions ...any) *Builder {
	e := bson.E{Key: DivideOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) DivideWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: DivideOp, Value: expressions})
	return b.parent
}

func (b *arithmeticBuilder) Mod(key string, expressions ...any) *Builder {
	e := bson.E{Key: ModOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) ModWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: ModOp, Value: expressions})
	return b.parent
}

func (b *arithmeticBuilder) Abs(key string, expression any) *Builder {
	e := bson.E{Key: Abs, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) AbsWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: Abs, Value: expression})
	return b.parent
}
