package generator

import "sync/atomic"

type AtomicIDGenerator struct {
	counter int64
}

func NewAtomicIDGenerator() *AtomicIDGenerator {
	return &AtomicIDGenerator{
		counter: 100000,
	}
}

func (g *AtomicIDGenerator) NextID() int64 {
	return atomic.AddInt64(&g.counter, 1)
}

// compiletime interafce validation

var _ IDGenerator = (*AtomicIDGenerator)(nil)
