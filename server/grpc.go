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
	maxMsgSize   = 1024 * 1024 * 150
)

type Grpc struct {
	Server   *grpc.Server
	listener net.Listener
	Port     string
}

func NewGrpc(port string /*, options ...grpc.ServerOption*/) *Grpc {
	// options = append(options, grpc.MaxSendMsgSize(maxMsgSize))
	// options = append(options, grpc.MaxRecvMsgSize(maxMsgSize))

	server := grpc.NewServer( /*options...*/ )

	return &Grpc{
		Server: server,
		Port:   port,
	}
}

func NewGrpcServer(port string) *Grpc {
	// var options grpc.ServerOption
	// add option unary interceptor
	server := NewGrpc(port /*, options*/)
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

// func defaultUnaryServerInterceptor() []grpc.UnaryServerInterceptor {
// 	options := []grpc.UnaryServerInterceptor{
// 		grpc_auth.UnaryServerInterceptor(commonJwt.Authorize)
// 	}

// 	return options
// }
