package fakultas

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/fakultas/builder"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	fakultas := builder.BuildFakultasHandler(cfg, db, grpcConn)
	pb.RegisterFakultasServiceServer(server, fakultas)
}
