package core

//go:generate mockgen -source number_converter.go -destination number_converter_mock.go -package core

type N2SReplacer interface {
	Match(carry string, number int) bool
	Apply(carry string, number int) string
}

type NumberConverter struct {
	Replacers []N2SReplacer
}

func (nc NumberConverter) Convert(number int) string {
	result := ""
	for _, replacer := range nc.Replacers {
		if replacer.Match(result, number) {
			result = replacer.Apply(result, number)
		}
	}
	return result
}
