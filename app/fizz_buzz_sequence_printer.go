package app

import "fmt"

//go:generate mockgen -source fizz_buzz_sequence_printer.go -destination fizz_buzz_sequence_printer_mock.go -package app

type NumberConverter interface {
	Convert(number int) string
}

type OutputWriter interface {
	Write(output string) error
}

type FizzBuzzSequencePrinter struct {
	FizzBuzz NumberConverter
	IOWriter OutputWriter
}

func (p FizzBuzzSequencePrinter) PrintRange(start int, end int) {
	for i := start; i <= end; i++ {
		formatted_text := fmt.Sprintf("%d %s\n", i, p.FizzBuzz.Convert(i))
		p.IOWriter.Write(formatted_text)
	}
}
