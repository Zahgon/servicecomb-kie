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

package pubsub

import (
	"sync"
	"time"

	"github.com/hashicorp/serf/cmd/serf/command/agent"
)

var once sync.Once
var bus *Bus

// const
const (
	EventKVChange             = "kv-chg"
	DefaultEventBatchSize     = 5000
	DefaultEventBatchInterval = 500 * time.Millisecond
)

var topics sync.Map

func Topics() *sync.Map {
	_ = "STUB: not implemented"

	// Bus is message bug
	return nil
}

type Bus struct {
	agent *agent.Agent
}

// Init create serf agent
func Init() { _ = "STUB: not implemented"; return }

// splitHostPort split input string to host port
func splitHostPort(advertiseAddr string, defaultHost string, defaultPort int) (string, int) {
	_ = "STUB: not implemented"
	return "", 0
}

// Start start serf agent
func Start() { _ = "STUB: not implemented"; return }

func join(addresses []string) error { _ = "STUB: not implemented"; return nil }

// Publish send event
func Publish(event *KVChangeEvent) error { _ = "STUB: not implemented"; return nil }

// AddObserver observe key changes by (key or labels) or (key and labels)
func AddObserver(o *Observer, topic *Topic) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func RemoveObserver(uuid string, topic *Topic) { _ = "STUB: not implemented"; return }
