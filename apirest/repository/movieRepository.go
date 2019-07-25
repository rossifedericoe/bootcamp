package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rossifedericoe/bootcamp/apirest/domain"
)

var database *gorm.DB

func init() {
	db, dbErr := gorm.Open("mysql", "root:go@tcp(127.0.0.1:3306)/bootcamp_data")
	if dbErr != nil {
		fmt.Println("No se pudo conectar con la base de datos")
		return
	}
	database = db
}

func ListarTodos() []domain.Movie {
	var allMovies []domain.Movie
	database.Find(&allMovies)
	return allMovies
}

func ObtenerPorId(id int) *domain.Movie {
	var movie domain.Movie
	database.First(&movie, id)
	return &movie
}

func Eliminar(movie *domain.Movie) {
	database.Delete(movie)
}

func Crear(movie *domain.Movie) {
	database.Create(movie)
}
