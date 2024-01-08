package spec

type CyclicNumberRule struct {
	Base        int
	Replacement string
}

func (cnr CyclicNumberRule) Match(carry string, number int) bool {
	return cnr.Base != 0 && number%cnr.Base == 0
}

func (cnr CyclicNumberRule) Apply(carry string, number int) string {
	return carry + cnr.Replacement
}
