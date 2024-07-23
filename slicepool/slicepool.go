// Package slicepool provides a sync.Pool to repeatedly use a []T, avoiding unnecessary allocations.
package slicepool

import "sync"

// Pool defines a sync.Pool of []T, where T is any type.
type Pool[T any] struct {
	pool sync.Pool
}

// New creates a new Pool[T] holding []T with a minimum capacity of size for
// reuse.  A non-zero size will make sure new slices don't need to allocate
// more than once to be able to hold the specified capacity.
func New[T any](size int) *Pool[T] {
	return &Pool[T]{
		pool: sync.Pool{
			New: func() any {
				sl := make([]T, 0, size)
				return &sl
			},
		},
	}
}

// Get returns a []T from the pool.  The []T must be later Put to reuse it
// in the future; it will be garbage collected if not.
func (p *Pool[T]) Get() []T {
	sl := p.pool.Get().(*[]T)
	return *sl
}

// Put clears, resets to 0 length, and returns a []T to the pool, to be used
// by a future Get; this is only safe if the caller doesn't attempt to use
// the slice for anything else.  Somewhat oddly, Put is passed a pointer to
// avoid a small allocation.
func (p *Pool[T]) Put(sl *[]T) {
	clear(*sl)
	*sl = (*sl)[:0]

	p.pool.Put(sl)
}
