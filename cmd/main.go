package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/danievanzyl/ya-echo-server/internal/server"
)

func main() {
	ctx := context.Background()
	s := server.NewServer(":8080")
	s.Start()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	s.Stop(ctx)
	fmt.Println("Received shutdown signal:", ctx.Err())
}
