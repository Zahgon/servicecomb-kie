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

package config

// Configurations is kie config items
var Configurations = &Config{}

// Init initiate config files
func Init() error { _ = "STUB: not implemented"; return nil }

// GetDB return db configs
func GetDB() DB {
	_ = "STUB: not implemented"
	return *

	// GetRBAC return rbac config
	new(DB)
}

func GetRBAC() RBAC { _ = "STUB: not implemented"; return *new(RBAC) }

// GetSync return sync config
func GetSync() Sync { _ = "STUB: not implemented"; return *new(Sync) }
