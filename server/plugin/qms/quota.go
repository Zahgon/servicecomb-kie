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

package qms

import (
	"github.com/go-chassis/go-chassis/v2/pkg/backends/quota"
)

// const
const (
	DefaultQuota   = 10000
	QuotaConfigKey = "QUOTA_CONFIG"
)

// BuildInManager read env config to max config item number, and db total usage
// it is not a centralized QMS.
type BuildInManager struct {
}

func (m *BuildInManager) SetLimit(domain, project, resourceType string, limit int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *BuildInManager) GetQuota(domain, project, resource string) (*quota.Quota, error) {
	_ = "STUB: not implemented"
	return nil,

		// GetQuotas get usage and quota
		nil
}

func (m *BuildInManager) GetQuotas(domain, project string) ([]*quota.Quota, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IncreaseUsed no use
func (m *BuildInManager) IncreaseUsed(domain, project, resource string, used int64) error {
	_ = "STUB: not implemented"

	// DecreaseUsed no use
	return nil
}

func (m *BuildInManager) DecreaseUsed(domain, project, resource string, used int64) error {
	_ = "STUB: not implemented"
	return nil
}

func newQMS(opts quota.Options) (quota.Manager, error) {
	_ = "STUB: not implemented"
	return *new(quota.Manager), nil
}

func init() {
	quota.Install("build-in", newQMS)
}
