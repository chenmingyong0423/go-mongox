// Copyright 2024 chenmingyong0423

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
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	StageLookUpOp         = "$lookup"
	AddFieldsOp           = "$addFields"
	SetOp                 = "$set"
	AddOp                 = "$add"
	SumOp                 = "$sum"
	BucketOp              = "$bucket"
	BucketAutoOp          = "$bucketAuto"
	ContactOp             = "$concat"
	PushOp                = "$push"
	MatchOp               = "$match"
	GroupOp               = "$group"
	DateToStringOp        = "$dateToString"
	MultiplyOp            = "$multiply"
	AvgOp                 = "$avg"
	SortOp                = "$sort"
	ProjectOp             = "$project"
	SkipOp                = "$skip"
	UnwindOp              = "$unwind"
	ReplaceWithOp         = "$replaceWith"
	ArrayToObjectOp       = "$arrayToObject"
	FacetOp               = "$facet"
	SortByCountOp         = "$sortByCount"
	CountOp               = "$count"
	SubtractOp            = "$subtract"
	DivideOp              = "$divide"
	ModOp                 = "$mod"
	EqOp                  = "$eq"
	NeOp                  = "$ne"
	GtOp                  = "$gt"
	GteOp                 = "$gte"
	LtOp                  = "$lt"
	LteOp                 = "$lte"
	AndOp                 = "$and"
	NotOp                 = "$not"
	OrOp                  = "$or"
	ConcatOp              = "$concat"
	SubstrBytesOp         = "$substrBytes"
	ToLowerOp             = "$toLower"
	ToUpperOp             = "$toUpper"
	ArrayElemAtOp         = "$arrayElemAt"
	ConcatArraysOp        = "$concatArrays"
	SizeOp                = "$size"
	SliceOp               = "$slice"
	DayOfMonthOp          = "$dayOfMonth"
	DayOfWeekOp           = "$dayOfWeek"
	DayOfYearOp           = "$dayOfYear"
	YearOp                = "$year"
	MonthOp               = "$month"
	WeekOp                = "$week"
	CondOp                = "$cond"
	IfNullOp              = "$ifNull"
	SwitchOp              = "$switch"
	FirstOp               = "$first"
	LastOp                = "$last"
	MinOp                 = "$min"
	MaxOp                 = "$max"
	MapOp                 = "$map"
	FilterOp              = "$filter"
	InputOp               = "input"
	AsOp                  = "as"
	InOp                  = "in"
	LimitOp               = "$limit"
	LIMIT                 = "limit"
	CondWithoutOperatorOp = "cond"
)

type LookUpOptions struct {
	LocalField   string
	ForeignField string
	Let          bson.D
	Pipeline     mongo.Pipeline
}
