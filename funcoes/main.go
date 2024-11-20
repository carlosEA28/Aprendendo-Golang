package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//multiplos retornos
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main() {
	result := somar(10, 10)
	fmt.Println(result)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("texto")
	fmt.Println(resultado)

	resultadosCalculoSoma, resultadosCalculoSub := calculosMatematicos(10, 5)
	// resultadosCalculoSoma, _ := calculosMatematicos(10, 5) caso nao queira usar duas variaveis
	// _, resultadosCalculoSub := calculosMatematicos(10, 5) caso nao queira usar duas variaveis
	fmt.Println(resultadosCalculoSoma, resultadosCalculoSub)

}
