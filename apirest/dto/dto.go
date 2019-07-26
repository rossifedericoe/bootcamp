package dto

// DTO significa Data Transfer Object: se usa para tener pre-establecido que objeto voy a recibir o enviar
// hacia otras aplicaciones mediante HTTP. (No es exclusivo de las API Rest, en SOAP ya se usaba)

// Con `binding:"required"` decimos que el json que recibimos y que queremos mapear dentro de un MovieDTO tiene
// que tener todos esos campos.
// En nuestro caso si los tiene pero son string vacio, por ejemplo, es algo que validamos con el contentService
type MovieDTO struct {
	ID       int    `binding:"required"`
	Title    string `binding:"required"`
	Language string `binding:"required"`
	Budget   int64  `binding:"required"`
	Revenue  int64  `binding:"required"`
	IMDB     string `binding:"required"`
}
