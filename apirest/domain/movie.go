package domain

type Movie struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Title    string `json:"title"`
	Language string `json:"language"`
	Budget   int64  `json:"budget"`
	Revenue  int64  `json:"ah"`
	IMDB     string `json:"imdb"`
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
