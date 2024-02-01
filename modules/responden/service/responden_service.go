package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
	"tracer-study-grpc/common/config"
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
	Create(ctx context.Context, nim string, semester, tipe, nama, kodeprodi, jk, tgl_wisuda, tgl_sidang, thn_sidang string) (*entity.Responden, error)
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

func (scv *RespondenService) Update(ctx context.Context, nim string, fields *entity.Responden) (*entity.Responden, error) {
	updateMap := map[string]interface{}{
		"updated_at":    time.Now(),
		"status_update": "1",
		"updated_by":    "system",
	}

	utils.AddItemToMap(updateMap, "ipk", fields.Ipk)
	utils.AddItemToMap(updateMap, "kodedikti", fields.Kodedikti)
	utils.AddItemToMap(updateMap, "jenjang", fields.Jenjang)
	utils.AddItemToMap(updateMap, "namaprodi", fields.Namaprodi)
	utils.AddItemToMap(updateMap, "namaprodi2", fields.Namaprodi2)
	utils.AddItemToMap(updateMap, "kodeprodi", fields.Kodeprodi)
	utils.AddItemToMap(updateMap, "kodeprodi2", fields.Kodeprodi2)
	utils.AddItemToMap(updateMap, "kodefak", fields.Kodefak)
	utils.AddItemToMap(updateMap, "namafak", fields.Namafak)
	utils.AddItemToMap(updateMap, "jlrmasuk", fields.Jlrmasuk)
	utils.AddItemToMap(updateMap, "thnmasuk", fields.Thnmasuk)
	utils.AddItemToMap(updateMap, "thn_ak", fields.ThnAk)
	utils.AddItemToMap(updateMap, "lamastd", fields.Lamastd)
	utils.AddItemToMap(updateMap, "semester", fields.Semester)
	utils.AddItemToMap(updateMap, "email", fields.Email)
	utils.AddItemToMap(updateMap, "hp", fields.Hp)
	utils.AddItemToMap(updateMap, "nik", fields.Nik)
	utils.AddItemToMap(updateMap, "npwp", fields.Npwp)

	log.Print(updateMap)

	res, err := scv.respondenRepository.Update(ctx, nim, updateMap)
	if err != nil {
		log.Println("[RespondenService - Update] Error while update responden: ", err)
		return nil, err
	}

	return res, nil
}

func (scv *RespondenService) Create(ctx context.Context, nim string, semester, tipe, nama, kodeprodi, jk, tgl_wisuda, tgl_sidang, thn_sidang string) (*entity.Responden, error) {
	lamaStdInt, err := strconv.ParseUint(semester, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	reqEntity := &entity.Responden{
		Nim:          nim,
		Semester:     uint32(lamaStdInt),
		Lamastd:      semester,
		CreatedBy:    nim,
		UpdatedBy:    nim,
		Type:         tipe,
		StatusUpdate: "1",
		Nama:         nama,
		Kodeprodi:    kodeprodi,
		Kodeprodi2:   kodeprodi[:4],
		JK:           jk,
		TglWisuda:    tgl_wisuda,
		TglSidang:    tgl_sidang,
		ThnSidang:    thn_sidang,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	res, err := scv.respondenRepository.Create(ctx, reqEntity)
	if err != nil {
		log.Println("[RespondenService - Create] Error while create responden: ", err)
		return nil, err
	}

	return res, nil
}