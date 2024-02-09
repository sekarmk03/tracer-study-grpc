package builder

import (
	"tracer-study-grpc/common/config"
	commonJwt "tracer-study-grpc/common/jwt"
	"tracer-study-grpc/modules/auth/handler"
	mhsSvc "tracer-study-grpc/modules/mhsbiodata/service"
	"tracer-study-grpc/modules/pkts/repository"
	pktsSvc "tracer-study-grpc/modules/pkts/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildAuthHandler(cfg config.Config, db *gorm.DB, jwtManager *commonJwt.JWT, grpcConn *grpc.ClientConn) *handler.AuthHandler {
	mhsBiodataSvc := mhsSvc.NewMhsBiodataService(cfg)

	pktsRepository := repository.NewPKTSRepository(db)
	pktsSvc := pktsSvc.NewPKTSService(cfg, pktsRepository)
	return handler.NewAuthHandler(cfg, mhsBiodataSvc, pktsSvc, jwtManager)
}
