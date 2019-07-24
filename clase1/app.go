package main

import (
	"fmt"

	"github.com/rossifedericoe/bootcamp/calculadora"
)

func main() {
	var resultado int
	resultado = calculadora.Sumar(1, 2)
	fmt.Println(resultado)

	resultado = calculadora.SumarConConstante(1, 2)
	fmt.Println(resultado)
}
