package service

import (
	"context"
	"log"
	"time"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/common/errors"
	"tracer-study-grpc/common/utils"
	"tracer-study-grpc/modules/pkts/entity"
	"tracer-study-grpc/modules/pkts/repository"
)

type PKTSService struct {
	cfg            config.Config
	pktsRepository repository.PKTSRepositoryUseCase
}

type PKTSServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.PKTS, error)
	FindByNim(ctx context.Context, nim string) (*entity.PKTS, error)
	Create(ctx context.Context, nim, kodeprodi, thnSidang string, f8, f5_04 uint32, f5_02, f5_06, f5_05, f5a1, f5a2 string, f11_01 uint32, f11_02, f5b string, f5c, f5d, f18a uint32, f18b, f18c, f18d string, f12_01 uint32, f12_02 string, f14, f15, f1761, f1762, f1763, f1764, f1765, f1766, f1767, f1768, f1769, f1770, f1771, f1772, f1773, f1774, f21, f22, f23, f24, f25, f26, f27, f301 uint32, f302 string, f303 uint32, f4_01, f4_02, f4_03, f4_04, f4_05, f4_06, f4_07, f4_08, f4_09, f4_10, f4_11, f4_12, f4_13, f4_14, f4_15, f4_16 string, f6, f7, f7a, f10_01 uint32, f10_02, f16_01, f16_02, f16_03, f16_04, f16_05, f16_06, f16_07, f16_08, f16_09, f16_10, f16_11, f16_12, f16_13, f16_14, namaAtasan, hpAtasan, emailAtasan, tinggalSelamaKuliah string) (*entity.PKTS, error)
	Update(ctx context.Context, nim string, fields *entity.PKTS) (*entity.PKTS, error)
	FindByAtasan(ctx context.Context, namaA, hpA, emailA string) ([]*string, error)
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
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSService-FindAll] Error while find all pkts:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *PKTSService) FindByNim(ctx context.Context, nim string) (*entity.PKTS, error) {
	res, err := svc.pktsRepository.FindByNim(ctx, nim)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSService-FindByNim] Error while find pkts by nim:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *PKTSService) Create(ctx context.Context, nim, kodeprodi, thnSidang string, f8, f5_04 uint32, f5_02, f5_06, f5_05, f5a1, f5a2 string, f11_01 uint32, f11_02, f5b string, f5c, f5d, f18a uint32, f18b, f18c, f18d string, f12_01 uint32, f12_02 string, f14, f15, f1761, f1762, f1763, f1764, f1765, f1766, f1767, f1768, f1769, f1770, f1771, f1772, f1773, f1774, f21, f22, f23, f24, f25, f26, f27, f301 uint32, f302 string, f303 uint32, f4_01, f4_02, f4_03, f4_04, f4_05, f4_06, f4_07, f4_08, f4_09, f4_10, f4_11, f4_12, f4_13, f4_14, f4_15, f4_16 string, f6, f7, f7a, f10_01 uint32, f10_02, f16_01, f16_02, f16_03, f16_04, f16_05, f16_06, f16_07, f16_08, f16_09, f16_10, f16_11, f16_12, f16_13, f16_14, namaAtasan, hpAtasan, emailAtasan, tinggalSelamaKuliah string) (*entity.PKTS, error) {
	reqEntity := &entity.PKTS{
		Nim:                 nim,
		Kodeprodi:           kodeprodi,
		ThnSidang:           thnSidang,
		F8:                  f8,
		F5_04:               f5_04,
		F5_02:               f5_02,
		F5_06:               f5_06,
		F5_05:               f5_05,
		F5a1:                f5a1,
		F5a2:                f5a2,
		F11_01:              f11_01,
		F11_02:              f11_02,
		F5b:                 f5b,
		F5c:                 f5c,
		F5d:                 f5d,
		F18a:                f18a,
		F18b:                f18b,
		F18c:                f18c,
		F18d:                f18d,
		F12_01:              f12_01,
		F12_02:              f12_02,
		F14:                 f14,
		F15:                 f15,
		F1761:               f1761,
		F1762:               f1762,
		F1763:               f1763,
		F1764:               f1764,
		F1765:               f1765,
		F1766:               f1766,
		F1767:               f1767,
		F1768:               f1768,
		F1769:               f1769,
		F1770:               f1770,
		F1771:               f1771,
		F1772:               f1772,
		F1773:               f1773,
		F1774:               f1774,
		F21:                 f21,
		F22:                 f22,
		F23:                 f23,
		F24:                 f24,
		F25:                 f25,
		F26:                 f26,
		F27:                 f27,
		F301:                f301,
		F302:                f302,
		F303:                f303,
		F4_01:               f4_01,
		F4_02:               f4_02,
		F4_03:               f4_03,
		F4_04:               f4_04,
		F4_05:               f4_05,
		F4_06:               f4_06,
		F4_07:               f4_07,
		F4_08:               f4_08,
		F4_09:               f4_09,
		F4_10:               f4_10,
		F4_11:               f4_11,
		F4_12:               f4_12,
		F4_13:               f4_13,
		F4_14:               f4_14,
		F4_15:               f4_15,
		F4_16:               f4_16,
		F6:                  f6,
		F7:                  f7,
		F7a:                 f7a,
		F10_01:              f10_01,
		F10_02:              f10_02,
		F16_01:              f16_01,
		F16_02:              f16_02,
		F16_03:              f16_03,
		F16_04:              f16_04,
		F16_05:              f16_05,
		F16_06:              f16_06,
		F16_07:              f16_07,
		F16_08:              f16_08,
		F16_09:              f16_09,
		F16_10:              f16_10,
		F16_11:              f16_11,
		F16_12:              f16_12,
		F16_13:              f16_13,
		F16_14:              f16_14,
		NamaAtasan:          namaAtasan,
		HpAtasan:            hpAtasan,
		EmailAtasan:         emailAtasan,
		TinggalSelamaKuliah: tinggalSelamaKuliah,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	res, err := svc.pktsRepository.Create(ctx, reqEntity)
	if err != nil {
		log.Println("ERROR: [PKTSService-Create] Error while create pkts:", err)
		return nil, err
	}

	return res, nil
}

func (svc *PKTSService) Update(ctx context.Context, nim string, fields *entity.PKTS) (*entity.PKTS, error) {
	updateMap := make(map[string]interface{})

	utils.AddItemToMap(updateMap, "f8", fields.F8)
	utils.AddItemToMap(updateMap, "f5_04", fields.F5_04)
	utils.AddItemToMap(updateMap, "f5_02", fields.F5_02)
	utils.AddItemToMap(updateMap, "f5_06", fields.F5_06)
	utils.AddItemToMap(updateMap, "f5_05", fields.F5_05)
	utils.AddItemToMap(updateMap, "f5a1", fields.F5a1)
	utils.AddItemToMap(updateMap, "f5a2", fields.F5a2)
	utils.AddItemToMap(updateMap, "f11_01", fields.F11_01)
	utils.AddItemToMap(updateMap, "f11_02", fields.F11_02)
	utils.AddItemToMap(updateMap, "f5b", fields.F5b)
	utils.AddItemToMap(updateMap, "f5c", fields.F5c)
	utils.AddItemToMap(updateMap, "f5d", fields.F5d)
	utils.AddItemToMap(updateMap, "f18a", fields.F18a)
	utils.AddItemToMap(updateMap, "f18b", fields.F18b)
	utils.AddItemToMap(updateMap, "f18c", fields.F18c)
	utils.AddItemToMap(updateMap, "f18d", fields.F18d)
	utils.AddItemToMap(updateMap, "f12_01", fields.F12_01)
	utils.AddItemToMap(updateMap, "f12_02", fields.F12_02)
	utils.AddItemToMap(updateMap, "f14", fields.F14)
	utils.AddItemToMap(updateMap, "f15", fields.F15)
	utils.AddItemToMap(updateMap, "f1761", fields.F1761)
	utils.AddItemToMap(updateMap, "f1762", fields.F1762)
	utils.AddItemToMap(updateMap, "f1763", fields.F1763)
	utils.AddItemToMap(updateMap, "f1764", fields.F1764)
	utils.AddItemToMap(updateMap, "f1765", fields.F1765)
	utils.AddItemToMap(updateMap, "f1766", fields.F1766)
	utils.AddItemToMap(updateMap, "f1767", fields.F1767)
	utils.AddItemToMap(updateMap, "f1768", fields.F1768)
	utils.AddItemToMap(updateMap, "f1769", fields.F1769)
	utils.AddItemToMap(updateMap, "f1770", fields.F1770)
	utils.AddItemToMap(updateMap, "f1771", fields.F1771)
	utils.AddItemToMap(updateMap, "f1772", fields.F1772)
	utils.AddItemToMap(updateMap, "f1773", fields.F1773)
	utils.AddItemToMap(updateMap, "f1774", fields.F1774)
	utils.AddItemToMap(updateMap, "f21", fields.F21)
	utils.AddItemToMap(updateMap, "f22", fields.F22)
	utils.AddItemToMap(updateMap, "f23", fields.F23)
	utils.AddItemToMap(updateMap, "f24", fields.F24)
	utils.AddItemToMap(updateMap, "f25", fields.F25)
	utils.AddItemToMap(updateMap, "f26", fields.F26)
	utils.AddItemToMap(updateMap, "f27", fields.F27)
	utils.AddItemToMap(updateMap, "f301", fields.F301)
	utils.AddItemToMap(updateMap, "f302", fields.F302)
	utils.AddItemToMap(updateMap, "f303", fields.F303)
	utils.AddItemToMap(updateMap, "f4_01", fields.F4_01)
	utils.AddItemToMap(updateMap, "f4_02", fields.F4_02)
	utils.AddItemToMap(updateMap, "f4_03", fields.F4_03)
	utils.AddItemToMap(updateMap, "f4_04", fields.F4_04)
	utils.AddItemToMap(updateMap, "f4_05", fields.F4_05)
	utils.AddItemToMap(updateMap, "f4_06", fields.F4_06)
	utils.AddItemToMap(updateMap, "f4_07", fields.F4_07)
	utils.AddItemToMap(updateMap, "f4_08", fields.F4_08)
	utils.AddItemToMap(updateMap, "f4_09", fields.F4_09)
	utils.AddItemToMap(updateMap, "f4_10", fields.F4_10)
	utils.AddItemToMap(updateMap, "f4_11", fields.F4_11)
	utils.AddItemToMap(updateMap, "f4_12", fields.F4_12)
	utils.AddItemToMap(updateMap, "f4_13", fields.F4_13)
	utils.AddItemToMap(updateMap, "f4_14", fields.F4_14)
	utils.AddItemToMap(updateMap, "f4_15", fields.F4_15)
	utils.AddItemToMap(updateMap, "f4_16", fields.F4_16)
	utils.AddItemToMap(updateMap, "f6", fields.F6)
	utils.AddItemToMap(updateMap, "f7", fields.F7)
	utils.AddItemToMap(updateMap, "f7a", fields.F7a)
	utils.AddItemToMap(updateMap, "f10_01", fields.F10_01)
	utils.AddItemToMap(updateMap, "f10_02", fields.F10_02)
	utils.AddItemToMap(updateMap, "f16_01", fields.F16_01)
	utils.AddItemToMap(updateMap, "f16_02", fields.F16_02)
	utils.AddItemToMap(updateMap, "f16_03", fields.F16_03)
	utils.AddItemToMap(updateMap, "f16_04", fields.F16_04)
	utils.AddItemToMap(updateMap, "f16_05", fields.F16_05)
	utils.AddItemToMap(updateMap, "f16_06", fields.F16_06)
	utils.AddItemToMap(updateMap, "f16_07", fields.F16_07)
	utils.AddItemToMap(updateMap, "f16_08", fields.F16_08)
	utils.AddItemToMap(updateMap, "f16_09", fields.F16_09)
	utils.AddItemToMap(updateMap, "f16_10", fields.F16_10)
	utils.AddItemToMap(updateMap, "f16_11", fields.F16_11)
	utils.AddItemToMap(updateMap, "f16_12", fields.F16_12)
	utils.AddItemToMap(updateMap, "f16_13", fields.F16_13)
	utils.AddItemToMap(updateMap, "f16_14", fields.F16_14)
	utils.AddItemToMap(updateMap, "nama_atasan", fields.NamaAtasan)
	utils.AddItemToMap(updateMap, "hp_atasan", fields.HpAtasan)
	utils.AddItemToMap(updateMap, "email_atasan", fields.EmailAtasan)
	utils.AddItemToMap(updateMap, "tinggal_selama_kuliah", fields.TinggalSelamaKuliah)

	log.Print(updateMap)

	res, err := svc.pktsRepository.Update(ctx, nim, updateMap)
	if err != nil {
		log.Println("ERROR: [PKTSService - Update] Error while update pkts:", err)
		return nil, err
	}

	return res, nil
}

func (svc *PKTSService) FindByAtasan(ctx context.Context, namaA, hpA, emailA string) ([]*string, error) {
	res, err := svc.pktsRepository.FindByAtasan(ctx, namaA, hpA, emailA)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSService-FindByAtasan] Error while find pkts by atasan:", parseError.Message)
		return nil, err
	}

	return res, nil
}
