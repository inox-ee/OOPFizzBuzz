package main

func main() {
	factory := FizzBuzzFactory{}
	p := factory.Create()
	p.PrintRange(1, 100)
}
