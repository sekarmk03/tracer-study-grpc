package provinsi

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/provinsi/builder"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	provinsi := builder.BuildProvinsiHandler(cfg, db, grpcConn)
	pb.RegisterProvinsiServiceServer(server, provinsi)
}
