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
	AllOp                = "$all"
	AndOp                = "$and"
	CaseSensitiveOp      = "$caseSensitive"
	DiacriticSensitiveOp = "$diacriticSensitive"
	ElemMatchOp          = "$elemMatch"
	EqOp                 = "$eq"
	ExistsOp             = "$exists"
	ExprOp               = "$expr"
	GtOp                 = "$gt"
	GteOp                = "$gte"
	IdOp                 = "_id"
	InOp                 = "$in"
	JsonSchemaOp         = "$jsonSchema"
	LanguageOp           = "$language"
	LtOp                 = "$lt"
	LteOp                = "$lte"
	ModOp                = "$mod"
	NeOp                 = "$ne"
	NinOp                = "$nin"
	NorOp                = "$nor"
	NotOp                = "$not"
	OptionsOp            = "$options"
	OrOp                 = "$or"
	RegexOp              = "$regex"
	SearchOp             = "$search"
	SizeOp               = "$size"
	SliceOp              = "$slice"
	TextOp               = "$text"
	TypeOp               = "$type"
	WhereOp              = "$where"
)

type TextOptions struct {
	Language           string
	CaseSensitive      bool
	DiacriticSensitive bool
}
