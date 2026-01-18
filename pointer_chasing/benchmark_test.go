package pointerchasing

import (
	"runtime"
	"testing"
)

const N = 1_000_000

var sink int

func buildPointersAllocEach() []*Node {
	nodes := make([]*Node, N)
	for i := range N {
		nodes[i] = &Node{Value: i}
	}
	return nodes
}

func buildValuesPrealloc() []NodeVal {
	vals := make([]NodeVal, N)
	for i := range N {
		vals[i] = NodeVal{Value: i}
	}
	return vals
}

func iteratePointers(nodes []*Node) {
	for i := range nodes {
		sink = nodes[i].Value
	}
	runtime.KeepAlive(nodes)
}

func iterateValues(vals []NodeVal) {
	for i := range vals {
		sink = vals[i].Value
	}
	runtime.KeepAlive(vals)
}

func BenchmarkIter_Pointers_1M(b *testing.B) {
	nodes := buildPointersAllocEach()
	b.ReportAllocs()

	for b.Loop() {
		iteratePointers(nodes)
	}
}

func BenchmarkIter_Values_1M(b *testing.B) {
	vals := buildValuesPrealloc()
	b.ReportAllocs()

	for b.Loop() {
		iterateValues(vals)
	}
}

func BenchmarkBuild_PointersAllocEach_1M(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		nodes := buildPointersAllocEach()
		runtime.KeepAlive(nodes)
	}
}

func BenchmarkBuild_ValuesPrealloc_1M(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		vals := buildValuesPrealloc()
		runtime.KeepAlive(vals)
	}
}
