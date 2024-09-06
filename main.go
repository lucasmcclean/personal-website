package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	err := run(ctx)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	server := NewServer("localhost:8080")

	go func() {
		fmt.Printf("listening on %s\n", server.Addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, time.Second*10)
		defer cancel()
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			fmt.Printf("error shutting down http server: %s\n", err)
		}
	}()

	wg.Wait()
	return nil
}
