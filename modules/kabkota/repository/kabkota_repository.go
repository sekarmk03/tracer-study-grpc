package repository

import (
	"context"
	"log"
	"tracer-study-grpc/modules/kabkota/entity"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		log.Println("ERROR [KabKotaRepository - FindAll] Internal error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return kabkota, nil
}
