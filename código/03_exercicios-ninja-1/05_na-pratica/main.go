package main

import "fmt"

type voceconheceositedosmenes int

var x voceconheceositedosmenes

var y int

func main() {
	x = 42
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
