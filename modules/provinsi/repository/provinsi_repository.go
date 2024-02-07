package repository

import (
	"context"
	"log"
	"tracer-study-grpc/modules/provinsi/entity"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		log.Println("ERROR: [ProvinsiRepository-FindAll] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return provinsi, nil
}
