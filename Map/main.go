package main

import "fmt"

func main() {
	//var cores map[string]string

	//cores := make(map[string]string)

	cores := map[string]string{
		"vermelho": "#FF0000",
		"verde":    "#GG0000",
		"azul":     "#AAAAA",
	}

	//cores["Azul"] = "#FFFFFF"
	//delete(cores, "Azul")

	fmt.Println(cores)
	printMap(cores)

}

func printMap(c map[string]string) {
	for chave, valor := range c {
		fmt.Printf("%v - %v", chave, valor)
		fmt.Println()
	}
}
