package faas

import (
	"net/http"

	"github.com/domicmeia/gcp_practice/handler/rest"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	rest.TranslateHandler(w, r)
}
