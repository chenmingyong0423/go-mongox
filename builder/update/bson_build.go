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
	"go.mongodb.org/mongo-driver/v2/bson"
)

func AddToSet(key string, value any) bson.D {
	return bson.D{{Key: AddToSetOp, Value: bson.D{{Key: key, Value: value}}}}
}

func Pop(key string, value any) bson.D {
	return bson.D{{Key: PopOp, Value: bson.D{{Key: key, Value: value}}}}
}

func Pull(key string, value any) bson.D {
	return bson.D{{Key: PullOp, Value: bson.D{{Key: key, Value: value}}}}
}

func Push(key string, value any) bson.D {
	return bson.D{{Key: PushOp, Value: bson.D{{Key: key, Value: value}}}}
}

func PullAll[T any](key string, values ...T) bson.D {
	return bson.D{{Key: PullAllOp, Value: bson.D{bson.E{Key: key, Value: values}}}}
}

func Each[T any](values ...T) bson.D {
	return bson.D{{Key: EachOp, Value: values}}
}

func Position(value any) bson.D {
	return bson.D{{Key: PositionOp, Value: value}}
}

func Slice(num int) bson.D {
	return bson.D{{Key: SliceForUpdateOp, Value: num}}
}

func Sort(value any) bson.D {
	return bson.D{{Key: SortOp, Value: value}}
}

func Set(key string, value any) bson.D {
	return bson.D{{Key: SetOp, Value: bson.D{{Key: key, Value: value}}}}
}

func SetFields(value any) bson.D {
	return bson.D{{Key: SetOp, Value: value}}
}

func Unset(keys ...string) bson.D {
	value := bson.D{}
	for i := range keys {
		value = append(value, bson.E{Key: keys[i], Value: ""})
	}
	return bson.D{{Key: UnsetOp, Value: value}}
}

func SetOnInsert(key string, value any) bson.D {
	return bson.D{{Key: SetOnInsertOp, Value: bson.D{{Key: key, Value: value}}}}
}

func SetFieldsOnInsert(value any) bson.D {
	return bson.D{{Key: SetOnInsertOp, Value: value}}
}

func CurrentDate(key string, value any) bson.D {
	return bson.D{{Key: CurrentDateOp, Value: bson.D{{Key: key, Value: value}}}}
}

func Inc(key string, value any) bson.D {
	return bson.D{{Key: IncOp, Value: bson.D{{Key: key, Value: value}}}}
}

func Min(key string, value any) bson.D {
	return bson.D{{Key: MinOp, Value: bson.D{{Key: key, Value: value}}}}
}

func Max(key string, value any) bson.D {
	return bson.D{{Key: MaxOp, Value: bson.D{{Key: key, Value: value}}}}
}

func Mul(key string, value any) bson.D {
	return bson.D{{Key: MulOp, Value: bson.D{{Key: key, Value: value}}}}
}

func Rename(key string, value any) bson.D {
	return bson.D{{Key: RenameOp, Value: bson.D{{Key: key, Value: value}}}}
}
