package prodi

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/prodi/builder"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	prodi := builder.BuildProdiHandler(cfg, db, grpcConn)
	pb.RegisterProdiServiceServer(server, prodi)
}
