package movieService

import (
	"errors"

	"github.com/rossifedericoe/bootcamp/apirest/domain"
	"github.com/rossifedericoe/bootcamp/apirest/repository"
	"github.com/rossifedericoe/bootcamp/apirest/services/contentService"
)

// NO HACER ESTO EN PRODU
//var moviesDatabase []domain.Movie = []domain.Movie{}

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
	return repository.ListarTodos()
}

func EliminarMovie(idToDelete int) error {
	movieAEliminar := repository.ObtenerPorId(idToDelete)
	if movieAEliminar == nil {
		return errors.New("No se puede eliminar la pelicula porque no existe")
	}
	repository.Eliminar(movieAEliminar)
	return nil
}
