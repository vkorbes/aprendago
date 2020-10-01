package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	threads := 50

	fmt.Println("Início. Goroutines:", runtime.NumGoroutine())

	for i := 0; i < threads; i++ {
		k := i
		go func() {
			fmt.Print("Thread #", k, " inicializando...\n")
			for j := 0; j < 10; j++ {
				fmt.Print("Thread: #", k, " Contagem: ", j, " Operação: ", math.Sqrt(float64(j)*math.Pi), ".\n")
			}
			threads--
			fmt.Print("Thread #", k, " finalizada. Restam ", threads, " threads.\n")
		}()
	}

	fmt.Println("Meio. Goroutines:", runtime.NumGoroutine())
	for threads != 0 {
	}

	fmt.Println("Fim. Goroutines:", runtime.NumGoroutine())
	fmt.Println("Todas as threads encerradas.")

}
