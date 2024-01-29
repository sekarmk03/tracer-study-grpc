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
	FindAllFakultas(ctx context.Context, req any) ([]*entity.Fakultas, error)
}

func (p *ProdiRepository) FindAll(ctx context.Context, req any) ([]*entity.Prodi, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProdiRepository - FindAll")
	defer span.End()

	var prodi []*entity.Prodi
	if err := p.db.Debug().WithContext(ctxSpan).Find(&prodi).Error; err != nil {
		return nil, err
	}

	return prodi, nil
}

func (p *ProdiRepository) FindAllFakultas(ctx context.Context, req any) ([]*entity.Fakultas, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProdiRepository - FindAllFakultas")
	defer span.End()

	var fakultas []*entity.Fakultas
	if err := p.db.Debug().WithContext(ctxSpan).Table(entity.ProdiTableName).Select("DISTINCT kode_fak, nama_fak, created_at, updated_at, deleted_at").Find(&fakultas).Error; err != nil {
		return nil, err
	}

	return fakultas, nil
}