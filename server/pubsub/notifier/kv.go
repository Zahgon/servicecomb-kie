package notifier

import (
	"sync"
	"time"

	"github.com/apache/servicecomb-kie/server/pubsub"
	"github.com/hashicorp/serf/serf"
)

// KVHandler handler serf custom event, it is singleton
type KVHandler struct {
	BatchSize          int
	BatchInterval      time.Duration
	Immediate          bool
	pendingEvents      sync.Map
	pendingEventsCount int
}

func (h *KVHandler) RunFlushTask() { _ = "STUB: not implemented"; return }

func (h *KVHandler) HandleEvent(e serf.Event) { _ = "STUB: not implemented"; return }

//never retain event, not recommended

func (h *KVHandler) mergeAndSave(ke *pubsub.KVChangeEvent) { _ = "STUB: not implemented"; return }

func (h *KVHandler) fireEvents() { _ = "STUB: not implemented"; return }

func (h *KVHandler) FindTopicAndFire(ke *pubsub.KVChangeEvent) { _ = "STUB: not implemented"; return }

//range all topics

func notifyAndRemoveObservers(value interface{}, ke *pubsub.KVChangeEvent) {
	_ = "STUB: not implemented"
	return
}

func init() {
	h := &KVHandler{
		BatchInterval: pubsub.DefaultEventBatchInterval,
		BatchSize:     pubsub.DefaultEventBatchSize,
		Immediate:     true,
	}
	pubsub.RegisterHandler(pubsub.EventKVChange, h)
	go h.RunFlushTask()
}
