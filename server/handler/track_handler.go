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

package handler

import (
	"github.com/go-chassis/go-chassis/v2/core/handler"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/openlog"
)

// const of noop auth handler
const (
	TrackHandlerName = "track-handler"
)

// TrackHandler tracks polling data
type TrackHandler struct{}

// Handle set local attribute to http request
func (h *TrackHandler) Handle(chain *handler.Chain, inv *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

func newTrackHandler() handler.Handler {
	_ = "STUB: not implemented"
	return *

	// Name is handler name
	new(handler.Handler)
}

func (h *TrackHandler) Name() string { _ = "STUB: not implemented"; return "" }

func init() {
	if err := handler.RegisterHandler(TrackHandlerName, newTrackHandler); err != nil {
		openlog.Fatal("register handler failed: " + err.Error())
	}
}
