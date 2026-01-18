package heap_escapes

import "fmt"

// RetPtr returns a pointer to a local variable.
// In Go this is safe: the compiler makes x escape to the heap.
func RetPtr() *int {
	x := 42
	return &x
}

// RetVal returns a value; x can stay on the stack.
func RetVal() int {
	x := 42
	return x
}

// StoreGlobal forces escape because the address is stored globally.
var GlobalPtr *int

func StoreGlobal() {
	x := 7
	GlobalPtr = &x
}

// CaptureClosure forces escape because captured variables must outlive the function call.
func CaptureClosure() func() int {
	x := 123
	return func() int { return x }
}

// InterfaceBoxing often forces escape because values stored in interface{} need stable storage.
func InterfaceBoxing() any {
	x := 99
	return x
}

// PrintfBoxing usually escapes because passing to fmt causes interface conversions.
// (Great example of "accidental escapes" from logging.)
func PrintfBoxing() {
	x := 11
	fmt.Sprintf("%d", x)
}

// AddrInLoopBad: taking address of a loop local each iteration -> escape per element.
type Node struct{ V int }

func AddrInLoopBad(n int) []*Node {
	out := make([]*Node, 0, n)
	for i := 0; i < n; i++ {
		x := Node{V: i}
		out = append(out, &x)
	}
	return out
}

// ValuesPrealloc: same data but stored contiguously, minimal escaping.
// Note: returning the slice means the backing array lives on the heap,
// but you avoid per-element heap allocations and pointer graph overhead.
func ValuesPrealloc(n int) []Node {
	out := make([]Node, n)
	for i := 0; i < n; i++ {
		out[i] = Node{V: i}
	}
	return out
}
