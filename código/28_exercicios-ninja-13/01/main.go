package starting_code

import (
	"fmt"

	"github.com/ellenkorbes/go-aprenda-a-programar/28_exercicios-ninja-13/01/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
