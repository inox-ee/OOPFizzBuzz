package spec_test

import (
	"inoxe_e/OOP_fizzbuzz/spec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCyclicNumberRule(t *testing.T) {
	rule := spec.CyclicNumberRule{Base: 3, Replacement: "Fizz"}
	assert.Equal(t, rule.Replace(1), "")
	assert.Equal(t, rule.Replace(3), "Fizz")
	assert.Equal(t, rule.Replace(6), "Fizz")
}
