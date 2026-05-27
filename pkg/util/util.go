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

package util

import (
	"github.com/go-chassis/cari/pkg/errsvc"
)

// IsEquivalentLabel compares whether two labels are equal.
// In particular, if one is nil and another is an empty map, it return true
func IsEquivalentLabel(x, y map[string]string) bool { _ = "STUB: not implemented"; return false }

// IsContainLabel compares whether x contain y
func IsContainLabel(x, y map[string]string) bool { _ = "STUB: not implemented"; return false }

func SvcErr(err error) *errsvc.Error { _ = "STUB: not implemented"; return nil }
