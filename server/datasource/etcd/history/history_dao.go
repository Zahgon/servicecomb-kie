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

package history

import (
	"context"

	"github.com/apache/servicecomb-kie/pkg/model"
	"github.com/apache/servicecomb-kie/server/datasource"
)

// Dao is the implementation
type Dao struct {
}

// GetHistory get all history by label id
func (s *Dao) GetHistory(ctx context.Context, kvID, project, domain string, options ...datasource.FindOption) (*model.KVResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pagingResult(histories []*model.KVDoc, offset, limit int64) []*model.KVDoc {
	_ = "STUB: not implemented"
	return nil
}

// AddHistory add kv history
func (s *Dao) AddHistory(ctx context.Context, kv *model.KVDoc) error {
	_ = "STUB: not implemented"
	return nil
}

// DelayDeletionTime add delete time to all revisions of the kv,
// thus these revisions will be automatically deleted by TTL index.
// TODO support delay deletion
func (s *Dao) DelayDeletionTime(ctx context.Context, kvIDs []string, project, domain string) error {
	_ = "STUB: not implemented"
	return nil
}

// historyRotate delete historical versions for a key that exceeds the limited number
func (s *Dao) historyRotate(ctx context.Context, kvID, project, domain string) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteMany(ctx context.Context, kvs []*model.KVDoc) error {
	_ = "STUB: not implemented"
	return nil
}
