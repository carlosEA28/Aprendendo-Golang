package main

import (
	"fmt"
)

func main() {

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	fmt.Println(i)

	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j < 10; j++ {

	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)

	// }

	nomes := []string{"carlos", "ricardo", "ana"}

	for index, names := range nomes {
		fmt.Println("id:", index, "Nome:", names)

	}

	for index, letra := range "PALAVRA" {
		fmt.Println(index, string(letra))
	}

	usuario := map[string]map[string]string{
		"nome": {
			"primeiro":  "Joao",
			"sobrenome": "vitor",
		},
	}

	for index, valor := range usuario {
		fmt.Println(index, valor)
	}

	for {
	} //loop infinito
}
