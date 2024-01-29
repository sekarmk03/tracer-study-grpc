package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
	"tracer-study-grpc/common/config"
	mhsbEntity "tracer-study-grpc/modules/mhsbiodata/entity"
	respEntity "tracer-study-grpc/modules/responden/entity"
	"tracer-study-grpc/modules/responden/repository"
)

type RespondenService struct {
	cfg                 config.Config
	respondenRepository repository.RespondenRepositoryUseCase
}

type RespondenServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*respEntity.Responden, error)
	FindByNim(ctx context.Context, nim string) (*respEntity.Responden, error)
	Update(ctx context.Context, nim string, fields *mhsbEntity.MhsBiodata) (*respEntity.Responden, error)
	Create(ctx context.Context, nim string, semester, tipe, nama, kodeprodi, jk, tgl_wisuda, tgl_sidang, thn_sidang string) (*respEntity.Responden, error)
}

func NewRespondenService(cfg config.Config, respondenRepository repository.RespondenRepositoryUseCase) *RespondenService {
	return &RespondenService{
		cfg:                 cfg,
		respondenRepository: respondenRepository,
	}
}

func (svc *RespondenService) FindAll(ctx context.Context, req any) ([]*respEntity.Responden, error) {
	res, err := svc.respondenRepository.FindAll(ctx, req)
	if err != nil {
		log.Println("[RespondenService - FindAll] Error while find all responden: ", err)
		return nil, err
	}

	return res, nil
}

func (scv *RespondenService) FindByNim(ctx context.Context, nim string) (*respEntity.Responden, error) {
	res, err := scv.respondenRepository.FindByNim(ctx, nim)
	if err != nil {
		log.Println("[RespondenService - FindByNim] Error while find responden by nim: ", err)
		return nil, err
	}

	return res, nil
}

func (scv *RespondenService) Update(ctx context.Context, nim string, fields *mhsbEntity.MhsBiodata) (*respEntity.Responden, error) {
	updateMap := map[string]interface{}{
		"ipk":           fields.IPK,
		"kodedikti":     fields.KODEPSTD,
		"jenjang":       fields.JENJANG,
		"namaprodi":     fields.PRODI,
		"namaprodi2":    fields.NAMAPST,
		"kodeprodi":     fields.KODEPST,
		"kodeprodi2":    fields.KODEPST[:4],
		"kodefak":       fields.KODEFAK,
		"namafak":       fields.NAMAFAK,
		"jlrmasuk":      fields.JLRMASUK,
		"thnmasuk":      fields.THNMASUK,
		"thn_ak":        fields.THNMASUK,
		"lamastd":       fields.LAMASTD,
		"updated_at":    time.Now(),
		"status_update": "1",
		"updated_by":    "system",
	}

	res, err := scv.respondenRepository.Update(ctx, nim, updateMap)
	if err != nil {
		log.Println("[RespondenService - Update] Error while update responden: ", err)
		return nil, err
	}

	return res, nil
}

func (scv *RespondenService) Create(ctx context.Context, nim string, semester, tipe, nama, kodeprodi, jk, tgl_wisuda, tgl_sidang, thn_sidang string) (*respEntity.Responden, error) {
	lamaStdInt, err := strconv.ParseUint(semester, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	reqEntity := &respEntity.Responden{
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
