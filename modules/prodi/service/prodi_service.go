package service

import (
	"context"
	"log"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/common/errors"
	"tracer-study-grpc/common/utils"
	"tracer-study-grpc/modules/prodi/entity"
	"tracer-study-grpc/modules/prodi/repository"
)

type ProdiService struct {
	cfg             config.Config
	prodiRepository repository.ProdiRepositoryUseCase
}

type ProdiServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Prodi, error)
	FindProdiByKode(ctx context.Context, kode string) (*entity.Prodi, error)
	Create(ctx context.Context, kode, kodeDikti, kodeFakultas, kodeIntegrasi, nama, jenjang string) (*entity.Prodi, error)
	Update(ctx context.Context, kode string, fields *entity.Prodi) (*entity.Prodi, error)
	Delete(ctx context.Context, kode string) error
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
		log.Println("ERROR: [ProdiService - FindAll] Error while find all prodi: ", err)
		return nil, err
	}

	return res, nil
}

func (svc *ProdiService) FindProdiByKode(ctx context.Context, kode string) (*entity.Prodi, error) {
	res, err := svc.prodiRepository.FindProdiByKode(ctx, kode)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiService - FindProdiByKode] Error while find prodi by kode: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ProdiService) Create(ctx context.Context, kode, kodeDikti, kodeFakultas, kodeIntegrasi, nama, jenjang string) (*entity.Prodi, error) {
	prodi := &entity.Prodi{
		Kode:          kode,
		KodeDikti:     kodeDikti,
		KodeIntegrasi: kodeIntegrasi,
		Nama:          nama,
		Jenjang:       jenjang,
		KodeFakultas:  kodeFakultas,
	}

	res, err := svc.prodiRepository.Create(ctx, prodi)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiService - Create] Error while create prodi: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ProdiService) Update(ctx context.Context, kode string, fields *entity.Prodi) (*entity.Prodi, error) {
	prodi, err := svc.prodiRepository.FindProdiByKode(ctx, kode)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiService - Update] Error while find prodi by kode: ", parseError.Message)
		return nil, err
	}

	updatedMap := make(map[string]interface{})

	utils.AddItemToMap(updatedMap, "kode_dikti", fields.KodeDikti)
	utils.AddItemToMap(updatedMap, "kode_integrasi", fields.KodeIntegrasi)
	utils.AddItemToMap(updatedMap, "nama", fields.Nama)
	utils.AddItemToMap(updatedMap, "jenjang", fields.Jenjang)
	utils.AddItemToMap(updatedMap, "kode_fakultas", fields.KodeFakultas)

	res, err := svc.prodiRepository.Update(ctx, prodi, updatedMap)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiService - Update] Error while update prodi: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ProdiService) Delete(ctx context.Context, kode string) error {
	_, err := svc.prodiRepository.FindProdiByKode(ctx, kode)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiService - Delete] Error while find prodi by kode: ", parseError.Message)
		return err
	}

	err = svc.prodiRepository.Delete(ctx, kode)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiService - Delete] Error while delete prodi: ", parseError.Message)
		return err
	}

	return nil
}
