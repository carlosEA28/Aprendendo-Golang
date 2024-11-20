package auxilar

import (
	"fmt"
)

//pacotes são um conjunto de códigos
//funções em letras Maiusculas, são consideradas públicas(podem ser Importadas por outros pacotes), ja as funcoes de letras minusculas, são consideradas privadas(não podem ser acessadas em outros pacotes, apenas no pacote de criação)

func Escrever() {
	fmt.Print("Escrevendo do Auxiar")
	escrever2()
}
