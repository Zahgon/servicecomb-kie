package kv

import (
	"context"
	"regexp"
	"sync"
	"time"

	"github.com/little-cui/etcdadpt"
	goCache "github.com/patrickmn/go-cache"
	"go.etcd.io/etcd/api/v3/mvccpb"

	"github.com/apache/servicecomb-kie/pkg/model"
	"github.com/apache/servicecomb-kie/server/datasource"
)

func Init() { _ = "STUB: not implemented"; return }

var kvCache *Cache

const (
	prefixKvs            = "kvs"
	cacheExpirationTime  = 10 * time.Minute
	cacheCleanupInterval = 11 * time.Minute
	etcdWatchTimeout     = 1 * time.Hour
	backOffMinInterval   = 5 * time.Second
)

type Cache struct {
	timeOut    time.Duration
	client     etcdadpt.Client
	revision   int64
	kvIDCache  sync.Map
	kvDocCache *goCache.Cache
}

func NewKvCache() *Cache { _ = "STUB: not implemented"; return nil }

func Enabled() bool { _ = "STUB: not implemented"; return false }

type CacheSearchReq struct {
	Domain  string
	Project string
	Opts    *datasource.FindOptions
	Regex   *regexp.Regexp
}

func (kc *Cache) Refresh(ctx context.Context) { _ = "STUB: not implemented"; return }

func (kc *Cache) listWatch(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (kc *Cache) watch(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (kc *Cache) list(ctx context.Context) (*etcdadpt.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kc *Cache) watchCallBack(message string, rsp *etcdadpt.Response) error {
	_ = "STUB: not implemented"
	return nil
}

func (kc *Cache) cachePut(rsp *etcdadpt.Response) { _ = "STUB: not implemented"; return }

func (kc *Cache) cacheDelete(rsp *etcdadpt.Response) { _ = "STUB: not implemented"; return }

func (kc *Cache) LoadKvIDSet(cacheKey string) (*sync.Map, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (kc *Cache) StoreKvIDSet(cacheKey string, kvIds *sync.Map) { _ = "STUB: not implemented"; return }

func (kc *Cache) LoadKvDoc(kvID string) (*model.KVDoc, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (kc *Cache) StoreKvDoc(kvID string, kvDoc *model.KVDoc) { _ = "STUB: not implemented"; return }

func (kc *Cache) DeleteKvDoc(kvID string) { _ = "STUB: not implemented"; return }

func Search(ctx context.Context, req *CacheSearchReq) (*model.KVResponse, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (kc *Cache) getKvFromEtcd(ctx context.Context, req *CacheSearchReq, kvIdsLeft []string) ([]*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isMatch(req *CacheSearchReq, doc *model.KVDoc) bool { _ = "STUB: not implemented"; return false }

func (kc *Cache) GetKvDoc(kv *mvccpb.KeyValue) (*model.KVDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (kc *Cache) GetCacheKey(domain, project string, labels map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}
