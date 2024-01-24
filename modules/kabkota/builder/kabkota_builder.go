package builder

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/kabkota/handler"
	"tracer-study-grpc/modules/kabkota/repository"
	"tracer-study-grpc/modules/kabkota/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildKabKotaHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.KabKotaHandler {
	kabkotaRepo := repository.NewKabKotaRepository(db)
	kabkotaSvc := service.NewKabKotaService(cfg, kabkotaRepo)

	return handler.NewKabKotaHandler(cfg, kabkotaSvc)
}
