package b_map

// Map accepts a value A and a converter function (A) -> B
func Map[A, B any](a A, f func(A) B) B {
	return f(a)
}
