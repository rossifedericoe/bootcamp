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
