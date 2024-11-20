package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var user usuario
	user.nome = "carlos"
	user.idade = 19
	fmt.Println(user)

	user2 := usuario{"joao", 19, endereco{logradouro: "Rua Francico 123", numero: 40}}
	fmt.Println(user2)

	user3 := usuario{nome: "Ricardo"}
	fmt.Println(user3)

}
