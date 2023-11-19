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

func AddToSet(value any) bson.D {
	return bson.D{{Key: types.AddToSet, Value: value}}
}

func Pop(value any) bson.D {
	return bson.D{{Key: types.Pop, Value: value}}
}

func Pull(value any) bson.D {
	return bson.D{{Key: types.Pull, Value: value}}
}

func Push(value any) bson.D {
	return bson.D{{Key: types.Push, Value: value}}
}

func PullAll[T any](key string, values ...T) bson.D {
	return bson.D{{Key: types.PullAll, Value: bson.D{bson.E{Key: key, Value: values}}}}
}

func Each[T any](key string, values ...T) bson.D {
	return bson.D{{Key: key, Value: bson.D{{Key: types.Each, Value: values}}}}
}

func Position(key string, value any) bson.D {
	return bson.D{{Key: key, Value: bson.D{{Key: types.Position, Value: value}}}}
}

func Slice(key string, num int) bson.D {
	return bson.D{{Key: key, Value: bson.D{{Key: types.SliceForUpdate, Value: num}}}}
}

func Sort(key string, value any) bson.D {
	return bson.D{{Key: key, Value: bson.D{{Key: types.Sort, Value: value}}}}
}

func Set(value any) bson.D {
	return bson.D{{Key: types.Set, Value: value}}
}

func Unset(keys ...string) bson.D {
	value := bson.D{}
	for i := range keys {
		value = append(value, bson.E{Key: keys[i], Value: ""})
	}
	return bson.D{{Key: types.Unset, Value: value}}
}

func SetOnInsert(value any) bson.D {
	return bson.D{{Key: types.SetOnInsert, Value: value}}
}

func CurrentDate(value any) bson.D {
	return bson.D{{Key: types.CurrentDate, Value: value}}
}

func Inc(value any) bson.D {
	return bson.D{{Key: types.Inc, Value: value}}
}

func Min(value any) bson.D {
	return bson.D{{Key: types.Min, Value: value}}
}

func Max(value any) bson.D {
	return bson.D{{Key: types.Max, Value: value}}
}

func Mul(value any) bson.D {
	return bson.D{{Key: types.Mul, Value: value}}
}

func Rename(value any) bson.D {
	return bson.D{{Key: types.Rename, Value: value}}
}
