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

package kv

import (
	"context"

	"github.com/apache/servicecomb-kie/pkg/concurrency"
	"github.com/apache/servicecomb-kie/pkg/model"
	"github.com/apache/servicecomb-kie/server/datasource"
	"github.com/go-chassis/cari/pkg/errsvc"
)

var listSema = concurrency.NewSemaphore(concurrency.DefaultConcurrency)

func ListKV(ctx context.Context, request *model.ListKVRequest) (int64, *model.KVResponse, *errsvc.Error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// Create get latest revision from history
// and increase revision of label
// and insert key
func Create(ctx context.Context, kv *model.KVDoc) (*model.KVDoc, *errsvc.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//check whether the project has certain labels or not

func completeKV(kv *model.KVDoc, revision int64) error { _ = "STUB: not implemented"; return nil }

func Upload(ctx context.Context, request *model.UploadKVRequest) *model.DocRespOfUpload {
	_ = "STUB: not implemented"
	return nil
}

func appendFailedKVResult(err *errsvc.Error, kv *model.KVDoc, result *model.DocRespOfUpload) {
	_ = "STUB: not implemented"
	return
}

func appendAbortFailedKVResult(kvs []*model.KVDoc, result *model.DocRespOfUpload) {
	_ = "STUB: not implemented"
	return
}

func Publish(kv *model.KVDoc) { _ = "STUB: not implemented"; return }

// Update update key value and add new revision
func Update(ctx context.Context, kv *model.UpdateKVRequest) (*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindOneAndDelete(ctx context.Context, kvID string, project, domain string) (*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FindManyAndDelete(ctx context.Context, kvIDs []string, project, domain string) ([]*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Get(ctx context.Context, req *model.GetKVRequest) (*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func List(ctx context.Context, project, domain string, options ...datasource.FindOption) (*model.KVResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Exist(ctx context.Context, key, project, domain string, labels map[string]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func GetByKey(ctx context.Context, key, project, domain string, labels map[string]string) ([]*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
