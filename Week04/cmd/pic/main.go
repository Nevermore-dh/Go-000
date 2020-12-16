package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	pb "Week04/api/pic/v1"
	"Week04/internal/pkg/server"
	"Week04/internal/service"

	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

const (
	address = ":8080"
)

func initConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.SetConfigName("server")
	err := viper.MergeInConfig()
	if err != nil {
		log.Fatalf("Init configs failed: %v", err)
	}
}

func main() {
	initConfig()

	srv := server.NewServer(viper.GetString("ApiAddr"))
	pb.RegisterPicServer(srv.GSrv, service.NewPicService())

	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return srv.StartListen(ctx)
	})

	g.Go(func() error {
		sg := make(chan os.Signal)
		signal.Notify(sg, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
		select {
		case sg := <-sg:
			log.Printf("received sinal: %v", sg.String())
			cancel()
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Printf("error: %v", err)
	}
}
