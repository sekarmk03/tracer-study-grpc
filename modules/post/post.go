package post

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/post/builder"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	post := builder.BuildPostHandler(cfg, db, grpcConn)
	pb.RegisterPostServiceServer(server, post)
}
