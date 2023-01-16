package main

import "fmt"

func main() {
	cartas := novoDeck()
	fmt.Println("========= Deck")
	cartas.print()

	mao, restanteDeck := darMao(cartas, 5)

	fmt.Println("========= Mao")
	mao.print()
	fmt.Println("========= Restante")
	restanteDeck.print()

	cartas.salvarArquivo("DeckOriginal.txt")
	mao.salvarArquivo("Mao.txt")
	restanteDeck.salvarArquivo("Restante.txt")

	fmt.Println("========= Restante")
	cartasArquivo := novoDeckDeAquivo("Mao.txt")
	cartasArquivo.print()

	fmt.Println("========= Embaralhar")

	cartas.embaralhar()
	cartas.print()
}
