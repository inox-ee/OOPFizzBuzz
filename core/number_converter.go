package core

//go:generate mockgen -source number_converter.go -destination number_converter_mock.go -package core

type N2SReplacer interface {
	Replace(number int) string
}

type NumberConverter struct {
	Replacers []N2SReplacer
}

func (nc NumberConverter) Convert(number int) string {
	if len(nc.Replacers) == 0 {
		return "1"
	}
	return nc.Replacers[0].Replace(number)
}
