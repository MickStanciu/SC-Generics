package y_maybe

import "fmt"

type Maybe[T any] interface {
	Exists() bool
	Get() (T, error)
	GetOrElse(T) T

	Bind(func(Maybe[T]) Maybe[T]) Maybe[T]
}

type maybe[T any] struct {
	exists    bool
	something *Something[T]
	nothing   *Nothing[T]
}

type Something[T any] struct {
	*maybe[T]
	val T
}

type Nothing[T any] struct {
	*maybe[T]
}

func (m maybe[T]) Bind(f func(Maybe[T]) Maybe[T]) Maybe[T] {
	if m.Exists() {
		return f(m.something)
	}

	return m
}

func (m maybe[T]) Exists() bool {
	return m.exists
}

func (m maybe[T]) Get() (T, error) {
	if m.Exists() {
		return m.something.val, nil
	}
	return *new(T), fmt.Errorf("expected something but got nothing")
}

func (m maybe[T]) GetOrElse(or T) T {
	if m.Exists() {
		return m.something.val
	}
	return or
}

func NewSomething[T any](val T) Maybe[T] {
	s := maybe[T]{
		exists:    true,
		something: &Something[T]{val: val},
	}
	return s
}

func NewNothing[T any](_ T) Maybe[T] {
	n := maybe[T]{
		exists:  false,
		nothing: &Nothing[T]{},
	}
	return n
}
