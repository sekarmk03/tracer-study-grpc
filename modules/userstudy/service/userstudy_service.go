package service

import (
	"context"
	"log"
	"time"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/common/errors"
	"tracer-study-grpc/common/utils"
	"tracer-study-grpc/modules/userstudy/entity"
	"tracer-study-grpc/modules/userstudy/repository"
)

type UserStudyService struct {
	cfg                 config.Config
	userStudyRepository repository.UserStudyRepositoryUseCase
}

type UserStudyServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.UserStudy, error)
	FindByNim(ctx context.Context, nim, emailResponden, hpResponden string) (*entity.UserStudy, error)
	Update(ctx context.Context, nim, emailResponden, hpResponden string, fields *entity.UserStudy) (*entity.UserStudy, error)
	Create(ctx context.Context, namaResponden, emailResponden, hpResponden, namaInstansi, jabatan, alamatInstansi, nimLulusan, namaLulusan, prodiLulusan, tahunLulusan string) (*entity.UserStudy, error)
}

func NewUserStudyService(cfg config.Config, userStudyRepository repository.UserStudyRepositoryUseCase) *UserStudyService {
	return &UserStudyService{
		cfg:                 cfg,
		userStudyRepository: userStudyRepository,
	}
}

func (svc *UserStudyService) FindAll(ctx context.Context, req any) ([]*entity.UserStudy, error) {
	res, err := svc.userStudyRepository.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyService-FindAll] Error while find all user study:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *UserStudyService) FindByNim(ctx context.Context, nim, emailResponden, hpResponden string) (*entity.UserStudy, error) {
	res, err := svc.userStudyRepository.FindByNim(ctx, nim, emailResponden, hpResponden)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyService-FindByNim] Error while find user study by nim:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *UserStudyService) Update(ctx context.Context, nim, emailResponden, hpResponden string, fields *entity.UserStudy) (*entity.UserStudy, error) {
	usrStd, err := svc.userStudyRepository.FindByNim(ctx, nim, emailResponden, hpResponden)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyService-FindByNim] Error while find user study by nim:", parseError.Message)
		return nil, err
	}

	updateMap := map[string]interface{}{
		"updated_at": time.Now(),
		"updated_by": nim,
	}

	utils.AddItemToMap(updateMap, "nama_instansi", fields.NamaInstansi)
	utils.AddItemToMap(updateMap, "jabatan", fields.Jabatan)
	utils.AddItemToMap(updateMap, "alamat_instansi", fields.AlamatInstansi)
	utils.AddItemToMap(updateMap, "lama_mengenal_lulusan", fields.LamaMengenalLulusan)
	utils.AddItemToMap(updateMap, "etika", fields.Etika)
	utils.AddItemToMap(updateMap, "keahlian_bid_ilmu", fields.KeahlianBidIlmu)
	utils.AddItemToMap(updateMap, "bahasa_inggris", fields.BahasaInggris)
	utils.AddItemToMap(updateMap, "penggunaan_ti", fields.PenggunaanTi)
	utils.AddItemToMap(updateMap, "komunikasi", fields.Komunikasi)
	utils.AddItemToMap(updateMap, "kerjasama_tim", fields.KerjasamaTim)
	utils.AddItemToMap(updateMap, "pengembangan_diri", fields.PengembanganDiri)
	utils.AddItemToMap(updateMap, "kesiapan_terjun_masy", fields.KesiapanTerjunMasy)
	utils.AddItemToMap(updateMap, "keunggulan_lulusan", fields.KeunggulanLulusan)
	utils.AddItemToMap(updateMap, "kelemahan_lulusan", fields.KelemahanLulusan)
	utils.AddItemToMap(updateMap, "saran_peningkatan_kompetensi_lulusan", fields.SaranPeningkatanKompetensiLulusan)
	utils.AddItemToMap(updateMap, "saran_perbaikan_kurikulum", fields.SaranPerbaikanKurikulum)

	res, err := svc.userStudyRepository.Update(ctx, usrStd, updateMap)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyService-Update] Error while update user study:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *UserStudyService) Create(ctx context.Context, namaResponden, emailResponden, hpResponden, namaInstansi, jabatan, alamatInstansi, nimLulusan, namaLulusan, prodiLulusan, tahunLulusan string) (*entity.UserStudy, error) {
	reqEntity := &entity.UserStudy{
		NamaResponden:  namaResponden,
		EmailResponden: emailResponden,
		HpResponden:    hpResponden,
		NamaInstansi:   namaInstansi,
		Jabatan:        jabatan,
		AlamatInstansi: alamatInstansi,
		NimLulusan:     nimLulusan,
		NamaLulusan:    namaLulusan,
		ProdiLulusan:   prodiLulusan,
		TahunLulusan:   tahunLulusan,
		CreatedAt:      time.Now(),
		CreatedBy:      "system",
		UpdatedAt:      time.Now(),
		UpdatedBy:      "system",
	}

	res, err := svc.userStudyRepository.Create(ctx, reqEntity)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyService-Create] Error while create user study:", parseError.Message)
		return nil, err
	}

	return res, nil
}
