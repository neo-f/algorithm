package week12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLC1(t *testing.T) {
}
func TestLC11(t *testing.T) {}
func TestLC15(t *testing.T) {}
func TestLC42(t *testing.T) {}
func TestLC53(t *testing.T) {
	assert.Equal(t, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6)
}
func TestLC84(t *testing.T)  {}
func TestLC167(t *testing.T) {}
func TestLC239(t *testing.T) {}
func TestLC304(t *testing.T) {}
func TestLC560(t *testing.T) {}
func TestLC1109(t *testing.T) {
	assert.Equal(t, corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5), []int{10, 55, 45, 25, 25})
}
func TestLC1248(t *testing.T) {
	assert.Equal(t, numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3), 2)
}