package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/domicmeia/gcp_practice/handler/healthcheck"
	"github.com/domicmeia/gcp_practice/handler/rest"
	"github.com/domicmeia/gcp_practice/translation"
)

func main() {

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if addr == ":" {
		addr = ":8080"
	}

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
