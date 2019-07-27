package movieService

import (
	"errors"
	"fmt"
	"sync"
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

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(allMovies))

	for i, _ := range allMovies {
		go setPais(&allMovies[i], &waitGroup)
	}

	waitGroup.Wait()

	return allMovies
}

func ListarMoviesPopulares() []domain.Movie {
	allMovies := repository.ListarTodos()

	popularesCh := make(chan domain.Movie, len(allMovies))
	normalesCh := make(chan domain.Movie, len(allMovies))

	for _, movie := range allMovies {
		go calificar(movie, popularesCh, normalesCh)
	}

	var moviePopoulares []domain.Movie
	var moviesNormales []domain.Movie

	for index := 0; index < len(allMovies); index++ {
		select {
		case pop := <-popularesCh:
			moviePopoulares = append(moviePopoulares, pop)
		case normal := <-normalesCh:
			moviesNormales = append(moviesNormales, normal)
		}
	}

	fmt.Println("La cantidad de movies consideradas no-populares (normales) son: " + fmt.Sprint(len(moviesNormales)))

	return moviePopoulares
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
// en que pais se puede mostrar la pelicula en base al idioma.
// Tiene un sleep de 1 milisegundo para que parezca un poco mas real
func setPais(movie *domain.Movie, waitGroup *sync.WaitGroup) {
	time.Sleep(1 * time.Millisecond)
	switch movie.Language {
	case "en":
		movie.Country = "Estados Unidos"
	case "es":
		movie.Country = "Argentina"
	default:
		movie.Country = "Pakistan"
	}
	waitGroup.Done()
}

// Funcion que internamente simula la llamada a un servicio externo para que nos diga
// si una movie puede ser considerada popular o no.
// Tiene un sleep de 1 milisegundo para que parezca un poco mas real
func calificar(movie domain.Movie, popularesCh chan domain.Movie, normalesCh chan domain.Movie) {
	time.Sleep(500 * time.Millisecond)
	if movie.Revenue > (movie.Budget * 10) {
		popularesCh <- movie
	} else {
		normalesCh <- movie
	}
}
