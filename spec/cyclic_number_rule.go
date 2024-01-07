package spec

type CyclicNumberRule struct {
	Base        int
	Replacement string
}

func (cnr CyclicNumberRule) Replace(number int) string {
	if number%cnr.Base == 0 {
		return cnr.Replacement
	}
	return ""
}
