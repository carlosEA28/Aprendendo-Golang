package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "carlos",
		"sobrenome": "Alves",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro":  "Joao",
			"sobrenome": "vitor",
		},
		"curso": {
			"Nome":      "fisioterapia",
			"faculdade": "Uniritter",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Sargitario",
	}

	fmt.Println(usuario2)

}
