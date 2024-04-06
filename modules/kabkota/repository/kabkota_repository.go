package repository

import (
	"context"
	"errors"
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
	FindByIdWil(ctx context.Context, idWil string) (*entity.KabKota, error)
}

func (k *KabKotaRepository) FindAll(ctx context.Context, req any) ([]*entity.KabKota, error) {
	ctxSpan, span := trace.StartSpan(ctx, "KabKotaRepository - FindAll")
	defer span.End()

	var kabkota []*entity.KabKota
	if err := k.db.Debug().WithContext(ctxSpan).Find(&kabkota).Error; err != nil {
		log.Println("ERROR: [KabKotaRepository-FindAll] Internal error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return kabkota, nil
}

func (k *KabKotaRepository) FindByIdWil(ctx context.Context, idWil string) (*entity.KabKota, error) {
	ctxSpan, span := trace.StartSpan(ctx, "KabKotaRepository - FindByIdWil")
	defer span.End()

	var kabkota entity.KabKota
	if err := k.db.Debug().WithContext(ctxSpan).Where("id_wil = ?", idWil).First(&kabkota).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [KabKotaRepository-FindByIdWil] Record not found for idWil", idWil)
			return nil, status.Errorf(codes.NotFound, "record not found for idWil %s", idWil)
		}
		log.Println("ERROR: [KabKotaRepository-FindByIdWil] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &kabkota, nil
}
