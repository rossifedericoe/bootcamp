package main

import (
	"fmt"

	"github.com/rossifedericoe/bootcamp/clase2/services/movieService"
)

func main() {
	movie, err := movieService.Crear(123, "foo", "bar")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Película creada exitosamente: " + fmt.Sprintf("%+v", *movie))
	}

}
