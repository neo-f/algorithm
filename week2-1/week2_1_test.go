package week21

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test30(t *testing.T) {
	//assert.Equal(t, findSubstring("barfoothefoobarman", []string{"foo", "bar"}), []int{0, 9})
	assert.Equal(t, findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}), []int{8})
}
