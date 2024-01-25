package builder

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/mhsbiodata/handler"
	"tracer-study-grpc/modules/mhsbiodata/service"

	"google.golang.org/grpc"
)

func BuildMhsBiodataHandler(cfg config.Config, grpcConn *grpc.ClientConn) *handler.MhsBiodataHandler {
	mhsbiodataSvc := service.NewMhsBiodataService(cfg)
	return handler.NewMhsBiodataHandler(cfg, mhsbiodataSvc)
}
