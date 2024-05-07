package builder

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/post/handler"
	"tracer-study-grpc/modules/post/repository"
	"tracer-study-grpc/modules/post/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildPostHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.PostHandler {
	postRepo := repository.NewPostRepository(db)
	imageSvc := service.NewImageService(cfg)
	postSvc := service.NewPostService(cfg, postRepo)

	return handler.NewPostHandler(cfg, postSvc, imageSvc)
}
