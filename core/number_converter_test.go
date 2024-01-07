package core_test

import (
	"inoxe_e/OOP_fizzbuzz/core"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
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
	assert.Equal(t, nc.Convert(1), "")
}

func TestNumberConverterWithSingleRule(t *testing.T) {
	ctrl := gomock.NewController(t)
	n2sr_mock := core.NewMockN2SReplacer(ctrl)
	n2sr_mock.EXPECT().Replace(1).Return("Replaced")

	nc := core.NumberConverter{Replacers: []core.N2SReplacer{n2sr_mock}}
	assert.Equal(t, nc.Convert(1), "Replaced")
}

func TestNumberConverterWithFizzBuzzRule(t *testing.T) {
	ctrl := gomock.NewController(t)
	nc := core.NumberConverter{Replacers: []core.N2SReplacer{createMockRule(ctrl, 1, "Fizz"), createMockRule(ctrl, 1, "Buzz")}}

	assert.Equal(t, nc.Convert(1), "FizzBuzz")
}

func createMockRule(ctrl *gomock.Controller, number int, result string) core.N2SReplacer {
	n2sr_mock := core.NewMockN2SReplacer(ctrl)
	n2sr_mock.EXPECT().Replace(number).Return(result)
	return n2sr_mock
}

func TestNumberConverterWithUnmachedFizzBuzzRulesAndConstantRule(t *testing.T) {
	ctrl := gomock.NewController(t)
	nc := core.NumberConverter{Replacers: []core.N2SReplacer{createMockRule(ctrl, 1, ""), createMockRule(ctrl, 1, ""), createMockRule(ctrl, 1, "1")}}

	assert.Equal(t, nc.Convert(1), "1")
}
