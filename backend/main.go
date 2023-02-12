package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/spencerfcp/user-login-code-challenge/backend/env"
	"github.com/spencerfcp/user-login-code-challenge/backend/logerr"
	"github.com/spencerfcp/user-login-code-challenge/backend/routes"
)

func main() {
	environment := env.GetEnv()
	db := sqlx.MustOpen("postgres", environment.DatabaseUrl)
	defer func() {
		fmt.Println("shutting down database connection")
		if err := db.Close(); err != nil {
			logerr.FromError(errors.Wrap(err, "failed to shutdown database connection"))
		}
	}()

	mux := http.NewServeMux()

	routes.Handle(
		mux,
		db,
		environment,
	)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
		},
		AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH",
		},
		AllowCredentials: true,
		AllowedHeaders: []string{"accept",
			"authorization",
			"content-type",
			"locale",
			"gpu",
			"User-Agent",
			"X-FORWARDED-FOR",
			"operating_system",
		},
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %v\n", port)
	}

	log.Printf("serving on port %s over http", port)
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           c.Handler(mux),
		ReadHeaderTimeout: 5 * time.Second,
	}

	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-interruptChannel
		fmt.Println("Shutting down HTTP server")
		ctx, timeout := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			fmt.Println("HTTP server shutdown timeout")
			timeout()
		}()

		if err := server.Shutdown(ctx); err != nil {
			logerr.FromError(errors.Wrap(err, "failed to shutdown server"))
		}
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logerr.FromError(errors.Wrapf(err, "failed to bind to port %s", port))
		os.Exit(1)
	}
}
