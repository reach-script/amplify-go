package errors

import "runtime"

type Frame uintptr

func callers(skip int) []Frame {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip, pcs[:])

	var stack []Frame
	for _, pc := range pcs[0:n] {
		stack = append(stack, Frame(pc))
	}
	return stack
}
