package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/rating", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[SERVICE RATING] Request headers: %v", r.Header)

		// Логирование полученных заголовков для отладки (опционально)
		for _, header := range []string{"x-request-id", "traceparent", "x-b3-traceid"} {
			fmt.Printf("%s: %s\n", header, r.Header.Get(header))
		}

		log.Printf("[SERVICE RATING] Request headers: %v", r.Header)

		// Ответ микросервиса Ratings
		fmt.Fprintf(w, "Ratings response")
	})

	// Запуск HTTP сервера
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
