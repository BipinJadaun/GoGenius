package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	assert.Equal(t, true, PalindromeString("abcba"))
	assert.Equal(t, true, PalindromeNumber(12321))
}
