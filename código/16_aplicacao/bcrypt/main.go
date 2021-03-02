package main

import "fmt"
import "golang.org/x/crypto/bcrypt"

func main() {
	senha := "20julho1980"
	senhaerrada := "20julho1990"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada)))

}

/*
Quem estiver com problemas no importar do bcrypt, faz isso:
   1 - cd /usr/local/go/src/
   2 - sudo mkdir x
   3 - cd x
   4 - sudo git clone https://github.com/golang/crypto.git
   5 - restart vscode
*/
