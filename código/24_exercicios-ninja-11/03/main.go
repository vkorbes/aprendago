package main

import "fmt"

// - Crie um struct "erroEspecial" que implemente a interface builtin.error.
// - Crie uma função que tenha um valor do tipo error como parâmetro.
// - Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.

type erroEspecial struct {
	qualquercoisa string
}

func (e erroEspecial) Error() string {
	return fmt.Sprintf("deu zica! e olha isso: %v", e.qualquercoisa)
}

func erroComoParametro(e error) {
	fmt.Println("Passaram como argumento pra mim, o seguinte: ", e.(erroEspecial).qualquercoisa, "\nno método Error, eu tenho:", e)
}

func main() {
	arg := erroEspecial{"EMOJIS!!!!!!!!"}
	erroComoParametro(arg)
}
