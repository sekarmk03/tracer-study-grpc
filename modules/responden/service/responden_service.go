package service

import (
	"context"
	"log"
	"time"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/common/errors"
	"tracer-study-grpc/common/utils"
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
	Update(ctx context.Context, nim string, fields *entity.Responden) (*entity.Responden, error)
	Create(ctx context.Context, nim string) (*entity.Responden, error)
	FindByNimList(ctx context.Context, nimList []string) ([]*entity.Responden, error)
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
		parseError := errors.ParseError(err)
		log.Println("ERROR: [RespondenService - FindAll] Error while find all responden:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *RespondenService) FindByNim(ctx context.Context, nim string) (*entity.Responden, error) {
	res, err := svc.respondenRepository.FindByNim(ctx, nim)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [RespondenService - FindByNim] Error while find responden by nim:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *RespondenService) Update(ctx context.Context, nim string, fields *entity.Responden) (*entity.Responden, error) {
	responden, err := svc.respondenRepository.FindByNim(ctx, nim)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [RespondenService - Update] Error while find responden by nim:", parseError.Message)
		return nil, err
	}

	updateMap := map[string]interface{}{
		"updated_at":    time.Now(),
	}

	utils.AddItemToMap(updateMap, "nama", fields.Nama)
	utils.AddItemToMap(updateMap, "jalur_masuk", fields.JalurMasuk)
	utils.AddItemToMap(updateMap, "tahun_masuk", fields.TahunMasuk)
	utils.AddItemToMap(updateMap, "lama_studi", fields.LamaStudi)
	utils.AddItemToMap(updateMap, "kode_fakultas", fields.KodeFakultas)
	utils.AddItemToMap(updateMap, "kode_prodi", fields.KodeProdi)
	utils.AddItemToMap(updateMap, "jenis_kelamin", fields.JenisKelamin)
	utils.AddItemToMap(updateMap, "email", fields.Email)
	utils.AddItemToMap(updateMap, "hp", fields.Hp)
	utils.AddItemToMap(updateMap, "ipk", fields.Ipk)
	utils.AddItemToMap(updateMap, "nik", fields.Nik)
	utils.AddItemToMap(updateMap, "npwp", fields.Npwp)
	utils.AddItemToMap(updateMap, "tanggal_wisuda", fields.TanggalWisuda)

	res, err := svc.respondenRepository.Update(ctx, responden, updateMap)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [RespondenService - Update] Error while update responden: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *RespondenService) Create(ctx context.Context, nim string) (*entity.Responden, error) {
	reqEntity := &entity.Responden{
		Nim:           nim,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	res, err := svc.respondenRepository.Create(ctx, reqEntity)
	if err != nil {
		log.Println("ERROR: [RespondenService - Create] Error while create responden: ", err)
		return nil, err
	}

	return res, nil
}

func (svc *RespondenService) FindByNimList(ctx context.Context, nimList []string) ([]*entity.Responden, error) {
	res, err := svc.respondenRepository.FindByNimList(ctx, nimList)
	if err != nil {
		log.Println("ERROR: [RespondenService - FindByNimList] Error while find responden by nim list: ", err)
		return nil, err
	}

	return res, nil
}
