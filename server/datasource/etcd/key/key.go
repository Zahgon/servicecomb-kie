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

package key

const (
	split      = "/"
	keyKV      = "kvs"
	keyCounter = "counter"
	keyHistory = "kv-history"
	keyTrack   = "track"
	syncer     = "syncer"
	task       = "task"
	tombstone  = "tombstone"
)

func getSyncRootKey() string { _ = "STUB: not implemented"; return "" }

func getTombstoneRootKey() string { _ = "STUB: not implemented"; return "" }

func TaskKey(domain, project, taskID string, timestamp int64) string {
	_ = "STUB: not implemented"
	return ""
}

func TombstoneKey(domain, project, resourceType, resourceID string) string {
	_ = "STUB: not implemented"
	return ""
}

func KV(domain, project, kvID string) string { _ = "STUB: not implemented"; return "" }

func KVList(domain, project string) string { _ = "STUB: not implemented"; return "" }

func Counter(name, domain string) string { _ = "STUB: not implemented"; return "" }

func His(domain, project, kvID string, updateRevision int64) string {
	_ = "STUB: not implemented"
	return ""
}

func HisList(domain, project, kvID string) string { _ = "STUB: not implemented"; return "" }

func Track(domain, project, revision, sessionID string) string {
	_ = "STUB: not implemented"
	return ""
}

func TrackList(domain, project string) string { _ = "STUB: not implemented"; return "" }
