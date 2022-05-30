package c_or_else_test

import (
	"github.com/MickStanciu/SC-Generics/generics/c_or_else"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrElse_WhenNotNil(t *testing.T) {
	result := c_or_else.OrElse(intToP(44), intToP(22))
	assert.EqualValues(t, 44, *result)
}

func TestOrElse_WhenNil(t *testing.T) {
	result := c_or_else.OrElse(nil, intToP(22))
	assert.EqualValues(t, 22, *result)
}

type cat struct {
	breed string
}

func TestOrElseMap_WhenNotNil(t *testing.T) {

	tabby := cat{breed: "tabby"}

	result := c_or_else.OrElseMap(&tabby, func(value *cat) *cat {
		return &cat{breed: "unknown cat type"}
	})

	assert.EqualValues(t, "tabby", result.breed)
}

func TestOrElseMap_WhenNil(t *testing.T) {

	result := c_or_else.OrElseMap(nil, func(value *cat) *cat {
		return &cat{breed: "unknown cat type"}
	})

	assert.EqualValues(t, "unknown cat type", result.breed)
}

func intToP(val int) *int {
	return &val
}
