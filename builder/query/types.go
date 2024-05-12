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

package query

const (
	IdOp                 = "_id"
	SetOp                = "$set"
	InOp                 = "$in"
	EqOp                 = "$eq"
	GtOp                 = "$gt"
	GteOp                = "$gte"
	LtOp                 = "$lt"
	LteOp                = "$lte"
	NeOp                 = "$ne"
	NinOp                = "$nin"
	AndOp                = "$and"
	NotOp                = "$not"
	NorOp                = "$nor"
	OrOp                 = "$or"
	ExistsOp             = "$exists"
	TypeOp               = "$type"
	AllOp                = "$all"
	ElemMatchOp          = "$elemMatch"
	SizeOp               = "$size"
	UnsetOp              = "$unset"
	SetOnInsertOp        = "$setOnInsert"
	CurrentDateOp        = "$currentDate"
	IncOp                = "$inc"
	MinOp                = "$min"
	MaxOp                = "$max"
	MulOp                = "$mul"
	RenameOp             = "$rename"
	ExprOp               = "$expr"
	JsonSchemaOp         = "$jsonSchema"
	ModOp                = "$mod"
	RegexOp              = "$regex"
	OptionsOp            = "$options"
	TextOp               = "$text"
	SearchOp             = "$search"
	LanguageOp           = "$language"
	CaseSensitiveOp      = "$caseSensitive"
	DiacriticSensitiveOp = "$diacriticSensitive"
	WhereOp              = "$where"
	SliceOp              = "$slice"
	AddToSetOp           = "$addToSet"
	PopOp                = "$pop"
	PullOp               = "$pull"
	PushOp               = "$push"
	PullAllOp            = "$pullAll"
	EachOp               = "$each"
	PositionOp           = "$position"
	SliceForUpdateOp     = "$slice"
	SortOp               = "$sort"
)

type TextOptions struct {
	Language           string
	CaseSensitive      bool
	DiacriticSensitive bool
}
