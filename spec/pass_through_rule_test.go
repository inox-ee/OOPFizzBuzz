package spec_test

import (
	"inoxe_e/OOP_fizzbuzz/spec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassThroughRuleApply(t *testing.T) {
	rule := spec.PassThroughRule{}
	assert.Equal(t, rule.Apply("", 1), "1")
	assert.Equal(t, rule.Apply("", 2), "2")
	assert.Equal(t, rule.Apply("Fizz", 3), "3")
}

func TestPassThroughRuleMatch(t *testing.T) {
	rule := spec.PassThroughRule{}
	assert.Equal(t, rule.Match("", 1), true)
	assert.Equal(t, rule.Match("Fizz", 1), false)
}
