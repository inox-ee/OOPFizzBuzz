package core_test

import (
	"inoxe_e/OOP_fizzbuzz/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestNumberConverter(t *testing.T) {
// 	nc := core.NumberConverter{}
// 	assert.Equal(t, nc.Convert(1), "1")
// 	assert.Equal(t, nc.Convert(2), "2")
// 	assert.Equal(t, nc.Convert(3), "Fizz")
// 	assert.Equal(t, nc.Convert(4), "4")
// 	assert.Equal(t, nc.Convert(5), "Buzz")
// 	assert.Equal(t, nc.Convert(6), "Fizz")
// 	assert.Equal(t, nc.Convert(10), "Buzz")
// 	assert.Equal(t, nc.Convert(15), "FizzBuzz")
// 	assert.Equal(t, nc.Convert(30), "FizzBuzz")
// }

func TestNumberConverterWithEmptyRule(t *testing.T) {
	nc := core.NumberConverter{Replacers: []core.N2SReplacer{}}
	assert.Equal(t, nc.Convert(1), "1")
}
