package pkts

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/pkts/builder"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	pkts := builder.BuildPKTSHandler(cfg, db, grpcConn)
	pb.RegisterPKTSServiceServer(server, pkts)
}