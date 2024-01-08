package main

import (
	"fmt"
	"inoxe_e/OOP_fizzbuzz/app"
	"inoxe_e/OOP_fizzbuzz/core"
	"inoxe_e/OOP_fizzbuzz/spec"
)

type FizzBuzzFactory struct{}

func (fbf FizzBuzzFactory) Create() app.FizzBuzzSequencePrinter {
	return app.FizzBuzzSequencePrinter{
		FizzBuzz: fbf.createFizzBuzz(),
		IOWriter: fbf.createOutputWriter(),
	}
}

func (fbf FizzBuzzFactory) createFizzBuzz() app.NumberConverter {
	return core.NumberConverter{
		Replacers: []core.N2SReplacer{
			fbf.createFizzRule(),
			fbf.createBuzzRule(),
			fbf.createPassThroughRule(),
		},
	}
}

func (fbf FizzBuzzFactory) createFizzRule() spec.CyclicNumberRule {
	return spec.CyclicNumberRule{
		Base:        3,
		Replacement: "Fizz",
	}
}

func (fbf FizzBuzzFactory) createBuzzRule() spec.CyclicNumberRule {
	return spec.CyclicNumberRule{
		Base:        5,
		Replacement: "Buzz",
	}
}

func (fbf FizzBuzzFactory) createPassThroughRule() spec.PassThroughRule {
	return spec.PassThroughRule{}
}

func (fbf FizzBuzzFactory) createOutputWriter() app.OutputWriter {
	return ConsoleOutput{}
}

type ConsoleOutput struct{}

func (co ConsoleOutput) Write(output string) error {
	fmt.Print(output)
	return nil
}
