package main

import (
	"fmt"
	"time"
)

func fibonnaci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonnaci(posicao-2) + fibonnaci(posicao-1)
}

func main() {

	posicao := uint(15)

	for i := uint(0); i <= posicao; i++ {

		time.Sleep(time.Second)
		fmt.Println(fibonnaci(i))
	}
}
