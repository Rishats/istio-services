package main

import (
	"io"
	"log"
	"net/http"
)

func forwardHeaders(src, dest http.Header, headers []string) {
	for _, header := range headers {
		if value := src.Get(header); value != "" {
			dest.Set(header, value)
		}
	}
}

func main() {
	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		// Создание запроса к сервису Reviews
		req, err := http.NewRequest("GET", "http://reviews:8080/reviews", nil)
		if err != nil {
			log.Printf("[SERVICE PRODUCT] Error creating request to Reviews service: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Передача заголовков Istio
		forwardHeaders(r.Header, req.Header, []string{
			"x-request-id", "traceparent", "tracestate",
			"x-ot-span-context", "x-datadog-trace-id",
			"x-datadog-parent-id", "x-datadog-sampling-priority",
			"x-b3-traceid", "x-b3-spanid", "x-b3-parentspanid",
			"x-b3-sampled", "x-b3-flags", "x-cloud-trace-context",
			"grpc-trace-bin", "sw8", "end-user", "user-agent",
			"cookie", "authorization", "jwt",
		})

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("[SERVICE PRODUCT] Error forwarding request to Reviews service: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		log.Printf("[SERVICE PRODUCT] Forwarded request to Reviews service")
		io.Copy(w, resp.Body)
	})

	log.Println("[SERVICE PRODUCT] Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("[SERVICE PRODUCT] Error starting server: %v", err)
	}
}
