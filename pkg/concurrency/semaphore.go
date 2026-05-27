package concurrency

import "math"

const (
	DefaultConcurrency = 500
	MaxConcurrency     = math.MaxUint16
)

// Semaphore ctl the max concurrency
type Semaphore struct {
	tickets chan bool
}

// NewSemaphore accept concurrency number, not more than 65535
func NewSemaphore(concurrency int) *Semaphore { _ = "STUB: not implemented"; return nil }

func (b *Semaphore) Acquire() {
	_ = "STUB: not implemented"

	// Release return back signal
	return
}

func (b *Semaphore) Release() { _ = "STUB: not implemented"; return }
