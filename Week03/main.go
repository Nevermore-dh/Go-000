package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	http *http.Server
	ctx context.Context
}

func NewServer(addr string, mux http.Handler) *Server {
	return &Server{
		http: &http.Server{
			Addr:         addr,
			Handler:      mux,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		ctx: context.Background(),
	}
}

func (s *Server)Start() error {
	log.Println("Server start listening...")
	return s.http.ListenAndServe()
}

func (s *Server)Stop() error {
	log.Println("Server shutdown")
	return s.http.Shutdown(s.ctx)
}

func main() {
	eg, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	server := NewServer(":8080", mux)

	eg.Go(func() error {
		return server.Start()
	})

	eg.Go(func() error {
		sg := make(chan os.Signal, 1)
		signal.Notify(sg, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
		select {
		case <-sg:
			log.Println("Received close signal")
		case <-ctx.Done():
			log.Printf("Server error: %v, context canceled\n", ctx.Err())
		}
		return server.Stop()
	})

	if err := eg.Wait(); err != nil {
		log.Printf("Server error: %v, exit\n", err)
		return
	}
	log.Println("Server exit")
}
