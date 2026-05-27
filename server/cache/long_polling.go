package cache

import (
	"sync"

	"github.com/apache/servicecomb-kie/pkg/model"
	"github.com/go-chassis/cari/pkg/errsvc"
)

var pollingCache = &LongPollingCache{}

// LongPollingCache exchange space for time
type LongPollingCache struct {
	m sync.Map
}
type DBResult struct {
	KVs *model.KVResponse
	Err *errsvc.Error
	Rev int64
}

func CachedKV() *LongPollingCache { _ = "STUB: not implemented"; return nil }

// Read reads the cached query result
// only need to filter by labels if match pattern is exact
func (c *LongPollingCache) Read(topic string) (int64, *model.KVResponse, *errsvc.Error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (c *LongPollingCache) Write(topic string, r *DBResult) { _ = "STUB: not implemented"; return }
