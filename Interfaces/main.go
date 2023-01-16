package main

import "fmt"

type bot interface {
	ola() string
}
type inglesBot struct{}
type espanholBot struct{}

func main() {
	ib := inglesBot{}
	eb := espanholBot{}

	printOla(ib)
	printOla(eb)
}

func printOla(b bot) {
	fmt.Println(b.ola())
}

func (inglesBot) ola() string {
	return "Hi"
}

func (espanholBot) ola() string {
	return "Hola"
}
