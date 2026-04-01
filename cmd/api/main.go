package main

import (
	"context"
	"goauth/internal/app"
	"goauth/internal/httpserver"
	"log"
	"net/http"
	"time"
)

func main() {
	// Start Up Code
	ctx := context.Background()

	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("StartUp Failed %v", err)
	}

	defer func() {
		if err := a.Close(ctx); err != nil {
			log.Fatalf("ShutDown Warning: %v", err)
		}
	}()

	router := httpserver.NewRouter()

	// Standard http server in golang
	srv := &http.Server{
		Addr: ":5000",
		Handler: router,
		ReadHeaderTimeout: 5 * time.Second,
	}


	log.Printf("Api Running on %s", srv.Addr)


	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("Server Closed")
			return
		}
		log.Fatalf("Server Error: %v", err)
	}


}