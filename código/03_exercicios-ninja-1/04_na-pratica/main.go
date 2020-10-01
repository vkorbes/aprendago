package main

import "fmt"

type voceconheceositedosmenes int

var x voceconheceositedosmenes

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
