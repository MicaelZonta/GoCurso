package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	contatoinfo
}

type contatoinfo struct {
	email    string
	telefone string
}

func main() {
	//var p pessoa
	//p := pessoa{"Nome", "Sobrenome"}
	pes := pessoa{
		nome:      "Nome",
		sobrenome: "Sobrenome",
		contatoinfo: contatoinfo{
			email:    "email",
			telefone: "telefone",
		},
	}

	//ponteiro := &pes
	pes.setNome("teste")
	println()
	pes.print()
}

func (p *pessoa) setNome(n string) {
	(*p).nome = n
}

func (p pessoa) print() {
	fmt.Printf("%+v", p)
}
