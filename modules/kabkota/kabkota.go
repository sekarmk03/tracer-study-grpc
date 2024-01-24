package kabkota

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/kabkota/builder"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	kabkota := builder.BuildKabKotaHandler(cfg, db, grpcConn)
	pb.RegisterKabKotaServiceServer(server, kabkota)
}
