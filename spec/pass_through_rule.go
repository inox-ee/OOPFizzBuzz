package spec

import "fmt"

type PassThroughRule struct {
}

func (ptr PassThroughRule) Match(carry string, number int) bool {
	return carry == ""
}

func (ptr PassThroughRule) Apply(carry string, number int) string {
	return fmt.Sprintf("%d", number)
}
