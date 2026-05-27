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

package etcd

import (
	"github.com/apache/servicecomb-kie/server/datasource"
	rbacdao "github.com/apache/servicecomb-kie/server/datasource/rbac"
)

type Broker struct {
}

func NewFrom(c *datasource.Config) (datasource.Broker, error) {
	_ = "STUB: not implemented"
	return *new(datasource.Broker), nil
}

func (*Broker) GetRevisionDao() datasource.RevisionDao {
	_ = "STUB: not implemented"
	return *new(datasource.RevisionDao)
}

func (*Broker) GetKVDao() datasource.KVDao {
	_ = "STUB: not implemented"
	return *new(datasource.KVDao)
}

func (*Broker) GetHistoryDao() datasource.HistoryDao {
	_ = "STUB: not implemented"
	return *new(datasource.HistoryDao)
}

func (*Broker) GetTrackDao() datasource.TrackDao {
	_ = "STUB: not implemented"
	return *new(datasource.TrackDao)
}

func (*Broker) GetRbacDao() rbacdao.Dao { _ = "STUB: not implemented"; return *new(rbacdao.Dao) }

func init() {
	datasource.RegisterPlugin("etcd", NewFrom)
	datasource.RegisterPlugin("embedded_etcd", NewFrom)
}
