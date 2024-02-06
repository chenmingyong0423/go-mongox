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

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Id                 = "_id"
	Set                = "$set"
	In                 = "$in"
	Eq                 = "$eq"
	Gt                 = "$gt"
	Gte                = "$gte"
	Lt                 = "$lt"
	Lte                = "$lte"
	Ne                 = "$ne"
	Nin                = "$nin"
	And                = "$and"
	Not                = "$not"
	Nor                = "$nor"
	Or                 = "$or"
	Exists             = "$exists"
	Type               = "$type"
	All                = "$all"
	ElemMatch          = "$elemMatch"
	Size               = "$size"
	Unset              = "$unset"
	SetOnInsert        = "$setOnInsert"
	CurrentDate        = "$currentDate"
	Inc                = "$inc"
	Min                = "$min"
	Max                = "$max"
	Mul                = "$mul"
	Rename             = "$rename"
	Expr               = "$expr"
	JsonSchema         = "$jsonSchema"
	Mod                = "$mod"
	Regex              = "$regex"
	Options            = "$options"
	Text               = "$text"
	Search             = "$search"
	Language           = "$language"
	CaseSensitive      = "$caseSensitive"
	DiacriticSensitive = "$diacriticSensitive"
	Where              = "$where"
	Slice              = "$slice"
	AddToSet           = "$addToSet"
	Pop                = "$pop"
	Pull               = "$pull"
	Push               = "$push"
	PullAll            = "$pullAll"
	Each               = "$each"
	Position           = "$position"
	SliceForUpdate     = "$slice"
	Sort               = "$sort"
)

const (
	AggregationStageAddFields      = "$addFields"
	AggregationStageSet            = "$set"
	AggregationAdd                 = "$add"
	AggregationSum                 = "$sum"
	AggregationStageBucket         = "$bucket"
	AggregationStageBucketAuto     = "$bucketAuto"
	AggregationContact             = "$concat"
	AggregationPush                = "$push"
	AggregationStageMatch          = "$match"
	AggregationStageGroup          = "$group"
	AggregationDateToString        = "$dateToString"
	AggregationMultiply            = "$multiply"
	AggregationAvg                 = "$avg"
	AggregationStageSort           = "$sort"
	AggregationStageProject        = "$project"
	AggregationStageLimit          = "$limit"
	AggregationStageSkip           = "$skip"
	AggregationStageUnwind         = "$unwind"
	AggregationStageReplaceWith    = "$replaceWith"
	AggregationArrayToObject       = "$arrayToObject"
	AggregationStageFacet          = "$facet"
	AggregationStageSortByCount    = "$sortByCount"
	AggregationStageCount          = "$count"
	AggregationSubtract            = "$subtract"
	AggregationDivide              = "$divide"
	AggregationMod                 = "$mod"
	AggregationEq                  = "$eq"
	AggregationNe                  = "$ne"
	AggregationGt                  = "$gt"
	AggregationGte                 = "$gte"
	AggregationLt                  = "$lt"
	AggregationLte                 = "$lte"
	AggregationAnd                 = "$and"
	AggregationNot                 = "$not"
	AggregationOr                  = "$or"
	AggregationConcat              = "$concat"
	AggregationSubstrBytes         = "$substrBytes"
	AggregationToLower             = "$toLower"
	AggregationToUpper             = "$toUpper"
	AggregationArrayElemAt         = "$arrayElemAt"
	AggregationConcatArrays        = "$concatArrays"
	AggregationSize                = "$size"
	AggregationSlice               = "$slice"
	AggregationDayOfMonth          = "$dayOfMonth"
	AggregationDayOfWeek           = "$dayOfWeek"
	AggregationDayOfYear           = "$dayOfYear"
	AggregationYear                = "$year"
	AggregationMonth               = "$month"
	AggregationWeek                = "$week"
	AggregationCond                = "$cond"
	AggregationIfNull              = "$ifNull"
	AggregationSwitch              = "$switch"
	AggregationFirst               = "$first"
	AggregationLast                = "$last"
	AggregationMin                 = "$min"
	AggregationMax                 = "$max"
	AggregationMap                 = "$map"
	AggregationFilter              = "$filter"
	AggregationInput               = "input"
	AggregationAs                  = "as"
	AggregationIn                  = "in"
	AggregationLimit               = "limit"
	AggregationCondWithoutOperator = "cond"
)

const (
	GroupBy     = "groupBy"
	Boundaries  = "boundaries"
	Default     = "default"
	Output      = "output"
	Buckets     = "buckets"
	Granularity = "granularity"
	Date        = "date"
	Format      = "format"
	Timezone    = "timezone"
	OnNull      = "onNull"
	Branches    = "branches"
	Case        = "case"
	Then        = "then"
	DefaultCase = "default"
)

type TestUser struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Age          int64
	UnknownField string    `bson:"-"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

func (tu *TestUser) DefaultCreatedAt() {
	if tu.CreatedAt.IsZero() {
		tu.CreatedAt = time.Now().Local()
	}
}

func (tu *TestUser) DefaultUpdatedAt() {
	tu.UpdatedAt = time.Now().Local()
}

type TestTempUser struct {
	Id           string `bson:"_id"`
	Name         string `bson:"name"`
	Age          int64
	UnknownField string `bson:"-"`
}

type IllegalUser struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Age  string
}

type UpdatedUser struct {
	Name string `bson:"name"`
	Age  int64
}

type UserName struct {
	Name string `bson:"name"`
}

type UnWindOptions struct {
	IncludeArrayIndex          string
	PreserveNullAndEmptyArrays bool
}

type BucketOptions struct {
	DefaultKey any
	Output     any
}

type BucketAutoOptions struct {
	Output      any
	Granularity string
}

type DateToStringOptions struct {
	Format   string
	Timezone string
	OnNull   any
}

type FilterOptions struct {
	As    string
	Limit int64
}

type TextOptions struct {
	Language           string
	CaseSensitive      bool
	DiacriticSensitive bool
}

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type CaseThen struct {
	Case any
	Then any
}
