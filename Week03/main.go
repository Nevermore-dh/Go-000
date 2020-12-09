package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg := errgroup.Group{}
	serErr := make(chan error, 1)
	sg := make(chan os.Signal, 1)

	s := http.Server{Addr: ":8080"}

	eg.Go(func() error {
		go func() {
			serErr <- s.ListenAndServe()
		}()
		select {
		case err := <-serErr:
			close(sg)
			close(serErr)
			return err
		}
	})

	eg.Go(func() error {
		signal.Notify(sg,
			syscall.SIGINT|syscall.SIGTERM|syscall.SIGKILL)
		<-sg
		return s.Shutdown(context.TODO())
	})

	log.Println(eg.Wait())
}
