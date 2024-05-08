package builder

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/user/handler"
	"tracer-study-grpc/modules/user/repository"
	"tracer-study-grpc/modules/user/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildUserHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.UserHandler {
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(cfg, userRepo)

	return handler.NewUserHandler(cfg, userSvc)
}
