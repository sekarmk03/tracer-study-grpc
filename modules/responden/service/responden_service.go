package service

import (
	"context"
	"log"
	"time"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/responden/entity"
	"tracer-study-grpc/modules/responden/repository"
	"tracer-study-grpc/pb"
)

type RespondenService struct {
	cfg                 config.Config
	respondenRepository repository.RespondenRepositoryUseCase
}

type RespondenServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Responden, error)
	FindByNim(ctx context.Context, nim string) (*entity.Responden, error)
	Update(ctx context.Context, nim string, fields *pb.RespondenUpdate) (*entity.Responden, error)
}

func NewRespondenService(cfg config.Config, respondenRepository repository.RespondenRepositoryUseCase) *RespondenService {
	return &RespondenService{
		cfg:                 cfg,
		respondenRepository: respondenRepository,
	}
}

func (svc *RespondenService) FindAll(ctx context.Context, req any) ([]*entity.Responden, error) {
	res, err := svc.respondenRepository.FindAll(ctx, req)
	if err != nil {
		log.Println("[RespondenService - FindAll] Error while find all responden: ", err)
		return nil, err
	}

	return res, nil
}

func (scv *RespondenService) FindByNim(ctx context.Context, nim string) (*entity.Responden, error) {
	res, err := scv.respondenRepository.FindByNim(ctx, nim)
	if err != nil {
		log.Println("[RespondenService - FindByNim] Error while find responden by nim: ", err)
		return nil, err
	}

	return res, nil
}

func (scv *RespondenService) Update(ctx context.Context, nim string, fields *pb.RespondenUpdate) (*entity.Responden, error) {
	updateMap := map[string]interface{}{
		"ipk":           fields.GetIpk(),
		"kodedikti":     fields.GetKodedikti(),
		"jenjang":       fields.GetJenjang(),
		"namaprodi":     fields.GetNamaprodi(),
		"namaprodi2":    fields.GetNamaprodi2(),
		"kodeprodi":     fields.GetKodeprodi(),
		"kodeprodi2":    fields.GetKodeprodi2(),
		"kodefak":       fields.GetKodefak(),
		"namafak":       fields.GetNamafak(),
		"jlrmasuk":      fields.GetJlrmasuk(),
		"thnmasuk":      fields.GetThnmasuk(),
		"lamastd":       fields.GetLamastd(),
		"updated_at":    time.Now(),
		"status_update": "1",
		"updated_by":    "system",
	}

	res, err := scv.respondenRepository.Update(ctx, nim, updateMap)
	if err != nil {
		log.Println("[RespondenService - Update] Error while update responden: ", err)
		return nil, err
	}

	return res, nil
}
