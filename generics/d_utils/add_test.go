package d_utils_test

import (
	"github.com/MickStanciu/SC-Generics/generics/d_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddThings_WhenString(t *testing.T) {
	result := d_utils.AddThings("5", "6")
	assert.EqualValues(t, "56", result)
}

func TestAddThings_WhenInt(t *testing.T) {
	result := d_utils.AddThings(5, 6)
	assert.EqualValues(t, 11, result)
}
