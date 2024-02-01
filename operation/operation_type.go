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

package operation

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type OpType string

const (
	OpTypeBeforeInsert OpType = "beforeInsert"
	OpTypeAfterInsert  OpType = "afterInsert"
	OpTypeBeforeUpdate OpType = "beforeUpdate"
	OpTypeAfterUpdate  OpType = "afterUpdate"
	OpTypeBeforeDelete OpType = "beforeDelete"
	OpTypeAfterDelete  OpType = "afterDelete"
	OpTypeBeforeUpsert OpType = "beforeUpsert"
	OpTypeAfterUpsert  OpType = "afterUpsert"
	OpTypeBeforeFind   OpType = "beforeFind"
	OpTypeAfterFind    OpType = "afterFind"
)

type OpContext struct {
	Col *mongo.Collection `opt:"-"`
	Doc any
	// filter also can be used as query query
	Filter      any
	Updates     any
	Replacement any
}

func (opCtx OpContext) GetPayload(opType OpType) any {
	switch opType {
	case OpTypeBeforeInsert, OpTypeAfterInsert:
		return opCtx.Doc
	case OpTypeBeforeUpdate, OpTypeAfterUpdate:
		return opCtx.Updates
	case OpTypeBeforeDelete, OpTypeAfterDelete:
		return opCtx.Filter
	case OpTypeBeforeUpsert, OpTypeAfterUpsert:
		return opCtx.Replacement
	}
	return nil
}
