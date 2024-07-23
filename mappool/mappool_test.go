package mappool

import (
	"testing"
)

func TestPool(t *testing.T) {
	p := New[int,bool](0)
	m := p.Get()
	m[123] = true
	p.Put(&m)
	m = p.Get()
	if len(m) > 0 {
		t.Error("Put didn't clear map")
	}
}

func BenchmarkNewPool(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		var p *Pool[int64,bool]
		for pb.Next() {
			p = New[int64,bool](512)
			_ = p
		}
	})
}

func BenchmarkGetPut(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		bp := New[int64,bool](512)

		var m map[int64]bool
		for pb.Next() {
			m = bp.Get()
			bp.Put(&m)
		}
	})
}

func BenchmarkMakeMap(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		var m map[int64]bool
		for pb.Next() {
			m = make(map[int64]bool, 512)
			_ = m
		}
	})
}
