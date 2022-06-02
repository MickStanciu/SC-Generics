package a_utils_test

import (
	"github.com/MickStanciu/SC-Generics/generics/a_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	assert.EqualValues(t, 5, a_utils.Max(5, 3))
	assert.EqualValues(t, 5, a_utils.Max(3, 5))
	assert.EqualValues(t, 5, a_utils.Max(5, 5))
}
