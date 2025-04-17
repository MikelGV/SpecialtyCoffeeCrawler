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

	"github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/api/routes"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/utils"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/database"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/config"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/logger"
	"github.com/rs/cors"
)

/**
    Function that sets up a new server
    Here we set up top-level http stuff such as CORS, auth middleware and logging
**/

func NewServer(
    logger *logger.Logger,
    cfg *config.Config,
    db *database.DBStore,
    userStore *database.UserStore,
    roastersStore *database.RoastersStore,
    productStore *database.ProductsStore,
    productTagsStore *database.ProductTagsStore,
    tagsStore *database.TagsStore,
    usrTags *database.User_TagsStore,
    ) http.Handler {
    mux := http.NewServeMux()

    routes.AddRoutes(
        mux, 
        *cfg, 
        logger, 
        userStore,
        roastersStore,
        productStore,
        productTagsStore,
        tagsStore,
        usrTags,
    )

    var handler http.Handler = mux

    corsHandler := cors.New(cors.Options{
        AllowedOrigins: []string{"http://roastnomads.com"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
        ExposedHeaders: []string{"Link"},
        MaxAge: 300,
        AllowCredentials: true,
    }).Handler(handler)

    return corsHandler
}


func run (
    ctx context.Context, 
    w io.Writer,
    getenv func(string) string,
) error {
    ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
    defer cancel()

    logg := logger.NewLogger()

    db, err := database.Connect()

    if err != nil {
        logg.Error("Failed to connect to database:", err.Error())
        return fmt.Errorf("Failed to connect to database: %w:", err)
    }

    if err := utils.InitializeDemoDB(db); err != nil {
        logg.Error("Failed to initialize database with dummy data:", err.Error())
        return fmt.Errorf("failed to initialize database: %w", err)
    }

    srv := NewServer(
        logg,
        &config.Config{},
        db,
        db.Users,
        db.Roasters,
        db.Products,
        db.ProductTags,
        db.Tags,
        db.UserTags,
    )

    httpServer := &http.Server{
        Addr: net.JoinHostPort(config.Env.Host, config.Env.Port),
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
