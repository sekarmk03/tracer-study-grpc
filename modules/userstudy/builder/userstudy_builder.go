package builder

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/userstudy/handler"
	"tracer-study-grpc/modules/userstudy/repository"
	"tracer-study-grpc/modules/userstudy/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildUserStudyHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.UserStudyHandler {
	userStudyRepo := repository.NewUserStudyRepository(db)
	userStudySvc := service.NewUserStudyService(cfg, userStudyRepo)

	return handler.NewUserStudyHandler(cfg, userStudySvc)
}
