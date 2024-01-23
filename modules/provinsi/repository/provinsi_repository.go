package repository

import (
	"context"
	"tracer-study-grpc/modules/provinsi/entity"

	"go.opencensus.io/trace"
	"gorm.io/gorm"
)

type ProvinsiRepository struct {
	db *gorm.DB
}

func NewProvinsiRepository(db *gorm.DB) *ProvinsiRepository {
	return &ProvinsiRepository{
		db: db,
	}
}

type ProvinsiRepositoryUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Provinsi, error)
}

func (p *ProvinsiRepository) FindAll(ctx context.Context, req any) ([]*entity.Provinsi, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProvinsiRepository - FindAll")
	defer span.End()

	var provinsi []*entity.Provinsi
	if err := p.db.Debug().WithContext(ctxSpan).Find(&provinsi).Error; err != nil {
		return nil, err
	}

	return provinsi, nil
}
