package auth

import (
	"tracer-study-grpc/common/config"
	commonJwt "tracer-study-grpc/common/jwt"
	"tracer-study-grpc/modules/auth/builder"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, jwtManager *commonJwt.JWT, grpcConn *grpc.ClientConn) {
	auth := builder.BuildAuthHandler(cfg, db, jwtManager, grpcConn)
	pb.RegisterAuthServiceServer(server, auth)
}
