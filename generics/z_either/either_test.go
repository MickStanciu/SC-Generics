package z_either

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEither_BindWhenNoErrors(t *testing.T) {
	e := Either[string]{
		Value: "Hello",
	}

	appendX := func(val string) Either[string] {
		return Either[string]{Value: val + "X"}
	}

	appendY := func(val string) Either[string] {
		return Either[string]{Value: val + "Y"}
	}

	appendZ := func(val string) Either[string] {
		return Either[string]{Value: val + "Z"}
	}

	test := e.
		Bind(appendX).
		Bind(appendY).
		Bind(appendZ)

	assert.EqualValues(t, "HelloXYZ", test.Value)
	assert.Nil(t, test.Error)
}

func TestEither_BindWhenErrors(t *testing.T) {
	e := Either[string]{
		Value: "Hello",
	}

	appendX := func(val string) Either[string] {
		return Either[string]{Value: val + "X"}
	}

	appendY := func(val string) Either[string] {
		return Either[string]{Error: fmt.Errorf("something went wrong in Y function")}
	}

	appendZ := func(val string) Either[string] {
		return Either[string]{Value: val + "Z"}
	}

	test := e.
		Bind(appendX).
		Bind(appendY).
		Bind(appendZ)

	assert.EqualValues(t, "", test.Value)
	assert.EqualValues(t, "something went wrong in Y function", test.Error.Error())
}
