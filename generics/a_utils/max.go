package a_utils

type Comparable interface {
	int | float64 | float32 | int8
}

func Max[A Comparable](a A, b A) A {
	if a >= b {
		return a
	}
	return b
}
