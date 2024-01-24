package repository

import (
	"context"
	"tracer-study-grpc/modules/responden/entity"

	"go.opencensus.io/trace"
	"gorm.io/gorm"
)

type RespondenRepository struct {
	db *gorm.DB
}

func NewRespondenRepository(db *gorm.DB) *RespondenRepository {
	return &RespondenRepository{
		db: db,
	}
}

type RespondenRepositoryUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Responden, error)
	FindByNim(ctx context.Context, nim string) (*entity.Responden, error)
}

func (r *RespondenRepository) FindAll(ctx context.Context, req any) ([]*entity.Responden, error) {
	ctxSpan, span := trace.StartSpan(ctx, "RespondenRepository - FindAll")
	defer span.End()

	var responden []*entity.Responden
	if err := r.db.Debug().WithContext(ctxSpan).Order("created_at desc").Limit(10).Find(&responden).Error; err != nil {
		return nil, err
	}

	return responden, nil
}

func (r *RespondenRepository) FindByNim(ctx context.Context, nim string) (*entity.Responden, error) {
	ctxSpan, span := trace.StartSpan(ctx, "RespondenRepository - FindByNim")
	defer span.End()

	var responden entity.Responden
	if err := r.db.Debug().WithContext(ctxSpan).Where("nim = ?", nim).First(&responden).Error; err != nil {
		return nil, err
	}

	return &responden, nil
}