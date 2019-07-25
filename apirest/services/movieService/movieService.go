package movieService

import (
	"errors"
	"fmt"

	"github.com/rossifedericoe/bootcamp/apirest/domain"
	"github.com/rossifedericoe/bootcamp/apirest/services/contentService"
)

// NO HACER ESTO EN PRODU
var moviesDatabase []domain.Movie = []domain.Movie{}

func Crear(id int, title string, lang string) (*domain.Movie, error) {
	movie := domain.Movie{
		ID:       id,
		Title:    title,
		Language: lang,
	}

	isValidMovie, invalidSection := contentService.IsValidContent(movie)
	if !isValidMovie {
		return nil, errors.New("La pelicula no es valida, no tiene " + invalidSection)
	}
	moviesDatabase = append(moviesDatabase, movie)
	return &movie, nil
}

func ListarMovies() []domain.Movie {
	return moviesDatabase
}

func EliminarMovie(idToDelete int) error {
	eliminado := false
	var newMoviesDatabase []domain.Movie = []domain.Movie{}
	for _, movie := range moviesDatabase {
		if movie.ID == idToDelete {
			eliminado = true
		} else {
			newMoviesDatabase = append(newMoviesDatabase, movie)
		}
	}
	moviesDatabase = newMoviesDatabase
	if eliminado {
		return nil
	} else {
		return errors.New("No se encontro movie con id " + fmt.Sprint(idToDelete))
	}
}
