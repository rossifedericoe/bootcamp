package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rossifedericoe/bootcamp/apirest/dto"
	"github.com/rossifedericoe/bootcamp/apirest/services/movieService"
)

func main() {
	engine := gin.Default()

	engine.GET("/movies", func(context *gin.Context) {
		context.JSON(200, movieService.ListarMovies())
	})

	engine.GET("/movies/populares", func(context *gin.Context) {
		context.JSON(200, movieService.ListarMoviesPopulares())
	})

	engine.POST("/movies", func(context *gin.Context) {
		var movieDTO dto.MovieDTO
		bindErr := context.BindJSON(&movieDTO)
		if bindErr != nil {
			context.JSON(400, bindErr.Error())
			return
		}

		movieCreada, crearErr := movieService.Crear(movieDTO.Title, movieDTO.Language, movieDTO.Budget, movieDTO.Revenue, movieDTO.IMDB)
		if crearErr != nil {
			context.JSON(400, crearErr.Error())
			return
		}

		context.JSON(200, movieCreada)
	})

	engine.DELETE("/movies/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		eliminarErr := movieService.EliminarMovie(id)
		if eliminarErr != nil {
			context.JSON(404, eliminarErr.Error())
			return
		}
		context.JSON(200, "pudo borrarse")
	})

	// Levanta en localhost:8080
	engine.Run()
}
