package repository

import (
	"context"
	"errors"
	"log"
	"time"
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
	FindProdiByKode(ctx context.Context, kodeProdi string) (*entity.Prodi, error)
	Create(ctx context.Context, req *entity.Prodi) (*entity.Prodi, error)
	Update(ctx context.Context, prodi *entity.Prodi, updatedFields map[string]interface{}) (*entity.Prodi, error)
	Delete(ctx context.Context, kodeProdi string) error
}

func (p *ProdiRepository) FindAll(ctx context.Context, req any) ([]*entity.Prodi, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProdiRepository - FindAll")
	defer span.End()

	var prodi []*entity.Prodi
	if err := p.db.Debug().WithContext(ctxSpan).Find(&prodi).Error; err != nil {
		log.Println("ERROR: [ProdiRepository - FindAll] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return prodi, nil
}

func (p *ProdiRepository) FindAllFakultas(ctx context.Context, req any) ([]*entity.Fakultas, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProdiRepository - FindAllFakultas")
	defer span.End()

	var fakultas []*entity.Fakultas
	if err := p.db.Debug().WithContext(ctxSpan).Table(entity.ProdiTableName).Select("DISTINCT kode_fak, nama_fak, created_at, updated_at, deleted_at").Find(&fakultas).Error; err != nil {
		log.Println("ERROR: [ProdiRepository - FindAllFakultas] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return fakultas, nil
}

func (p *ProdiRepository) FindProdiByKode(ctx context.Context, kodeProdi string) (*entity.Prodi, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProdiRepository - FindProdiByKode")
	defer span.End()

	var prodi entity.Prodi
	if err := p.db.Debug().WithContext(ctxSpan).Where("kode = ?", kodeProdi).First(&prodi).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [ProdiRepository - FindProdiByKode] Record not found for kode", kodeProdi)
			return nil, status.Errorf(codes.NotFound, "record not found for kode %s", kodeProdi)
		}
		log.Println("ERROR: [ProdiRepository - FindProdiByKode] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &prodi, nil
}

func (p *ProdiRepository) Create(ctx context.Context, req *entity.Prodi) (*entity.Prodi, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProdiRepository - Create")
	defer span.End()

	if err := p.db.Debug().WithContext(ctxSpan).Create(req).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Println("ERROR: [ProdiRepository - Create] Duplicated:", err)
			return nil, status.Errorf(codes.AlreadyExists, "duplicate key error: %v", err)
		}
		log.Println("ERROR: [ProdiRepository - Create] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return req, nil
}

func (p *ProdiRepository) Update(ctx context.Context, prodi *entity.Prodi, updatedFields map[string]interface{}) (*entity.Prodi, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProdiRepository - Update")
	defer span.End()

	updatedFields["updated_at"] = time.Now()
	if err := p.db.Debug().WithContext(ctxSpan).Model(&prodi).Where("kode", prodi.Kode).Updates(updatedFields).Error; err != nil {
		if errors.Is(err, gorm.ErrInvalidValue) {
			log.Println("ERROR: [ProdiRepository - Update] Invalid value:", err)
			return nil, status.Errorf(codes.InvalidArgument, "invalid value: %v", err)
		}
		log.Println("ERROR: [ProdiRepository - Update] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	oldNamaFak := prodi.NamaFak
	newNamaFak := updatedFields["nama_fak"]
	if newNamaFak != nil {
		if newNamaFak != "" {
			if err := p.db.Debug().Model(&entity.Prodi{}).Where("nama_fak = ?", oldNamaFak).Update("nama_fak", newNamaFak).Error; err != nil {
				if errors.Is(err, gorm.ErrInvalidValue) {
					log.Println("ERROR: [ProdiRepository - Update] Invalid value for nama_fak:", err)
					return nil, status.Errorf(codes.InvalidArgument, "invalid value for nama_fak: %v", err)
				}
				log.Println("ERROR: [ProdiRepository - Update] Failed to update nama_fak:", err)
				return nil, status.Errorf(codes.Internal, "failed to update nama_fak: %v", err)
			}
		}
	}

	return prodi, nil
}

func (p *ProdiRepository) Delete(ctx context.Context, kodeProdi string) error {
	ctxSpan, span := trace.StartSpan(ctx, "ProdiRepository - Delete")
	defer span.End()

	if err := p.db.Debug().WithContext(ctxSpan).Where("kode = ?", kodeProdi).Delete(&entity.Prodi{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [ProdiRepository - Delete] Record not found for kode", kodeProdi)
			return status.Errorf(codes.NotFound, "record not found for kode %s", kodeProdi)
		}
		log.Println("ERROR: [ProdiRepository - Delete] Internal server error:", err)
		return status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return nil
}