package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Criar novo tipo de Deck
// pega tudo que tem dentro de String
type deck []string

func novoDeck() deck {
	//var card string = "Ace of Spades"

	//card := novaCarta()
	//fmt.Println(card)

	//cartas := []string{"Ace of Diamonds", novaCarta()}
	//cartas := deck{"Ace of Diamonds", novaCarta()}
	//cartas = append(cartas, "Six of Spades")

	cartas := deck{}

	cartasTipos := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cartasValores := []string{"Ace", "Two", "Three", "Four"}

	for _, tipo := range cartasTipos {
		for _, valor := range cartasValores {
			cartas = append(cartas, valor+" of "+tipo)
		}
	}

	return cartas
}

func darMao(d deck, tamanhoMao int) (deck, deck) {
	return d[:tamanhoMao], d[tamanhoMao:]
}

func (d deck) print() {
	//	for _carta := range cartas {
	for i, carta := range d {
		fmt.Println(i, carta)
	}
}

func (d deck) salvarArquivo(nomeArquivo string) error {
	return os.WriteFile(nomeArquivo, d.toBytes(), 0777)
}

func novoDeckDeAquivo(nomeArquivo string) deck {
	byteSlice, err := os.ReadFile(nomeArquivo)

	if err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}

	stringDeck := string(byteSlice)
	sliceStringDeck := strings.Split(stringDeck, ",")

	cartas := deck(sliceStringDeck)

	return cartas
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) toBytes() []byte {
	return []byte(d.toString())
}

func (d deck) embaralhar() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		novaPosicao := r.Intn(len(d) - 1)
		d[i], d[novaPosicao] = d[novaPosicao], d[i]
	}
}
