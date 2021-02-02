package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mhutter/kroki"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	s := kroki.NewServer()

	srv := http.Server{
		Addr:         ":" + getEnvOr("PORT", "8443"),
		Handler:      s,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("Listening on https://localhost" + srv.Addr)
	return srv.ListenAndServeTLS(
		getEnvOr("TLS_CERT", "cert.pem"),
		getEnvOr("TLS_KEY", "key.pem"),
	)
}

func getEnvOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
