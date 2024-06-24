package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/domicmeia/gcp_practice/handler/rest"
)

func main() {

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if addr == ":" {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/translate/hello", rest.TranslateHandler)

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}
