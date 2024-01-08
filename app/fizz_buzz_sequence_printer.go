package app

import (
	"fmt"
	"inoxe_e/OOP_fizzbuzz/core"
	"inoxe_e/OOP_fizzbuzz/spec"
)

type FizzBuzzSequencePrinter struct{}

func (p FizzBuzzSequencePrinter) PrintRange(start int, end int) {
	fizzbuzz := []core.N2SReplacer{
		spec.CyclicNumberRule{Base: 3, Replacement: "Fizz"},
		spec.CyclicNumberRule{Base: 5, Replacement: "Buzz"},
		spec.PassThroughRule{},
	}
	nc := core.NumberConverter{Replacers: fizzbuzz}
	for i := start; i <= end; i++ {
		fmt.Printf("%d %s\n", i, nc.Convert(i))
	}
}
