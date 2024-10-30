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
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	AbsOp                 = "$abs"
	AddOp                 = "$add"
	AndOp                 = "$and"
	ArrayElemAtOp         = "$arrayElemAt"
	ArrayToObjectOp       = "$arrayToObject"
	AsOp                  = "as"
	AvgOp                 = "$avg"
	BranchesOp            = "branches"
	CaseOp                = "case"
	CeilOp                = "$ceil"
	ConcatArraysOp        = "$concatArrays"
	ConcatOp              = "$concat"
	CondOp                = "$cond"
	CondWithoutOperatorOp = "cond"
	ContactOp             = "$concat"
	DateOp                = "date"
	DateToStringOp        = "$dateToString"
	DayOfMonthOp          = "$dayOfMonth"
	DayOfWeekOp           = "$dayOfWeek"
	DayOfYearOp           = "$dayOfYear"
	DefaultCaseOp         = "default"
	DivideOp              = "$divide"
	EqOp                  = "$eq"
	ExpOp                 = "$exp"
	FilterOp              = "$filter"
	FirstOp               = "$first"
	FloorOp               = "$floor"
	FormatOp              = "format"
	GtOp                  = "$gt"
	GteOp                 = "$gte"
	IfNullOp              = "$ifNull"
	InOp                  = "in"
	InputOp               = "input"
	LastOp                = "$last"
	LIMIT                 = "limit"
	LnOp                  = "$ln"
	Log10Op               = "$log10"
	LogOp                 = "$log"
	LtOp                  = "$lt"
	LteOp                 = "$lte"
	MapOp                 = "$map"
	MaxOp                 = "$max"
	MinOp                 = "$min"
	ModOp                 = "$mod"
	MonthOp               = "$month"
	MultiplyOp            = "$multiply"
	NeOp                  = "$ne"
	NotOp                 = "$not"
	OnNullOp              = "onNull"
	OrOp                  = "$or"
	PowOp                 = "$pow"
	PushOp                = "$push"
	RoundOp               = "$round"
	SizeOp                = "$size"
	SliceOp               = "$slice"
	SqrtOp                = "$sqrt"
	SubstrBytesOp         = "$substrBytes"
	SubtractOp            = "$subtract"
	SumOp                 = "$sum"
	SwitchOp              = "$switch"
	ThenOp                = "then"
	TimezoneOp            = "timezone"
	ToLowerOp             = "$toLower"
	ToUpperOp             = "$toUpper"
	TruncOp               = "$trunc"
	WeekOp                = "$week"
	YearOp                = "$year"
)

// Stages
const (
	StageAddFieldsOp   = "$addFields"
	StageBoundariesOp  = "boundaries"
	StageBucketAutoOp  = "$bucketAuto"
	StageBucketOp      = "$bucket"
	StageBucketsOp     = "buckets"
	StageCountOp       = "$count"
	StageDefaultOp     = "default"
	StageFacetOp       = "$facet"
	StageGranularityOp = "granularity"
	StageGroupByOp     = "groupBy"
	StageGroupOp       = "$group"
	StageLimitOp       = "$limit"
	StageLookUpOp      = "$lookup"
	StageMatchOp       = "$match"
	StageOutputOp      = "output"
	StageProjectOp     = "$project"
	StageReplaceWithOp = "$replaceWith"
	StageSetOp         = "$set"
	StageSkipOp        = "$skip"
	StageSortByCountOp = "$sortByCount"
	StageSortOp        = "$sort"
	StageUnwindOp      = "$unwind"
)

type BucketAutoOptions struct {
	Output      any
	Granularity string
}

type BucketOptions struct {
	DefaultKey any
	Output     any
}

type CaseThen struct {
	Case any
	Then any
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

type LookUpOptions struct {
	LocalField   string
	ForeignField string
	Let          bson.D
	Pipeline     mongo.Pipeline
}

type UnWindOptions struct {
	IncludeArrayIndex          string
	PreserveNullAndEmptyArrays bool
}
