package service

import (
	"context"
	"log"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/pkts/entity"
	"tracer-study-grpc/modules/pkts/repository"
)

type PKTSService struct {
	cfg            config.Config
	pktsRepository repository.PKTSRepositoryUseCase
}

type PKTSServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.PKTS, error)
}

func NewPKTSService(cfg config.Config, pktsRepository repository.PKTSRepositoryUseCase) *PKTSService {
	return &PKTSService{
		cfg:            cfg,
		pktsRepository: pktsRepository,
	}
}

func (svc *PKTSService) FindAll(ctx context.Context, req any) ([]*entity.PKTS, error) {
	res, err := svc.pktsRepository.FindAll(ctx, req)
	if err != nil {
		log.Println("[PKTSService - FindAll] Error while find all pkts: ", err)
		return nil, err
	}

	return res, nil
}