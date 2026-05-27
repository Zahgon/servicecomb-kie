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

package datasource

import (
	"time"
)

const DefaultTimeout = 60 * time.Second

type Config struct {
}

// NewDefaultFindOpts return default options
func NewDefaultFindOpts() FindOptions { _ = "STUB: not implemented"; return *new(FindOptions) }

// NewDefaultWriteOptions return default options
func NewDefaultWriteOptions() WriteOptions { _ = "STUB: not implemented"; return *new(WriteOptions) }

// NewWriteOptions return options with write option
func NewWriteOptions(option ...WriteOption) WriteOptions {
	_ = "STUB: not implemented"
	return *new(WriteOptions)
}

// WriteOptions is option for create ,update and delete kv
type WriteOptions struct {
	SyncEnable bool
}

// FindOptions is option to find key value
type FindOptions struct {
	ExactLabels bool
	Status      string
	Depth       int
	ID          string
	Key         string
	Value       string
	Labels      map[string]string
	LabelFormat string
	ClearLabel  bool
	Timeout     time.Duration
	// Offset the offset of the response, start at 0
	Offset int64
	// Limit the page size of the response, dot not paging if limit=0
	Limit         int64
	CaseSensitive bool
}

// WriteOption is functional option to create, update and delete kv
type WriteOption func(*WriteOptions)

// FindOption is functional option to find key value
type FindOption func(*FindOptions)

// WithSync indicates that the synchronization function is on
func WithSync(enabled bool) WriteOption { _ = "STUB: not implemented"; return *new(WriteOption) }

// WithCaseSensitive tell model service whether to match case of letters or not.
func WithCaseSensitive() FindOption { _ = "STUB: not implemented"; return *new(FindOption) }

// WithExactLabels tell model service to return only one kv matches the labels
func WithExactLabels() FindOption { _ = "STUB: not implemented"; return *new(FindOption) }

// WithID find by kvID
func WithID(id string) FindOption { _ = "STUB: not implemented"; return *new(FindOption) }

// WithKey find by key
func WithKey(key string) FindOption { _ = "STUB: not implemented"; return *new(FindOption) }

// WithValue find by value
func WithValue(value string) FindOption { _ = "STUB: not implemented"; return *new(FindOption) }

// WithStatus enabled/disabled
func WithStatus(status string) FindOption { _ = "STUB: not implemented"; return *new(FindOption) }

// WithTimeout will return err if execution take too long
func WithTimeout(d time.Duration) FindOption { _ = "STUB: not implemented"; return *new(FindOption) }

// WithLabels find kv by labels
func WithLabels(labels map[string]string) FindOption {
	_ = "STUB: not implemented"
	return *new(FindOption)
}

// WithLabelFormat find kv by label string
func WithLabelFormat(label string) FindOption { _ = "STUB: not implemented"; return *new(FindOption) }

// WithLimit tells service paging limit
func WithLimit(l int64) FindOption { _ = "STUB: not implemented"; return *new(FindOption) }

// WithOffset tells service paging offset
func WithOffset(os int64) FindOption { _ = "STUB: not implemented"; return *new(FindOption) }
