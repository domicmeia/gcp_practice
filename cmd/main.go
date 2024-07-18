package main

import (
	"log"
	"net/http"
	"time"

	"github.com/domicmeia/gcp_practice/config"
	"github.com/domicmeia/gcp_practice/handler/healthcheck"
	"github.com/domicmeia/gcp_practice/handler/rest"
	"github.com/domicmeia/gcp_practice/translation"
)

func main() {

	cfg := config.LoadConfiguration()
	addr := cfg.Port

	timeout := 10 * time.Second

	server := &http.Server{
		Addr:         addr,
		Handler:      http.NewServeMux(),
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	mux := server.Handler.(*http.ServeMux)

	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)

	mux.HandleFunc("/translate/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/healthcheck", healthcheck.Healthcheck)

	log.Printf("listening on %s\n", addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}
}
