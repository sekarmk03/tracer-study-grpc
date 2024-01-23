package repository

import (
	"context"
	"tracer-study-grpc/modules/prodi/entity"

	"go.opencensus.io/trace"
	"gorm.io/gorm"
)

type ProdiRepository struct {
	db *gorm.DB
}

func NewProdiRepository(db *gorm.DB) *ProdiRepository {
	return &ProdiRepository{
		db: db,
	}
}

type ProdiRepositoryUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Prodi, error)
}

func (p *ProdiRepository) FindAll(ctx context.Context, req any) ([]*entity.Prodi, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProdiRepository-FindAll")
	defer span.End()

	var prodi []*entity.Prodi
	if err := p.db.Debug().WithContext(ctxSpan).Find(&prodi).Error; err != nil {
		return nil, err
	}

	return prodi, nil
}
