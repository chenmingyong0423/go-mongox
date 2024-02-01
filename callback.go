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

package mongox

import (
	"github.com/chenmingyong0423/go-mongox/callback"
	"github.com/chenmingyong0423/go-mongox/operation"
)

func RegisterBeforeInsert(name string, cb callback.CbFn) {
	callback.Callbacks.Register(operation.OpTypeBeforeInsert, name, cb)
}

func RegisterAfterInsert(name string, cb callback.CbFn) {
	callback.Callbacks.Register(operation.OpTypeAfterInsert, name, cb)
}

func RegisterBeforeUpdate(name string, cb callback.CbFn) {
	callback.Callbacks.Register(operation.OpTypeBeforeUpdate, name, cb)
}

func RegisterAfterUpdate(name string, cb callback.CbFn) {
	callback.Callbacks.Register(operation.OpTypeAfterUpdate, name, cb)
}

func RegisterBeforeDelete(name string, cb callback.CbFn) {
	callback.Callbacks.Register(operation.OpTypeBeforeDelete, name, cb)
}

func RegisterAfterDelete(name string, cb callback.CbFn) {
	callback.Callbacks.Register(operation.OpTypeAfterDelete, name, cb)
}

func RegisterBeforeUpsert(name string, cb callback.CbFn) {
	callback.Callbacks.Register(operation.OpTypeBeforeUpsert, name, cb)
}

func RegisterAfterUpsert(name string, cb callback.CbFn) {
	callback.Callbacks.Register(operation.OpTypeAfterUpsert, name, cb)
}

func RegisterBeforeFind(name string, cb callback.CbFn) {
	callback.Callbacks.Register(operation.OpTypeBeforeFind, name, cb)
}

func RegisterAfterFind(name string, cb callback.CbFn) {
	callback.Callbacks.Register(operation.OpTypeAfterFind, name, cb)
}
