package server

import (
	"context"
	"fmt"
	"net/http"
)

type HTTPServer interface {
	Start()
	Stop(context.Context)
}

func NewServer(addr string) HTTPServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", echoHandler)

	return &httpServer{
		http: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}

type httpServer struct {
	http *http.Server
}

func (h *httpServer) Start() {
	go func() {
		fmt.Println("Starting server on", h.http.Addr)
		if err := h.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server error:", err)
		}
	}()
}

func (h *httpServer) Stop(ctx context.Context) {
	if err := h.http.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	}
	fmt.Println("Server stopped")
}
