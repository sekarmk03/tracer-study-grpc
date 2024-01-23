package builder

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/prodi/handler"
	"tracer-study-grpc/modules/prodi/repository"
	"tracer-study-grpc/modules/prodi/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildProdiHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.ProdiHandler {
	prodiRepo := repository.NewProdiRepository(db)
	prodiSvc := service.NewProdiService(cfg, prodiRepo)

	return handler.NewProdiHandler(cfg, prodiSvc)
}
