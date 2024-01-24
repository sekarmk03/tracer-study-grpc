package responden

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/responden/builder"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	responden := builder.BuildRespondenHandler(cfg, db, grpcConn)
	pb.RegisterRespondenServiceServer(server, responden)
}