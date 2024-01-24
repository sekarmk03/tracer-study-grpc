package mhsbiodata

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/mhsbiodata/builder"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	mhsbiodata := builder.BuildMhsBiodataHandler(cfg, grpcConn)
	pb.RegisterMhsBiodataServiceServer(server, mhsbiodata)
}
