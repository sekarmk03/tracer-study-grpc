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
	FindByNim(ctx context.Context, nim string) (*entity.PKTS, error)
	Create(ctx context.Context, req *entity.PKTS) (*entity.PKTS, error)
}

func (p *PKTSRepository) FindAll(ctx context.Context, req any) ([]*entity.PKTS, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PKTSRepository - FindAll")
	defer span.End()
	
	var pkts []*entity.PKTS
	if err := p.db.Debug().WithContext(ctxSpan).Order("created_at desc").Limit(10).Find(&pkts).Error; err != nil {
		return nil, err
	}

	return pkts, nil
}

func (p *PKTSRepository) FindByNim(ctx context.Context, nim string) (*entity.PKTS, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PKTSRepository - FindByNim")
	defer span.End()

	var pkt entity.PKTS
	if err := p.db.Debug().WithContext(ctxSpan).Where("nim = ?", nim).First(&pkt).Error; err != nil {
		return nil, err
	}

	return &pkt, nil
}

func (p *PKTSRepository) Create(ctx context.Context, req *entity.PKTS) (*entity.PKTS, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PKTSRepository - Create")
	defer span.End()

	if err := p.db.Debug().WithContext(ctxSpan).Create(&req).Error; err != nil {
		return nil, err
	}

	return req, nil
}