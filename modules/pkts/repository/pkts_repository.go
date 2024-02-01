package repository

import (
	"context"
	"time"
	"tracer-study-grpc/modules/pkts/entity"

	"go.opencensus.io/trace"

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
	Update(ctx context.Context, nim string, updatedFields map[string]interface{}) (*entity.PKTS, error)
	FindByAtasan(ctx context.Context, namaA, hpA, emailA string) ([]*string, error)
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

func (p *PKTSRepository) Update(ctx context.Context, nim string, updatedFields map[string]interface{}) (*entity.PKTS, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PKTSRepository - Update")
	defer span.End()

	var pkts entity.PKTS
	if err := p.db.Debug().WithContext(ctxSpan).Where("nim = ?", nim).First(&pkts).Error; err != nil {
		return nil, err
	}

	updatedFields["updated_at"] = time.Now()
	if err := p.db.Debug().WithContext(ctxSpan).Model(&pkts).Updates(updatedFields).Error; err != nil {
		return nil, err
	}

	return &pkts, nil
}

func (p *PKTSRepository) FindByAtasan(ctx context.Context, namaA, hpA, emailA string) ([]*string, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PKTSRepository - FindByAtasan")
	defer span.End()

	var nims []*string
	if err := p.db.Debug().WithContext(ctxSpan).
		Table(entity.PKTSTableName).
		Select("nim").
		Where("hp_atasan LIKE ?", "%"+hpA+"%").
		Or("LOWER(email_atasan) LIKE LOWER(?)", "%"+emailA+"%").
		Or("LOWER(nama_atasan) LIKE LOWER(?)", "%"+namaA+"%").
		Pluck("nim", &nims).
		Error; err != nil {
		return nil, err
	}

	return nims, nil
}
