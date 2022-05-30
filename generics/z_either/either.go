package z_either

type Either[T any] struct {
	Value T
	Error error
}

func (e Either[T]) Bind(f func(T) Either[T]) Either[T] {
	// if we have an error, we return
	if e.Error != nil {
		return e
	}

	// otherwise we transform
	return f(e.Value)
}
