package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

/**
    Function that sets up a new server
    Here we set up top-level http stuff such as CORS, auth middleware and logging
**/

func NewServer() http.Handler {
    mux := http.NewServeMux()

    var handler http.Handler = mux

    return handler
}


func run (
    ctx context.Context, 
    w io.Writer,
    getenv func(string) string,
) error {
    ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
    defer cancel()

    srv := NewServer()

    httpServer := &http.Server{
        Addr: net.JoinHostPort("localhost", "8080"),
        Handler: srv,
    }

    go func() {
        log.Printf("listening on %s \n", httpServer.Addr)

        if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
        }
    }() 

    var wg sync.WaitGroup
    wg.Add(1)

    go func() {
        defer wg.Done()
        <-ctx.Done()

        shutdownCtx := context.Background()
        shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10 * time.Second)

        defer cancel()

        if err := httpServer.Shutdown(shutdownCtx); err != nil {
            fmt.Fprintf(os.Stderr, "error shutting down http server %s\n", err)
        }
    }()

    wg.Wait()
    return nil
}

func main() {
    ctx := context.Background()
    if err := run(ctx, os.Stdout, os.Getenv); err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        os.Exit(1)
    }
}
