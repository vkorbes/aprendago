package escrevendo

import "fmt"

// X é um número inútil.
const X = 10

// Doc é um monte de coisa nenhuma.
func Doc() {
	fmt.Println("Essa função não faz nada.")
}

// doc2 começa com letra minúscula.
func doc2() {
	fmt.Println("Essa função não faz nada.")
}

// Doc3 é a última!
func Doc3() {
	fmt.Println("Essa função não faz nada. X é:", X)
}

// Um comentário de uma linha.

/*
	Pode ser um comentário de mais de uma linha.
*/
