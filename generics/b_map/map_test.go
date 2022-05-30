package b_map_test

import (
	"fmt"
	"github.com/MickStanciu/SC-Generics/generics/b_map"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap_IntToString(t *testing.T) {

	converterIntPlus5 := func(value int) string {
		return fmt.Sprintf("%d", value+5)
	}

	resultIntPlus5 := b_map.Map(6, converterIntPlus5)
	assert.EqualValues(t, "11", resultIntPlus5)
}

func TestMap_Struct(t *testing.T) {

	type human struct {
		age  int
		name string
	}

	george := human{
		age:  40,
		name: "George",
	}

	ageFn := func(value human) human {
		return human{
			age:  value.age + 1,
			name: value.name,
		}
	}

	resultAge := b_map.Map(george, ageFn)
	assert.EqualValues(t, 41, resultAge.age)
}
