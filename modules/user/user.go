package user

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/user/builder"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	user := builder.BuildUserHandler(cfg, db, grpcConn)
	pb.RegisterUserServiceServer(server, user)
}
