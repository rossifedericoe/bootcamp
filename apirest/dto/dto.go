package dto

// DTO significa Data Transfer Object: se usa para tener pre-establecido que objeto voy a recibir o enviar
// hacia otras aplicaciones mediante HTTP. (No es exclusivo de las API Rest, en SOAP ya se usaba)
type MovieDTO struct {
	ID       int
	Title    string
	Language string
	Budget   int64
	Revenue  int64
	IMDB     string
}
