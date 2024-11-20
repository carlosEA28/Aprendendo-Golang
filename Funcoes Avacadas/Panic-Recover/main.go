package main

import "fmt"

func alunoEstaProvado(n1, n2 float32) bool {
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MEDIA E EXATAMENTE 6!")
}
func main() {
	fmt.Println(alunoEstaProvado(6, 6))
	fmt.Println("pos execucao")
}
