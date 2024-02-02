package builder

import (
	"tracer-study-grpc/common/config"
	commonJwt "tracer-study-grpc/common/jwt"
	"tracer-study-grpc/modules/auth/handler"
	"tracer-study-grpc/modules/mhsbiodata/service"

	"google.golang.org/grpc"
)

func BuildAuthHandler(cfg config.Config, jwtManager *commonJwt.JWT, grpcConn *grpc.ClientConn) *handler.AuthHandler {
	mhsBiodataSvc := service.NewMhsBiodataService(cfg)
	return handler.NewAuthHandler(cfg, mhsBiodataSvc, jwtManager)
}
