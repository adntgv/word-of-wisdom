package server

import (
	"applicationDesignTest/config"
	"applicationDesignTest/internal/http/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

type Api struct {
	server *http.Server
}

func NewApi() (*Api, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("could not read config: %v", err)
	}

	router := chi.NewRouter()

	routes.NewOrdersRoute(router).Register()

	return &Api{
		server: &http.Server{
			Addr:           fmt.Sprintf(":%v", cfg.Port),
			Handler:        router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}, nil
}

func (api *Api) Run() error {
	var wg = sync.WaitGroup{}
	wg.Add(1)

	// Running server in Goroutines
	go func() {
		defer wg.Done()

		log.Printf("Starting api server, listening at %s\n", api.server.Addr)

		if err := api.server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down the api server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := api.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shut down api server: %s", err.Error())
	}

	wg.Wait()
	log.Println("api server successfully shutdown")

	return nil
}
