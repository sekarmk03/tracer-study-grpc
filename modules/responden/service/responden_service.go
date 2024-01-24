package service

import (
	"context"
	"log"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/responden/entity"
	"tracer-study-grpc/modules/responden/repository"
)

type RespondenService struct {
	cfg                 config.Config
	respondenRepository repository.RespondenRepositoryUseCase
}

type RespondenServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Responden, error)
	FindByNim(ctx context.Context, nim string) (*entity.Responden, error)
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
