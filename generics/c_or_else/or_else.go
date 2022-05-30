package c_or_else

// OrElse returns "or" if "a" is nil
func OrElse[T any](a *T, or *T) *T {
	if a == nil {
		return or
	}
	return a
}

// OrElseMap returns f(T) if T is nil
func OrElseMap[T any](a *T, f func(*T) *T) *T {
	if a == nil {
		return f(a)
	}
	return a
}
