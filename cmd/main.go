package main

import (
	"log"
	"net/http"
	"time"

	"github.com/domicmeia/gcp_practice/config"
	"github.com/domicmeia/gcp_practice/handler/healthcheck"
	"github.com/domicmeia/gcp_practice/handler/info"
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

	mux := API(cfg)

	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}
}

func API(cfg config.Configuration) *http.ServeMux {
	mux := http.NewServeMux()

	var translationService rest.Translator

	translationService = translation.NewStaticService()

	if cfg.LegacyEndpoit != "" {
		log.Printf("creating external translation client: %s", cfg.LegacyEndpoit)
		client := translation.NewHelloClient(cfg.LegacyEndpoit)
		translationService = translation.NewRemoteService(client)
	}

	if cfg.DatabaseURL != "" {
		db := translation.NewDatabaseServic(cfg)
		translationService = db
	}

	translateHandler := rest.NewTranslateHandler(translationService)

	mux.HandleFunc("/translate/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/healthcheck", healthcheck.Healthcheck)
	mux.HandleFunc("/info", info.Info)

	return mux
}
