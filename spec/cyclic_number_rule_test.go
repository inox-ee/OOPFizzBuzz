package spec_test

import (
	"inoxe_e/OOP_fizzbuzz/spec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCyclicNumberRuleMatch(t *testing.T) {
	rule := spec.CyclicNumberRule{Base: 3, Replacement: ""}
	assert.Equal(t, rule.Match("", 1), false)
	assert.Equal(t, rule.Match("", 3), true)
	assert.Equal(t, rule.Match("", 6), true)
}

func TestCyclicNumberRuleApply(t *testing.T) {
	rule := spec.CyclicNumberRule{Base: 0, Replacement: "Buzz"}
	assert.Equal(t, rule.Apply("", 0), "Buzz")
	assert.Equal(t, rule.Apply("Fizz", 0), "FizzBuzz")
}
