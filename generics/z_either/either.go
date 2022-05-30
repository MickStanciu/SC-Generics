package z_either

type Either[T any] struct {
	Value T
	Error error
}

func (e Either[T]) Bind(f func(T) Either[T]) Either[T] {
	if e.Error != nil {
		return e
	}
	return f(e.Value)
}
