package generics

// Swap - swaps 2 values around
func Swap[A, B any](a A, b B) (B, A) {
	return b, a
}
