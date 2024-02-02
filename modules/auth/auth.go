package auth

import (
	"tracer-study-grpc/common/config"
	commonJwt "tracer-study-grpc/common/jwt"
	"tracer-study-grpc/modules/auth/builder"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc"
)

func InitGrpc(server *grpc.Server, cfg config.Config, jwtManager *commonJwt.JWT, grpcConn *grpc.ClientConn) {
	auth := builder.BuildAuthHandler(cfg, jwtManager, grpcConn)
	pb.RegisterAuthServiceServer(server, auth)
}