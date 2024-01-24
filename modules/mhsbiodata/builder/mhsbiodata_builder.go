package builder

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/mhsbiodata/handler"

	"google.golang.org/grpc"
)

func BuildMhsBiodataHandler(cfg config.Config, grpcConn *grpc.ClientConn) *handler.MhsBiodataHandler {
	return handler.NewMhsBiodataHandler(cfg)
}
