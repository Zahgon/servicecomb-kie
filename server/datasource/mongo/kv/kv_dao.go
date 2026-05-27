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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/apache/servicecomb-kie/pkg/model"
	"github.com/apache/servicecomb-kie/server/datasource"
)

const (
	MsgFindKvFailed    = "find kv failed, deadline exceeded"
	FmtErrFindKvFailed = "can not find kv in %s"
)

// Dao operate data in mongodb
type Dao struct {
}

func (s *Dao) Create(ctx context.Context, kv *model.KVDoc, options ...datasource.WriteOption) (*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if syncEnable is true, will create kv with task

func create(ctx context.Context, kv *model.KVDoc) (*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// txnCreate is to start transaction when creating KV, will create task in a transaction operation
func txnCreate(ctx context.Context, kv *model.KVDoc) (*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update update key value
func (s *Dao) Update(ctx context.Context, kv *model.KVDoc, options ...datasource.WriteOption) error {
	_ = "STUB: not implemented"
	return nil
}

// if syncEnable is true, will create kv with task

func update(ctx context.Context, kv *model.KVDoc) error { _ = "STUB: not implemented"; return nil }

// txnUpdate is to start transaction when updating kV, will create task in a transaction operation
func txnUpdate(ctx context.Context, kv *model.KVDoc) error { _ = "STUB: not implemented"; return nil }

// Extract key values
func getValue(str string) string { _ = "STUB: not implemented"; return "" }

func findKV(ctx context.Context, domain string, project string, opts datasource.FindOptions) (*mongo.Cursor, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func findOneKey(ctx context.Context, filter bson.M) ([]*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Exist supports you query a key value by label map or labels id
func (s *Dao) Exist(ctx context.Context, key, project, domain string, options ...datasource.FindOption) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// FindOneAndDelete deletes one kv by id and return the deleted kv as these appeared before deletion
// domain=tenant
func (s *Dao) FindOneAndDelete(ctx context.Context, kvID, project, domain string, options ...datasource.WriteOption) (*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if syncEnable is ture, will delete kv, create task and create tombstone

func findOneAndDelete(ctx context.Context, kvID, project, domain string) (*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// txnFindOneAndDelete is to start transaction when delete KV, will create task and tombstone in a transaction operation
func txnFindOneAndDelete(ctx context.Context, kvID, project, domain string) (*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindManyAndDelete deletes multiple kvs and return the deleted kv list as these appeared before deletion
func (s *Dao) FindManyAndDelete(ctx context.Context, kvIDs []string, project, domain string, options ...datasource.WriteOption) ([]*model.KVDoc, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// if sync enable is true, will delete kvs, create tasks and tombstones

func findManyAndDelete(ctx context.Context, kvIDs []string, project, domain string) ([]*model.KVDoc, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// txnFindManyAndDelete is to start transaction when delete KVs, will create tasks and tombstones in a transaction operation
func txnFindManyAndDelete(ctx context.Context, kvIDs []string, project, domain string) ([]*model.KVDoc, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func findKeys(ctx context.Context, filter interface{}, withoutLabel bool) ([]*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// findKVByLabel get kvs by key and label
// key can be empty, then it will return all key values
// if key is given, will return 0-1 key value
func findKVByLabel(ctx context.Context, domain, labelFormat, key string, project string) ([]*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get get kv by kv id
func (s *Dao) Get(ctx context.Context, req *model.GetKVRequest) (*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Dao) GetByKey(ctx context.Context, key, project, domain string, options ...datasource.FindOption) ([]*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Dao) Total(ctx context.Context, project, domain string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// List get kv list by key and criteria
func (s *Dao) List(ctx context.Context, project, domain string, options ...datasource.FindOption) (*model.KVResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
