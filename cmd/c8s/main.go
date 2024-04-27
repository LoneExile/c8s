package main

import (
	"c8s/config"
	"c8s/internal/app/handlers"
	"c8s/pkg/kube"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	route "c8s/api"
	ks "c8s/internal/service/kube"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	conf := config.GetConfig()

	kubeClient, err := kube.NewClient(conf)
	if err != nil {
		log.Fatalf("Failed to initialize Kubernetes client: %v", err)
	}
	kubeService := ks.NewService(kubeClient)

	router := chi.NewRouter()

	// Set up Huma API
	api := humachi.New(router, huma.DefaultConfig("Kubernetes API", "1.0.0"))
	route.Pod(api, kubeService)
	route.Node(api, kubeService)

	// Set up static file server
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Set up other routes
	router.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		r.NotFound(handlers.NewNotFoundHandler().ServeHTTP)
		// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// 	http.Redirect(w, r, "/app", http.StatusMovedPermanently)
		// })
		r.Get("/", handlers.NewHomeHandler().ServeHTTP)
		r.Get("/pods", handlers.NewPodList(kubeService).ServeHTTP)
		r.Get("/c/pods", handlers.NewPodList(kubeService).Component)
	})

	srv := &http.Server{
		Addr:    conf.Port,
		Handler: router,
	}

	// Graceful shutdown handling
	killSig := make(chan os.Signal, 1)
	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Printf("Environment: %s\n", os.Getenv("env"))
	log.Printf("Server started on %s\n", conf.Port)
	<-killSig

	log.Println("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}
	log.Println("Server shutdown complete")
}
