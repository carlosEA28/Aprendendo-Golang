package main

import "fmt"

func main() {

	var nome string = "carlos eduardo"
	dia := "segunda"
	fmt.Println(nome)
	fmt.Println(dia)

	var (
		num1      int  = 123
		isLoading bool = true
	)
	fmt.Println(num1, isLoading)

	varialve3, varialve4 := "oi", "tchau"
	fmt.Println(varialve3, varialve4)

	const PI float32 = 3.14

	varialve3, varialve4 = varialve4, varialve3 //inverte os valores

	fmt.Println(varialve3, varialve4)

}
