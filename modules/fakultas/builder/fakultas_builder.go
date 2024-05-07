package builder

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/fakultas/handler"
	"tracer-study-grpc/modules/fakultas/repository"
	"tracer-study-grpc/modules/fakultas/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildFakultasHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.FakultasHandler {
	fakultasRepo := repository.NewFakultasRepository(db)
	fakultasSvc := service.NewFakultasService(cfg, fakultasRepo)

	return handler.NewFakultasHandler(cfg, fakultasSvc)
}
