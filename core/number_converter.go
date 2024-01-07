package core

import "fmt"

type N2SReplacer interface {
	Replace(number int) string
}

type NumberConverter struct {
	Replacers []N2SReplacer
}

func (nc NumberConverter) Convert(number int) string {
	if number%3 == 0 && number%5 == 0 {
		return "FizzBuzz"
	}
	if number%3 == 0 {
		return "Fizz"
	}
	if number%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%d", number)
}
