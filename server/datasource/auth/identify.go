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

package auth

import (
	"context"
	"errors"

	rbacmodel "github.com/go-chassis/cari/rbac"
)

const (
	RootName = "root"
)

var ErrNoRoles = errors.New("no role found in token")

func Identify(ctx context.Context) (*rbacmodel.Account, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accountExist(ctx context.Context, user string) error {
	_ = "STUB: not implemented"
	// if root should pass, cause of root initialization
	return nil
}
