package faas

import (
	"net/http"

	"github.com/domicmeia/gcp_practice/handler/rest"
	"github.com/domicmeia/gcp_practice/translation"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)
	translateHandler.TranslateHandler(w, r)
}
