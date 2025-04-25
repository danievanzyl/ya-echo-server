package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/danievanzyl/ya-echo-server/internal/server"
)

func main() {
	ctx := context.Background()
	s := server.NewServer(":8080")
	s.Start()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Stop(shutdownCtx)
}
