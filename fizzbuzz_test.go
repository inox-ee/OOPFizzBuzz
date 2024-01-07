package main

import (
	"inoxe_e/OOP_fizzbuzz/core"
	"inoxe_e/OOP_fizzbuzz/spec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	nc := core.NumberConverter{Replacers: []core.N2SReplacer{
		spec.CyclicNumberRule{Base: 3, Replacement: "Fizz"},
		spec.CyclicNumberRule{Base: 5, Replacement: "Buzz"},
		spec.PassThroughRule{},
	}}
	assert.Equal(t, nc.Convert(1), "1")
	assert.Equal(t, nc.Convert(2), "2")
	assert.Equal(t, nc.Convert(3), "Fizz")
}
