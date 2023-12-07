package handler

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Tiburso/GoManager/common"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	http.Server
	shutdown chan bool
	reqCount uint32
}

func NewServer() *ApiServer {
	port := common.GetEnvWithDefault("PORT", "8080")

	srv := &ApiServer{
		Server: http.Server{
			Addr:         ":" + port,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		shutdown: make(chan bool),
	}

	r := mux.NewRouter().StrictSlash(true)

	// r.HandleFunc("/api/v1/applications", application.CreateApplication).Methods("POST")

	r.Use(mux.CORSMethodMiddleware(r))

	srv.Handler = r

	// Graceful shutdown
	return srv
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/api/v1/health", HealthCheckHandler).Methods("GET")
}

func (s *ApiServer) WaitForShutdown() {
	irq := make(chan os.Signal, 1)
	signal.Notify(irq, os.Interrupt)

	select {
	case <-irq:
		// Received SIGINT (Ctrl + C). Shut down gracefully...
		s.shutdown <- true
	case <-s.shutdown:
		// Received another shutdown request. Already shutting down...
	}

	log.Println("Shutting down...")

	// Create shutdown context with 5 seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// shut down gracefully, but wait no longer than the context deadline
	err := s.Shutdown(ctx)
	if err != nil {
		log.Println("Error:", err)
	}
}

func (s *ApiServer) Start() {
	log.Println("Starting server on port", s.Addr)

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Println("Error:", err)
		}
	}()

	s.WaitForShutdown()
}

func RunServer() {
	srv := NewServer()
	srv.Start()
}
