package main

import (
	"fmt"

	"github.com/ellenkorbes/go-aprenda-a-programar/28_exercicios-ninja-13/02/quote"
	"github.com/ellenkorbes/go-aprenda-a-programar/28_exercicios-ninja-13/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
