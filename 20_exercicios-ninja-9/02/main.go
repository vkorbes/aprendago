package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println(p.nome, "diz oi!")
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {

	p1 := pessoa{"Genghis Khan", 1000}

	p1.falar() // ← É um shortcut pra (&p1).falar()

	(&p1).falar() // ← É a maneira "mais correta."

	dizerAlgumaCoisa(&p1) // ← Funciona!

	// dizerAlgumaCoisa(p1) // ← Não funciona!
}
