package main

import "fmt"

func Soma(numbers ...int) int {
	total := 0

	for _, numero := range numbers {
		total += numero
	}
	return total
}

func escrever(txt string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(txt, numero)
	}
}

func main() {
	totalSoma := Soma(10, 20, 30)
	fmt.Println(totalSoma)

	escrever("Ola", 1, 2, 3, 4, 5)

}
