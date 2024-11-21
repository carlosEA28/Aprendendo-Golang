package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("dentro do metodo salvar", u.nome)
}

func main() {
	usuario1 := usuario{"user1", 39}
	fmt.Println(usuario1)
	usuario1.salvar()
}
