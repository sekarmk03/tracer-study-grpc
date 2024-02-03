package service

import (
	"context"
	"log"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/prodi/entity"
	"tracer-study-grpc/modules/prodi/repository"
)

type ProdiService struct {
	cfg             config.Config
	prodiRepository repository.ProdiRepositoryUseCase
}

type ProdiServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Prodi, error)
	FindAllFakultas(ctx context.Context, req any) ([]*entity.Fakultas, error)
}

func NewProdiService(cfg config.Config, prodiRepository repository.ProdiRepositoryUseCase) *ProdiService {
	return &ProdiService{
		cfg:             cfg,
		prodiRepository: prodiRepository,
	}
}

func (svc *ProdiService) FindAll(ctx context.Context, req any) ([]*entity.Prodi, error) {
	res, err := svc.prodiRepository.FindAll(ctx, req)
	if err != nil {
		log.Println("[ProdiService - FindAll] ERROR While find all prodi: ", err)
		return nil, err
	}

	return res, nil
}

func (svc *ProdiService) FindAllFakultas(ctx context.Context, req any) ([]*entity.Fakultas, error) {
	res, err := svc.prodiRepository.FindAllFakultas(ctx, req)
	if err != nil {
		log.Println("[ProdiService - FindAllFakultas] ERROR While find all fakultas: ", err)
		return nil, err
	}

	return res, nil
}
