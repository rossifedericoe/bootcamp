package movieService

import (
	"errors"

	"github.com/rossifedericoe/bootcamp/clase2/domain"
	"github.com/rossifedericoe/bootcamp/clase2/services/contentService"
)

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
	return &movie, nil
}
