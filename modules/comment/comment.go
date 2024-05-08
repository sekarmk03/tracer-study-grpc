package comment

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/comment/builder"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	comment := builder.BuildCommentHandler(cfg, db, grpcConn)
	pb.RegisterCommentServiceServer(server, comment)
}
