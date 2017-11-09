package main

import "fmt"
import "runtime"

func main() {
	fmt.Println("Esse é programa do exercício de compilação cruzada. Foi compilado num linux/amd64, e agora está rodando num sistema:", runtime.GOARCH, runtime.GOOS)

}
