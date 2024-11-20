package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a func 1")

}

func funcao2() {
	fmt.Println("Executando a func 2")
}

func alunoEstaProvado(n1, n2 float32) bool {
	defer fmt.Println("Media calculado.Resultado sera retornado")

	fmt.Println("Entrando na funcao de aprovacao")

	media := n1 + n2/2

	if media < 6 {
		return true

	}
	return false

}
func main() {

	//defer = adiar
	fmt.Println(alunoEstaProvado(4, 10))
}
