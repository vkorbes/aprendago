package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var contador int

const quantidadeDeGoroutines = 50000

func main() {

	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de goroutines:\t", quantidadeDeGoroutines, "\nTotal do contador:\t", contador)

}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
}

// - Utilizando goroutines, crie um programa incrementador:
// - Tenha uma variável com o valor da contagem
// - Crie um monte de goroutines, onde cada uma deve:
// 	- Ler o valor do contador
// 	- Salvar este valor
// 	- Fazer yield da thread com runtime.Gosched()
// 	- Incrementar o valor salvo
// 	- Copiar o novo valor para a variável inicial
// - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
// - Demonstre que há uma condição de corrida utilizando a flag -race
