package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

const (
	connProtocol = "tcp"
)

type Grpc struct {
	Server   *grpc.Server
	listener net.Listener
	Port     string
}

func NewGrpc(port string) *Grpc {
	server := grpc.NewServer()
	return &Grpc{
		Server: server,
		Port:   port,
	}
}

func NewGrpcServer(port string) *Grpc {
	server := NewGrpc(port)
	return server
}

func (g *Grpc) Run() error {
	var err error
	g.listener, err = net.Listen(connProtocol, fmt.Sprintf(":%s", g.Port))
	if err != nil {
		return err
	}

	go g.serve()
	log.Printf("grpc server is running on port %s\n", g.Port)
	return nil
}

func (g *Grpc) serve() {
	if err := g.Server.Serve(g.listener); err != nil {
		panic(err)
	}
}

func (g *Grpc) AwaitTermination() error {
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	<-sign

	g.Server.GracefulStop()
	return g.listener.Close()
}