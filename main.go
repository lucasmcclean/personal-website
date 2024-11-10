package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/lucasmcclean/personal-website/server"
)

const port = ":3000"

func main() {
	ctx := context.Background()

	err := run(ctx)
	if err != nil {
		log.Fatalf("server closed with error: %s\n", err)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	srv := server.New(port)

	go func() {
		log.Printf("listening and serving on %s\n", srv.Addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("error listening and serving: %s\n", err)
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
		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			log.Printf("error shutting down http srv: %s\n", err)
		}
	}()

	wg.Wait()
	return nil
}
