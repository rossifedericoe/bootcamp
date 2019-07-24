package main

import (
	"fmt"

	"github.com/rossifedericoe/bootcamp/clase2/services"
)

func main() {
	_, err := services.Crear(123, "123", "")

	fmt.Println(err)
}
