package builder

import (
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/comment/handler"
	"tracer-study-grpc/modules/comment/repository"
	"tracer-study-grpc/modules/comment/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildCommentHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.CommentHandler {
	commentRepo := repository.NewCommentRepository(db)
	commentSvc := service.NewCommentService(cfg, commentRepo)

	return handler.NewCommentHandler(cfg, commentSvc)
}
