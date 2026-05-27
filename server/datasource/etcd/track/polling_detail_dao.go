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

package track

import (
	"context"

	"github.com/apache/servicecomb-kie/pkg/model"
)

// Dao is the implementation
type Dao struct {
}

// CreateOrUpdate create a record or update exist record
// If revision and session_id exists then update else insert
func (s *Dao) CreateOrUpdate(ctx context.Context, detail *model.PollingDetail) (*model.PollingDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPollingDetail is to get a track data
func (s *Dao) GetPollingDetail(ctx context.Context, detail *model.PollingDetail) ([]*model.PollingDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
