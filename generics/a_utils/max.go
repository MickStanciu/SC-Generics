package a_utils

import "fmt"

type Comparable interface {
	int | float64 | float32 | int8
}

func Max[A Comparable](a A, b A) A {
	if a >= b {
		return a
	}
	return b
}

func PrintAny[T int | float32](a T) {
	fmt.Printf("%v\n", a)
}
