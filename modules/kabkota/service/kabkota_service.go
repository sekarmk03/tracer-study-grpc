package service

import (
	"context"
	"log"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/kabkota/entity"
	"tracer-study-grpc/modules/kabkota/repository"
)

type KabKotaService struct {
	cfg               config.Config
	kabkotaRepository repository.KabKotaRepositoryUseCase
}

type KabKotaServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.KabKota, error)
}

func NewKabKotaService(cfg config.Config, kabkotaRepository repository.KabKotaRepositoryUseCase) *KabKotaService {
	return &KabKotaService{
		cfg:               cfg,
		kabkotaRepository: kabkotaRepository,
	}
}

func (svc *KabKotaService) FindAll(ctx context.Context, req any) ([]*entity.KabKota, error) {
	res, err := svc.kabkotaRepository.FindAll(ctx, req)
	if err != nil {
		log.Println("ERROR: [KabKotaService-FindAll] Error while find all kabkota:", err)
		return nil, err
	}

	return res, nil
}
