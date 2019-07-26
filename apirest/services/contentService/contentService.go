// Servicio basico que se encarga de recibir cualquier cosa que immplemente la interfaz "content", osea,
// que es un contenido; y valida que tenga minimamente algo como titulo y como idioma.

// Si tuvieramos la entidad de dominio serie o documental (ademas de pelicula), podria implementar la misma
// interfaz y usar este servicio para validar su titulo e idioma.
package contentService

import (
	"github.com/rossifedericoe/bootcamp/apirest/domain"
)

func IsValidContent(content domain.IContent) (bool, string) {
	if !content.ValidTitle() {
		return false, "title"
	}

	if !content.ValidLanguage() {
		return false, "lang"
	}

	return true, ""
}
