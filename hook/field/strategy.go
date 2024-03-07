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

package field

import (
	"github.com/chenmingyong0423/go-mongox/hook"
	"github.com/chenmingyong0423/go-mongox/operation"
)

var strategies = map[operation.OpType]func(doc any) error{
	operation.OpTypeBeforeInsert: BeforeInsert,
	operation.OpTypeBeforeUpdate: BeforeUpdate,
	operation.OpTypeBeforeUpsert: BeforeUpsert,
}

func BeforeInsert(doc any) error {
	if tsh, ok := doc.(hook.DefaultModelHook); ok {
		tsh.DefaultId()
		tsh.DefaultCreatedAt()
	}
	return nil
}

func BeforeUpdate(doc any) error {
	if tsh, ok := doc.(hook.DefaultModelHook); ok {
		tsh.DefaultUpdatedAt()
	}
	return nil
}

func BeforeUpsert(doc any) error {
	if tsh, ok := doc.(hook.DefaultModelHook); ok {
		tsh.DefaultId()
		tsh.DefaultCreatedAt()
		tsh.DefaultUpdatedAt()
	}
	return nil
}
