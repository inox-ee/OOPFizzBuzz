package spec_test

import (
	"inoxe_e/OOP_fizzbuzz/spec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassThroughRule(t *testing.T) {
	rule := spec.PassThroughRule{}
	assert.Equal(t, rule.Replace(1), "1")
	assert.Equal(t, rule.Replace(2), "2")
	assert.Equal(t, rule.Replace(3), "3")
}
