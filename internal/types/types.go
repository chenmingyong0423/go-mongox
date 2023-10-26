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

package types

const (
	Id        = "_id"
	Set       = "$set"
	In        = "$in"
	Eq        = "$eq"
	Gt        = "$gt"
	Gte       = "$gte"
	Lt        = "$lt"
	Lte       = "$lte"
	Ne        = "$ne"
	Nin       = "$nin"
	And       = "$and"
	Not       = "$not"
	Nor       = "$nor"
	Or        = "$or"
	Exists    = "$exists"
	Type      = "$type"
	All       = "$all"
	ElemMatch = "$elemMatch"
	Size      = "$size"
)

type TestUser struct {
	Id           string `bson:"_id"`
	Name         string `bson:"name"`
	Age          int64
	UnknownField string `bson:"-"`
}

type IllegalUser struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
	Age  string
}

type UpdatedUser struct {
	Name string `bson:"name"`
	Age  int64
}

//
//type userName struct {
//	Name string `bson:"name"`
//}