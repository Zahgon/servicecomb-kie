/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package v1 hold http rest v1 API
package v1

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"

	"github.com/apache/servicecomb-kie/pkg/model"
)

// KVResource has API about kv operations
type KVResource struct {
}

// Upload upload kvs
func (r *KVResource) Upload(rctx *restful.Context) { _ = "STUB: not implemented"; return }

// Post create a kv
func (r *KVResource) Post(rctx *restful.Context) { _ = "STUB: not implemented"; return }

// Put update a kv
func (r *KVResource) Put(rctx *restful.Context) { _ = "STUB: not implemented"; return }

// Get search key by kv id
func (r *KVResource) Get(rctx *restful.Context) { _ = "STUB: not implemented"; return }

// List response kv list
func (r *KVResource) List(rctx *restful.Context) { _ = "STUB: not implemented"; return }

func returnData(rctx *restful.Context, request *model.ListKVRequest) {
	_ = "STUB: not implemented"
	return
}

func isLegalWaitRequest(rctx *restful.Context, request *model.ListKVRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func watch(rctx *restful.Context, request *model.ListKVRequest, wait string) bool {
	_ = "STUB: not implemented"
	return false
}

// Delete deletes one kv by id
func (r *KVResource) Delete(rctx *restful.Context) { _ = "STUB: not implemented"; return }

// DeleteList deletes multiple kvs by ids
func (r *KVResource) DeleteList(rctx *restful.Context) { _ = "STUB: not implemented"; return }

// URLPatterns defined config operations
func (r *KVResource) URLPatterns() []restful.Route { _ = "STUB: not implemented"; return nil }
