package main

import "fmt"

func main() {

	retorno := func(txt string) string {
		fmt.Println(txt)
		return fmt.Sprintf("Recebido: %s", txt)
	}("oi")

	fmt.Println(retorno)
}
