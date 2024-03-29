package main

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Rishats/istio-services/pkg/server_info"
	"google.golang.org/grpc"
)

// ServerInfo represents the data structure for server info
type ServerInfo struct {
	Hostname    string `json:"hostname"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
	Uptime      string `json:"uptime"`
}

func fetchServerInfo(r *http.Request) (*ServerInfo, error) {
	// Address of the server_info gRPC server
	serverAddress := "server-info:50051"

	// Establish a connection to the server_info gRPC server
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Create a gRPC client
	client := server_info.NewServerInfoServiceClient(conn)

	istioHeaders := []string{
		"x-request-id", "traceparent", "tracestate",
		"x-ot-span-context", "x-datadog-trace-id",
		"x-datadog-parent-id", "x-datadog-sampling-priority",
		"x-b3-traceid", "x-b3-spanid", "x-b3-parentspanid",
		"x-b3-sampled", "x-b3-flags", "x-cloud-trace-context",
		"grpc-trace-bin", "sw8", "end-user", "user-agent",
		"cookie", "authorization", "jwt",
	}

	// Создание контекста с метаданными
	md := metadata.New(map[string]string{})
	for _, header := range istioHeaders {
		if value := r.Header.Get(header); value != "" {
			// Добавление заголовка и его значения в метаданные
			md.Append(strings.ToLower(header), value)
		}
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Make a request to get server information
	resp, err := client.GetServerInfo(ctx, &server_info.GetServerInfoRequest{})
	if err != nil {
		return nil, err
	}

	return &ServerInfo{
		Hostname:    resp.Hostname,
		Environment: resp.Environment,
		Version:     resp.Version,
		Uptime:      resp.UpTime,
	}, nil
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	info, err := fetchServerInfo(r)
	if err != nil {
		http.Error(w, "[SERVICE Metrics] Failed to fetch server info", http.StatusInternalServerError)
		log.Printf("[SERVICE Metrics] Failed to fetch server info: %v", err)
		return
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(info); err != nil {
		http.Error(w, "[SERVICE Metrics] Failed to encode server info", http.StatusInternalServerError)
		log.Printf("[SERVICE Metrics] Failed to encode server info: %v", err)
	}
}

func main() {
	http.HandleFunc("/metrics", metricsHandler)

	log.Println("[SERVICE Metrics] Starting web server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start web server: %v", err)
	}
}
