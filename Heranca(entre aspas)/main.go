package main

import "fmt"

type pessoa struct {
	nome      string
	idade     uint8
	sobrenome string
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa{"cadu", 21, "alves"}
	fmt.Println(p1)

	estudante1 := estudante{p1, "ads", "uniritter"}
	fmt.Println(estudante1)
	fmt.Println(estudante1.nome)

}
