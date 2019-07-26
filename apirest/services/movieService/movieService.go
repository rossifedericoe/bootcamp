package movieService

import (
	"errors"
	"time"

	"github.com/rossifedericoe/bootcamp/apirest/domain"
	"github.com/rossifedericoe/bootcamp/apirest/repository"
	"github.com/rossifedericoe/bootcamp/apirest/services/contentService"
)

func Crear(title string, lang string, budget int64, revenue int64, imdb string) (*domain.Movie, error) {
	movie := domain.Movie{
		Title:    title,
		Language: lang,
		Budget:   budget,
		Revenue:  revenue,
		IMDB:     imdb,
	}

	isValidMovie, invalidSection := contentService.IsValidContent(movie)
	if !isValidMovie {
		return nil, errors.New("La pelicula no es valida, no tiene " + invalidSection)
	}
	// moviesDatabase = append(moviesDatabase, movie)
	repository.Crear(&movie)
	return &movie, nil
}

func ListarMovies() []domain.Movie {
	allMovies := repository.ListarTodos()
	for i, _ := range allMovies {
		allMovies[i].Country = GetPaisPorIdioma(allMovies[i].Language)
	}
	return allMovies
}

func EliminarMovie(idToDelete int) error {
	movieAEliminar := repository.ObtenerPorId(idToDelete)
	if movieAEliminar == nil || movieAEliminar.ID == 0 {
		return errors.New("No se puede eliminar la pelicula porque no existe")
	}
	repository.Eliminar(movieAEliminar)
	return nil
}

// Funcion que internamente simula la llamada a un servicio externo para que nos diga
// en que pais se puede mostrar la pelicula en base al idioma
func GetPaisPorIdioma(idioma string) string {
	time.Sleep(1 * time.Millisecond)
	switch idioma {
	case "en":
		return "Estados Unidos"
	case "es":
		return "Argentina"
	default:
		return "Pakistan"
	}
}
