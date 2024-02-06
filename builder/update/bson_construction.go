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

func AddToSet(key string, value any) bson.D {
	return bson.D{{Key: types.AddToSet, Value: bson.D{{Key: key, Value: value}}}}
}

func Pop(key string, value any) bson.D {
	return bson.D{{Key: types.Pop, Value: bson.D{{Key: key, Value: value}}}}
}

func Pull(key string, value any) bson.D {
	return bson.D{{Key: types.Pull, Value: bson.D{{Key: key, Value: value}}}}
}

func Push(key string, value any) bson.D {
	return bson.D{{Key: types.Push, Value: bson.D{{Key: key, Value: value}}}}
}

func PullAll[T any](key string, values ...T) bson.D {
	return bson.D{{Key: types.PullAll, Value: bson.D{bson.E{Key: key, Value: values}}}}
}

func Each[T any](values ...T) bson.D {
	return bson.D{{Key: types.Each, Value: values}}
}

func Position(value any) bson.D {
	return bson.D{{Key: types.Position, Value: value}}
}

func Slice(num int) bson.D {
	return bson.D{{Key: types.SliceForUpdate, Value: num}}
}

func Sort(value any) bson.D {
	return bson.D{{Key: types.Sort, Value: value}}
}

func Set(key string, value any) bson.D {
	return bson.D{{Key: types.Set, Value: bson.D{{Key: key, Value: value}}}}
}

func Unset(keys ...string) bson.D {
	value := bson.D{}
	for i := range keys {
		value = append(value, bson.E{Key: keys[i], Value: ""})
	}
	return bson.D{{Key: types.Unset, Value: value}}
}

func SetOnInsert(key string, value any) bson.D {
	return bson.D{{Key: types.SetOnInsert, Value: bson.D{{Key: key, Value: value}}}}
}

func CurrentDate(key string, value any) bson.D {
	return bson.D{{Key: types.CurrentDate, Value: bson.D{{Key: key, Value: value}}}}
}

func Inc(key string, value any) bson.D {
	return bson.D{{Key: types.Inc, Value: bson.D{{Key: key, Value: value}}}}
}

func Min(key string, value any) bson.D {
	return bson.D{{Key: types.Min, Value: bson.D{{Key: key, Value: value}}}}
}

func Max(key string, value any) bson.D {
	return bson.D{{Key: types.Max, Value: bson.D{{Key: key, Value: value}}}}
}

func Mul(key string, value any) bson.D {
	return bson.D{{Key: types.Mul, Value: bson.D{{Key: key, Value: value}}}}
}

func Rename(key string, value any) bson.D {
	return bson.D{{Key: types.Rename, Value: bson.D{{Key: key, Value: value}}}}
}
