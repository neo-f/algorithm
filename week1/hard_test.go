package week1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	assert.Equal(t, calculate("1+2*3"), 7)
	assert.Equal(t, calculate("(1+2)*3"), 9)
}
