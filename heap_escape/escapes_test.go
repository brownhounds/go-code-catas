package heap_escapes

import (
	"runtime"
	"testing"
)

var sinkAny any
var sinkPtr *int
var sinkInt int
var sinkFn func() int
var sinkNodes []*Node
var sinkVals []Node

func BenchmarkRetPtr(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		p := RetPtr()
		sinkPtr = p
		runtime.KeepAlive(p)
	}
}

func BenchmarkRetVal(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		v := RetVal()
		sinkInt = v
		runtime.KeepAlive(v)
	}
}

func BenchmarkStoreGlobal(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		StoreGlobal()
	}
}

func BenchmarkCaptureClosure(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		fn := CaptureClosure()
		sinkFn = fn
		runtime.KeepAlive(fn)
	}
}

func BenchmarkInterfaceBoxing(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		v := InterfaceBoxing()
		sinkAny = v
		runtime.KeepAlive(v)
	}
}

func BenchmarkPrintfBoxing(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		PrintfBoxing()
	}
}

func BenchmarkAddrInLoopBad_100k(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		nodes := AddrInLoopBad(100_000)
		sinkNodes = nodes
		runtime.KeepAlive(nodes)
	}
}

func BenchmarkValuesPrealloc_100k(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		vals := ValuesPrealloc(100_000)
		sinkVals = vals
		runtime.KeepAlive(vals)
	}
}
