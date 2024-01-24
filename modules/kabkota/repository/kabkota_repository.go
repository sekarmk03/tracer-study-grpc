package repository

import (
	"context"
	"tracer-study-grpc/modules/kabkota/entity"

	"go.opencensus.io/trace"
	"gorm.io/gorm"
)

type KabKotaRepository struct {
	db *gorm.DB
}

func NewKabKotaRepository(db *gorm.DB) *KabKotaRepository {
	return &KabKotaRepository{
		db: db,
	}
}

type KabKotaRepositoryUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.KabKota, error)
}

func (k *KabKotaRepository) FindAll(ctx context.Context, req any) ([]*entity.KabKota, error) {
	ctxSpan, span := trace.StartSpan(ctx, "KabKotaRepository - FindAll")
	defer span.End()

	var kabkota []*entity.KabKota
	if err := k.db.Debug().WithContext(ctxSpan).Find(&kabkota).Error; err != nil {
		return nil, err
	}

	return kabkota, nil
}
