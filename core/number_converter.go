package core

import "fmt"

type NumberConverter struct{}

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
