package domain

type Movie struct {
	ID       int
	Title    string
	Language string
}

func (movie Movie) ValidTitle() bool {
	return movie.Title != ""
}

func (movie Movie) ValidLanguage() bool {
	return movie.Language != ""
}

type IContent interface {
	ValidTitle() bool
	ValidLanguage() bool
}
