package repository

import (
	"context"
	"go.opencensus.io/trace"
	"tracer-study-grpc/modules/pkts/entity"

	"gorm.io/gorm"
)

type PKTSRepository struct {
	db *gorm.DB
}

func NewPKTSRepository(db *gorm.DB) *PKTSRepository {
	return &PKTSRepository{
		db: db,
	}
}

type PKTSRepositoryUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.PKTS, error)
}

func (r *PKTSRepository) FindAll(ctx context.Context, req any) ([]*entity.PKTS, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PKTSRepository - FindAll")
	defer span.End()
	
	var pkts []*entity.PKTS
	if err := r.db.Debug().WithContext(ctxSpan).Order("created_at desc").Limit(10).Find(&pkts).Error; err != nil {
		return nil, err
	}

	return pkts, nil
}