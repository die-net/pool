// Package mappool provides a sync.Pool to repeatedly use a map[K]V, avoiding unnecessary allocations.
package mappool

import "sync"

// Pool defines a sync.Pool of map[K]V, where K is a comparable type, and V is any type.
type Pool[K comparable, V any] struct {
	pool sync.Pool
}

// New creates a new Pool[K,V] holding map[K]V with a minimum capacity of size for
// reuse.  A non-zero size will make sure new maps don't need to allocate
// more than once to be able to hold the specified capacity.
func New[K comparable, V any](size int) *Pool[K,V] {
	return &Pool[K,V]{
		pool: sync.Pool{
			New: func() any {
				m := make(map[K]V, size)
				return &m
			},
		},
	}
}

// Get returns a map[K]V from the pool.  The map[K]V must be later Put to
// reuse it in the future; it will be garbage collected if not.
func (p *Pool[K,V]) Get() map[K]V {
	m := p.pool.Get().(*map[K]V)
	return *m
}

// Put clears and returns a map[K]V to the pool to be used by a future Get;
// this is only safe if the caller doesn't attempt to use the map for
// anything else.  Somewhat oddly, Put is passed a pointer to avoid a
// small allocation.
func (p *Pool[K,V]) Put(m *map[K]V) {
	clear(*m)

	p.pool.Put(m)
}
