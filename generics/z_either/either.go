package z_either

type Either[L, R any] interface {
	IsLeft() bool
	GetRightOrPanic() R
	GetLeftOrPanic() L
	Map(func(R) Either[L, R]) Either[L, R]
	Bind(func(Either[L, R]) Either[L, R]) Either[L, R]
}

type either[L, R any] struct {
	left  *left[L, R]
	right *right[L, R]

	isLeft bool
}

type left[L, R any] struct {
	*either[L, R]
	cargo L
}

type right[L, R any] struct {
	*either[L, R]
	cargo R
}

func NewRight[L, R any](r R) Either[L, R] {
	valR := right[L, R]{
		cargo: r,
	}

	valR.either = &either[L, R]{
		isLeft: false,
		right:  &valR,
	}

	return valR
}

func NewLeft[L, R any](l L) Either[L, R] {
	valL := left[L, R]{
		cargo: l,
	}

	valL.either = &either[L, R]{
		isLeft: true,
		left:   &valL,
	}

	return valL
}

func (e either[L, R]) Map(f func(R) Either[L, R]) Either[L, R] {
	// if we have an error, we return
	if e.isLeft {
		return e
	}

	// otherwise we map
	return f(e.right.cargo)
}

func (e either[L, R]) Bind(f func(Either[L, R]) Either[L, R]) Either[L, R] {
	// if we have an error, we return
	if e.isLeft {
		return e
	}

	return f(e.right)
}

func (e either[L, R]) GetRightOrPanic() R {
	if e.isLeft {
		panic("supposed to be a right-value here!!!")
	}
	return e.right.cargo
}

func (e either[L, R]) GetLeftOrPanic() L {
	if e.isLeft {
		return e.left.cargo
	}
	panic("supposed to be a left-value here!!!")
}

func (e either[L, R]) IsLeft() bool {
	return e.isLeft
}
