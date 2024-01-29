package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
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

	addItemToMap(updateMap, "ipk", fields.Ipk)
	addItemToMap(updateMap, "kodedikti", fields.Kodedikti)
	addItemToMap(updateMap, "jenjang", fields.Jenjang)
	addItemToMap(updateMap, "namaprodi", fields.Namaprodi)
	addItemToMap(updateMap, "namaprodi2", fields.Namaprodi2)
	addItemToMap(updateMap, "kodeprodi", fields.Kodeprodi)
	addItemToMap(updateMap, "kodeprodi2", fields.Kodeprodi2)
	addItemToMap(updateMap, "kodefak", fields.Kodefak)
	addItemToMap(updateMap, "namafak", fields.Namafak)
	addItemToMap(updateMap, "jlrmasuk", fields.Jlrmasuk)
	addItemToMap(updateMap, "thnmasuk", fields.Thnmasuk)
	addItemToMap(updateMap, "thn_ak", fields.ThnAk)
	addItemToMap(updateMap, "lamastd", fields.Lamastd)
	addItemToMap(updateMap, "semester", fields.Semester)
	addItemToMap(updateMap, "email", fields.Email)
	addItemToMap(updateMap, "hp", fields.Hp)
	addItemToMap(updateMap, "nik", fields.Nik)
	addItemToMap(updateMap, "npwp", fields.Npwp)

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

func addItemToMap(m map[string]interface{}, key string, value any) {
	if value != nil && value != "" {
		m[key] = value
	}
}
