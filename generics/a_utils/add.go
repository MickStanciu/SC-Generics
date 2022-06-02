package a_utils

// AddThings custom add function
func AddThings[T int | string](a, b T) T {
	return a + b
}
