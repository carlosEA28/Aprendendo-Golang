package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//PONTEIRO E UMA REFERENCIA DE MEMORIA

	var var3 int
	var ponteiro *int

	var3 = 100
	ponteiro = &var3

	fmt.Println(var3, *ponteiro)

	var3 = 150
	fmt.Println(var3, *ponteiro)

}
