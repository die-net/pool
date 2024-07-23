package slicepool

import (
	"testing"
)

func TestPool(t *testing.T) {
        p := New[int](1)
        sl := p.Get()
        if cap(sl) != 1 {
                t.Error("Get didn't return slice of cap 1")
        }
        sl = append(sl, 1, 2)
        p.Put(&sl)

        sl = p.Get()
        if cap(sl) != 2 {
                t.Error("Get didn't reuse slice of cap 2")
        }
        if len(sl) > 0 {
                t.Error("Put didn't reset slice")
        }
        sl = sl[:2:2]
        if sl[0] != 0 || sl[1] != 0 {
        	t.Error("Put didn't clear slice")
        }
}

func BenchmarkNewPool(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		var p *Pool[int64]
		for pb.Next() {
			p = New[int64](512)
			_ = p
		}
	})
}

func BenchmarkGetPut(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		bp := New[int64](512)

		var b []int64
		for pb.Next() {
			b = bp.Get()
			_ = b
			bp.Put(&b)
		}
	})
}

func BenchmarkMakeSlice(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		var b []int64
		for pb.Next() {
			b = make([]int64, 0, 512)
			_ = b
		}
	})
}
