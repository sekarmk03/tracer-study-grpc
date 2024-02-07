package service

import (
	"context"
	"log"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/provinsi/entity"
	"tracer-study-grpc/modules/provinsi/repository"
)

type ProvinsiService struct {
	cfg                config.Config
	provinsiRepository repository.ProvinsiRepositoryUseCase
}

type ProvinsiServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Provinsi, error)
}

func NewProvinsiService(cfg config.Config, provinsiRepository repository.ProvinsiRepositoryUseCase) *ProvinsiService {
	return &ProvinsiService{
		cfg:                cfg,
		provinsiRepository: provinsiRepository,
	}
}

func (svc *ProvinsiService) FindAll(ctx context.Context, req any) ([]*entity.Provinsi, error) {
	res, err := svc.provinsiRepository.FindAll(ctx, req)
	if err != nil {
		log.Println("ERROR: [ProvinsiService - FindAll] Error while find all provinsi: ", err)
		return nil, err
	}

	return res, nil
}
