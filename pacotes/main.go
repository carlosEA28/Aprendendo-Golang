package main

import (
	"fmt"
	auxilar "modulo/auxiliar"

	"github.com/badoux/checkmail" //pacote externo
)

func main() {
	fmt.Print("Escrevendo do main")
	auxilar.Escrever()
	// auxilar.escrever2() não consegue acessar

	erro := checkmail.ValidateFormat("cadu")
	fmt.Println(erro)
}
