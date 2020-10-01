package main

import "fmt"

func Add(x, y int) int {
	s := []int{x, y}
	return s[0] + s[1]
}

func main() {
	fmt.Println(Add(1, 2))
}
