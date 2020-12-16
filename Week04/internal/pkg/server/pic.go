package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	GSrv *grpc.Server
	addr string
}

func NewServer(addr string) *Server {
	return &Server{GSrv: grpc.NewServer(), addr: addr}
}

func (s *Server) StartListen(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	log.Printf("Pic server start listening: %s", s.addr)

	go func() {
		<-ctx.Done()
		s.GSrv.GracefulStop()
		log.Printf("Pic server gracefully stoped.")
	}()

	return s.GSrv.Serve(listener)
}