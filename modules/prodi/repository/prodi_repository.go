package repository

import (
	"context"
	"log"
	"tracer-study-grpc/modules/prodi/entity"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	FindAllFakultas(ctx context.Context, req any) ([]*entity.Fakultas, error)
}

func (p *ProdiRepository) FindAll(ctx context.Context, req any) ([]*entity.Prodi, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProdiRepository - FindAll")
	defer span.End()

	var prodi []*entity.Prodi
	if err := p.db.Debug().WithContext(ctxSpan).Find(&prodi).Error; err != nil {
		log.Println("ERROR: [ProdiRepository-FindAll] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return prodi, nil
}

func (p *ProdiRepository) FindAllFakultas(ctx context.Context, req any) ([]*entity.Fakultas, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProdiRepository - FindAllFakultas")
	defer span.End()

	var fakultas []*entity.Fakultas
	if err := p.db.Debug().WithContext(ctxSpan).Table(entity.ProdiTableName).Select("DISTINCT kode_fak, nama_fak, created_at, updated_at, deleted_at").Find(&fakultas).Error; err != nil {
		log.Println("ERROR: [ProdiRepository-FindAllFakultas] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return fakultas, nil
}
