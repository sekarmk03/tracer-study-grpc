package builder

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/pkts/handler"
	"tracer-study-grpc/modules/pkts/repository"
	"tracer-study-grpc/modules/pkts/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildPKTSHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.PKTSHandler {
	pktsRepo := repository.NewPKTSRepository(db)
	pktsSvc := service.NewPKTSService(cfg, pktsRepo)

	return handler.NewPKTSHandler(cfg, pktsSvc)
}