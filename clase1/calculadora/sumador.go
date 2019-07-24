package calculadora

import "fmt"

const constante int = 500

func Sumar(num1 int, num2 int) int {
	resultado := num1 + num2
	return resultado
}

func SumarConConstante(num1 int, num2 int) int {
	resultado := num1 + num2 + constante
	return resultado
}

func init() {
	fmt.Println("inicializando calculadora!")
}
