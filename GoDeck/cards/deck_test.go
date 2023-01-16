package main

import (
	"os"
	"testing"
)

func TestNovoDeck(t *testing.T) {
	d := novoDeck()

	if len(d) != 16 {
		t.Errorf("Tamanho incorreto no Deck, esperava 16 mas recebeu %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("o primeiro elemento não é um Ace of Spades e sim um %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("o último elemento não é um Four of Clubs e sim um %v", d[len(d)-1])
	}
}

func TestSalvarArquivo(t *testing.T) {
	os.Remove("_decktesting")
	d := novoDeck()
	d.embaralhar()
	d.salvarArquivo("_decktesting")

	d2 := novoDeckDeAquivo("_decktesting")

	if len(d2) != 16 {
		t.Errorf("Tamanho incorreto no Deck, esperava 16 mas recebeu %v", len(d2))
	}

	if d2[0] != d[0] {
		t.Errorf("o primeiro elemento não é um %v e sim um %v", d[0], d2[0])
	}

	if d2[len(d)-1] != d[len(d)-1] {
		t.Errorf("o último elemento não é um %v e sim um %v", d[len(d)-1], d2[len(d)-1])
	}
	os.Remove("_decktesting")
}
