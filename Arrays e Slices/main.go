package main

import "fmt"

func main() {

	var array1 [5]int
	array1[0] = 109
	fmt.Println(array1)

	array2 := [5]string{"Posicao1", "posicao2"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} //os tres pontos significam que o tamanho do array sera determindo pela quantida de itens adicionados na criacao do array
	fmt.Println(array3)

	//Slice

	slice := []int{}
	slice = append(slice, 10)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	//arrays Internos

	fmt.Println("<-------||------->")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 1)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //ver o tamanho
	fmt.Println(cap(slice3)) //ver a capacidade

	//no momento que o go ve que vai estourar a capacidade do slice, o go dobra a capacidade

}
