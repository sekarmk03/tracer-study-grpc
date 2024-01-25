package builder

import (
	"tracer-study-grpc/common/config"
	mhsbSvc "tracer-study-grpc/modules/mhsbiodata/service"
	"tracer-study-grpc/modules/responden/handler"
	"tracer-study-grpc/modules/responden/repository"
	respSvc "tracer-study-grpc/modules/responden/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildRespondenHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.RespondenHandler {
	respondenRepo := repository.NewRespondenRepository(db)
	respondenSvc := respSvc.NewRespondenService(cfg, respondenRepo)

	mhsbiodataSvc := mhsbSvc.NewMhsBiodataService(cfg)

	return handler.NewRespondenHandler(cfg, respondenSvc, mhsbiodataSvc)
}
