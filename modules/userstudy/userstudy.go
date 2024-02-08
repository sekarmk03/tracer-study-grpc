package userstudy

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/userstudy/builder"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	userstudy := builder.BuildUserStudyHandler(cfg, db, grpcConn)
	pb.RegisterUserStudyServiceServer(server, userstudy)
}
