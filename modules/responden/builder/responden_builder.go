package builder

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/responden/handler"
	"tracer-study-grpc/modules/responden/repository"
	"tracer-study-grpc/modules/responden/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildRespondenHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.RespondenHandler {
	respondenRepo := repository.NewRespondenRepository(db)
	respondenSvc := service.NewRespondenService(cfg, respondenRepo)

	return handler.NewRespondenHandler(cfg, respondenSvc)

}