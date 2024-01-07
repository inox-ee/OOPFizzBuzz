package spec

import "fmt"

type PassThroughRule struct {
}

func (ptr PassThroughRule) Replace(number int) string {
	return fmt.Sprintf("%d", number)
}
