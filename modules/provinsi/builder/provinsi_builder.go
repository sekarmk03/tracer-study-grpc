package builder

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/provinsi/handler"
	"tracer-study-grpc/modules/provinsi/repository"
	"tracer-study-grpc/modules/provinsi/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildProvinsiHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.ProvinsiHandler {
	provinsiRepo := repository.NewProvinsiRepository(db)
	provinsiSvc := service.NewProvinsiService(cfg, provinsiRepo)

	return handler.NewProvinsiHandler(cfg, provinsiSvc)
}
