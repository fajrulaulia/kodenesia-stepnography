package stepnography

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	//question no 1
	result := generateStepnography([]int{1, 1, 0, 0, 0, 0, 1})
	assert.Equal(t, result, "a", "they should be equal")

	//question no 2
	result = generateStepnography([]int{31, 31, 20, 20, 20, 21, 8})
	assert.Equal(t, result, "b", "they should be equal")

	//question no 3
	result = generateStepnography([]int{51, 61, 10, 0, 0, 15, 17})
	assert.Equal(t, result, "c", "they should be equal")
}
