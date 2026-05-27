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

// const
const (
	ActionPut    = "put"
	ActionDelete = "del"
)

// KVChangeEvent is event between kie nodes, and broadcast by serf
type KVChangeEvent struct {
	Key      string
	Action   string //include: put,delete
	Labels   map[string]string
	DomainID string
	Project  string
}

func (e *KVChangeEvent) String() string { _ = "STUB: not implemented"; return "" }

// NewKVChangeEvent create a struct base on event payload
func NewKVChangeEvent(payload []byte) (*KVChangeEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Topic can be subscribe
type Topic struct {
	Labels       map[string]string `json:"-"`
	LabelsFormat string            `json:"labels,omitempty"`
	DomainID     string            `json:"domainID,omitempty"`
	Project      string            `json:"project,omitempty"`
	MatchType    string            `json:"match,omitempty"`
}

func (t *Topic) Encode() (string, error) { _ = "STUB: not implemented"; return "", nil }

// ParseTopic parse topic string to topic struct
func ParseTopic(s string) (*Topic, error) { _ = "STUB: not implemented"; return nil, nil }

// Match compare event with topic
// If the match type is set to exact in long pulling request, only update request with exactly
// the same label of pulling request will match the request and will trigger an immediate return.
//
// If the match type is not set, it will be matched when pulling request labels is equal to
// update request labels or a subset of it.
func (t *Topic) Match(event *KVChangeEvent) bool { _ = "STUB: not implemented"; return false }

// Observer represents a client polling request
type Observer struct {
	UUID  string
	Event chan *KVChangeEvent
}
