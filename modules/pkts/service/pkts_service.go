package service

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/common/errors"
	"tracer-study-grpc/common/utils"
	"tracer-study-grpc/modules/pkts/entity"
	"tracer-study-grpc/modules/pkts/repository"

	"github.com/xuri/excelize/v2"
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
	ExportPKTSReport(ctx context.Context, req any) (*bytes.Buffer, error)
	FindPKTSRekap(ctx context.Context, kodeprodi string) ([]*entity.PKTSRekap, error)
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
	pkts, err := svc.pktsRepository.FindByNim(ctx, nim)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSService-Update] Error while find pkts by nim:", parseError.Message)
		return nil, err
	}

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

	res, err := svc.pktsRepository.Update(ctx, pkts, updateMap)
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

func (svc *PKTSService) ExportPKTSReport(ctx context.Context, req any) (*bytes.Buffer, error) {
	rows, err := svc.pktsRepository.FindAllReport(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSService-FindAll] Error while find all pkts:", parseError.Message)
		return nil, err
	}

	file := excelize.NewFile()
	sheetName := "Sheet1"
	index, err := file.NewSheet(sheetName)
	if err != nil {
		log.Println("ERROR: [PKTSService-ExportPKTSReport] Error while create new sheet:", err)
		return nil, err
	}

	file.SetCellValue(sheetName, "A1", "created_at")
	file.SetCellValue(sheetName, "B1", "updated_at")
	file.SetCellValue(sheetName, "C1", "kodeprodi")
	file.SetCellValue(sheetName, "D1", "nama_prodi")
	file.SetCellValue(sheetName, "E1", "kdptimsmh")
	file.SetCellValue(sheetName, "F1", "kdpstmsmh")
	file.SetCellValue(sheetName, "G1", "jenjang")
	file.SetCellValue(sheetName, "H1", "nimhsmsmh")
	file.SetCellValue(sheetName, "I1", "nmmhsmsmh")
	file.SetCellValue(sheetName, "J1", "jenis_kelamin")
	file.SetCellValue(sheetName, "K1", "telpomsmh")
	file.SetCellValue(sheetName, "L1", "emailmsmh")
	file.SetCellValue(sheetName, "M1", "thn_lulus")
	file.SetCellValue(sheetName, "N1", "nik")
	file.SetCellValue(sheetName, "O1", "npwp")
	file.SetCellValue(sheetName, "P1", "f8")
	file.SetCellValue(sheetName, "Q1", "f504")
	file.SetCellValue(sheetName, "R1", "f502")
	file.SetCellValue(sheetName, "S1", "f505")
	file.SetCellValue(sheetName, "T1", "f506")
	file.SetCellValue(sheetName, "U1", "f5a1")
	file.SetCellValue(sheetName, "V1", "f5a2")
	file.SetCellValue(sheetName, "W1", "f1101")
	file.SetCellValue(sheetName, "X1", "f1102")
	file.SetCellValue(sheetName, "Y1", "f5b")
	file.SetCellValue(sheetName, "Z1", "f5c")
	file.SetCellValue(sheetName, "AA1", "f5d")
	file.SetCellValue(sheetName, "AB1", "f18a")
	file.SetCellValue(sheetName, "AC1", "f18b")
	file.SetCellValue(sheetName, "AD1", "f18c")
	file.SetCellValue(sheetName, "AE1", "f18d")
	file.SetCellValue(sheetName, "AF1", "f1201")
	file.SetCellValue(sheetName, "AG1", "f1202")
	file.SetCellValue(sheetName, "AH1", "f14")
	file.SetCellValue(sheetName, "AI1", "f15")
	file.SetCellValue(sheetName, "AJ1", "f1761")
	file.SetCellValue(sheetName, "AK1", "f1762")
	file.SetCellValue(sheetName, "AL1", "f1763")
	file.SetCellValue(sheetName, "AM1", "f1764")
	file.SetCellValue(sheetName, "AN1", "f1765")
	file.SetCellValue(sheetName, "AO1", "f1766")
	file.SetCellValue(sheetName, "AP1", "f1767")
	file.SetCellValue(sheetName, "AQ1", "f1768")
	file.SetCellValue(sheetName, "AR1", "f1769")
	file.SetCellValue(sheetName, "AS1", "f1770")
	file.SetCellValue(sheetName, "AT1", "f1771")
	file.SetCellValue(sheetName, "AU1", "f1772")
	file.SetCellValue(sheetName, "AV1", "f1773")
	file.SetCellValue(sheetName, "AW1", "f1774")
	file.SetCellValue(sheetName, "AX1", "f21")
	file.SetCellValue(sheetName, "AY1", "f22")
	file.SetCellValue(sheetName, "AZ1", "f23")
	file.SetCellValue(sheetName, "BA1", "f24")
	file.SetCellValue(sheetName, "BB1", "f25")
	file.SetCellValue(sheetName, "BC1", "f26")
	file.SetCellValue(sheetName, "BD1", "f27")
	file.SetCellValue(sheetName, "BE1", "f301")
	file.SetCellValue(sheetName, "BF1", "f302")
	file.SetCellValue(sheetName, "BG1", "f303")
	file.SetCellValue(sheetName, "BH1", "f401")
	file.SetCellValue(sheetName, "BI1", "f402")
	file.SetCellValue(sheetName, "BJ1", "f403")
	file.SetCellValue(sheetName, "BK1", "f404")
	file.SetCellValue(sheetName, "BL1", "f405")
	file.SetCellValue(sheetName, "BM1", "f406")
	file.SetCellValue(sheetName, "BN1", "f407")
	file.SetCellValue(sheetName, "BO1", "f408")
	file.SetCellValue(sheetName, "BP1", "f409")
	file.SetCellValue(sheetName, "BQ1", "f410")
	file.SetCellValue(sheetName, "BR1", "f411")
	file.SetCellValue(sheetName, "BS1", "f412")
	file.SetCellValue(sheetName, "BT1", "f413")
	file.SetCellValue(sheetName, "BU1", "f414")
	file.SetCellValue(sheetName, "BV1", "f415")
	file.SetCellValue(sheetName, "BW1", "f416")
	file.SetCellValue(sheetName, "BX1", "f6")
	file.SetCellValue(sheetName, "BY1", "f7")
	file.SetCellValue(sheetName, "BZ1", "f7a")
	file.SetCellValue(sheetName, "CA1", "f1001")
	file.SetCellValue(sheetName, "CB1", "f1002")
	file.SetCellValue(sheetName, "CC1", "f1601")
	file.SetCellValue(sheetName, "CD1", "f1602")
	file.SetCellValue(sheetName, "CE1", "f1603")
	file.SetCellValue(sheetName, "CF1", "f1604")
	file.SetCellValue(sheetName, "CG1", "f1605")
	file.SetCellValue(sheetName, "CH1", "f1606")
	file.SetCellValue(sheetName, "CI1", "f1607")
	file.SetCellValue(sheetName, "CJ1", "f1608")
	file.SetCellValue(sheetName, "CK1", "f1609")
	file.SetCellValue(sheetName, "CL1", "f1610")
	file.SetCellValue(sheetName, "CM1", "f1611")
	file.SetCellValue(sheetName, "CN1", "f1612")
	file.SetCellValue(sheetName, "CO1", "f1613")
	file.SetCellValue(sheetName, "CP1", "f1614")
	file.SetCellValue(sheetName, "CQ1", "nama_atasan")
	file.SetCellValue(sheetName, "CR1", "hp_atasan")
	file.SetCellValue(sheetName, "CS1", "email_atasan")

	kodePT := "001034"

	for i, row := range rows {
		file.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), row.CreatedAt)
		file.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), row.UpdatedAt)
		file.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), row.Kodeprodi)
		file.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), row.NamaProdi)
		file.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), kodePT)
		file.SetCellValue(sheetName, fmt.Sprintf("F%d", i+2), row.KodeDikti)
		file.SetCellValue(sheetName, fmt.Sprintf("G%d", i+2), row.Jenjang)
		file.SetCellValue(sheetName, fmt.Sprintf("H%d", i+2), row.Nim)
		file.SetCellValue(sheetName, fmt.Sprintf("I%d", i+2), row.Nama)
		file.SetCellValue(sheetName, fmt.Sprintf("J%d", i+2), row.JK)
		file.SetCellValue(sheetName, fmt.Sprintf("K%d", i+2), row.Hp)
		file.SetCellValue(sheetName, fmt.Sprintf("L%d", i+2), row.Email)
		file.SetCellValue(sheetName, fmt.Sprintf("M%d", i+2), row.ThnSidang)
		file.SetCellValue(sheetName, fmt.Sprintf("N%d", i+2), row.Nik)
		file.SetCellValue(sheetName, fmt.Sprintf("O%d", i+2), row.Npwp)
		file.SetCellValue(sheetName, fmt.Sprintf("P%d", i+2), row.F8)
		file.SetCellValue(sheetName, fmt.Sprintf("Q%d", i+2), row.F5_04)
		file.SetCellValue(sheetName, fmt.Sprintf("R%d", i+2), row.F5_02)
		file.SetCellValue(sheetName, fmt.Sprintf("S%d", i+2), row.F5_05)
		file.SetCellValue(sheetName, fmt.Sprintf("T%d", i+2), row.F5_06)
		file.SetCellValue(sheetName, fmt.Sprintf("U%d", i+2), row.F5a1)
		file.SetCellValue(sheetName, fmt.Sprintf("V%d", i+2), row.F5a2)
		file.SetCellValue(sheetName, fmt.Sprintf("W%d", i+2), row.F11_01)
		file.SetCellValue(sheetName, fmt.Sprintf("X%d", i+2), row.F11_02)
		file.SetCellValue(sheetName, fmt.Sprintf("Y%d", i+2), row.F5b)
		file.SetCellValue(sheetName, fmt.Sprintf("Z%d", i+2), row.F5c)
		file.SetCellValue(sheetName, fmt.Sprintf("AA%d", i+2), row.F5d)
		file.SetCellValue(sheetName, fmt.Sprintf("AB%d", i+2), row.F18a)
		file.SetCellValue(sheetName, fmt.Sprintf("AC%d", i+2), row.F18b)
		file.SetCellValue(sheetName, fmt.Sprintf("AD%d", i+2), row.F18c)
		file.SetCellValue(sheetName, fmt.Sprintf("AE%d", i+2), row.F18d)
		file.SetCellValue(sheetName, fmt.Sprintf("AF%d", i+2), row.F12_01)
		file.SetCellValue(sheetName, fmt.Sprintf("AG%d", i+2), row.F12_02)
		file.SetCellValue(sheetName, fmt.Sprintf("AH%d", i+2), row.F14)
		file.SetCellValue(sheetName, fmt.Sprintf("AI%d", i+2), row.F15)
		file.SetCellValue(sheetName, fmt.Sprintf("AJ%d", i+2), row.F1761)
		file.SetCellValue(sheetName, fmt.Sprintf("AK%d", i+2), row.F1762)
		file.SetCellValue(sheetName, fmt.Sprintf("AL%d", i+2), row.F1763)
		file.SetCellValue(sheetName, fmt.Sprintf("AM%d", i+2), row.F1764)
		file.SetCellValue(sheetName, fmt.Sprintf("AN%d", i+2), row.F1765)
		file.SetCellValue(sheetName, fmt.Sprintf("AO%d", i+2), row.F1766)
		file.SetCellValue(sheetName, fmt.Sprintf("AP%d", i+2), row.F1767)
		file.SetCellValue(sheetName, fmt.Sprintf("AQ%d", i+2), row.F1768)
		file.SetCellValue(sheetName, fmt.Sprintf("AR%d", i+2), row.F1769)
		file.SetCellValue(sheetName, fmt.Sprintf("AS%d", i+2), row.F1770)
		file.SetCellValue(sheetName, fmt.Sprintf("AT%d", i+2), row.F1771)
		file.SetCellValue(sheetName, fmt.Sprintf("AU%d", i+2), row.F1772)
		file.SetCellValue(sheetName, fmt.Sprintf("AV%d", i+2), row.F1773)
		file.SetCellValue(sheetName, fmt.Sprintf("AW%d", i+2), row.F1774)
		file.SetCellValue(sheetName, fmt.Sprintf("AX%d", i+2), row.F21)
		file.SetCellValue(sheetName, fmt.Sprintf("AY%d", i+2), row.F22)
		file.SetCellValue(sheetName, fmt.Sprintf("AZ%d", i+2), row.F23)
		file.SetCellValue(sheetName, fmt.Sprintf("BA%d", i+2), row.F24)
		file.SetCellValue(sheetName, fmt.Sprintf("BB%d", i+2), row.F25)
		file.SetCellValue(sheetName, fmt.Sprintf("BC%d", i+2), row.F26)
		file.SetCellValue(sheetName, fmt.Sprintf("BD%d", i+2), row.F27)
		file.SetCellValue(sheetName, fmt.Sprintf("BE%d", i+2), row.F301)
		file.SetCellValue(sheetName, fmt.Sprintf("BF%d", i+2), row.F302)
		file.SetCellValue(sheetName, fmt.Sprintf("BG%d", i+2), row.F303)
		file.SetCellValue(sheetName, fmt.Sprintf("BH%d", i+2), row.F4_01)
		file.SetCellValue(sheetName, fmt.Sprintf("BI%d", i+2), row.F4_02)
		file.SetCellValue(sheetName, fmt.Sprintf("BJ%d", i+2), row.F4_03)
		file.SetCellValue(sheetName, fmt.Sprintf("BK%d", i+2), row.F4_04)
		file.SetCellValue(sheetName, fmt.Sprintf("BL%d", i+2), row.F4_05)
		file.SetCellValue(sheetName, fmt.Sprintf("BM%d", i+2), row.F4_06)
		file.SetCellValue(sheetName, fmt.Sprintf("BN%d", i+2), row.F4_07)
		file.SetCellValue(sheetName, fmt.Sprintf("BO%d", i+2), row.F4_08)
		file.SetCellValue(sheetName, fmt.Sprintf("BP%d", i+2), row.F4_09)
		file.SetCellValue(sheetName, fmt.Sprintf("BQ%d", i+2), row.F4_10)
		file.SetCellValue(sheetName, fmt.Sprintf("BR%d", i+2), row.F4_11)
		file.SetCellValue(sheetName, fmt.Sprintf("BS%d", i+2), row.F4_12)
		file.SetCellValue(sheetName, fmt.Sprintf("BT%d", i+2), row.F4_13)
		file.SetCellValue(sheetName, fmt.Sprintf("BU%d", i+2), row.F4_14)
		file.SetCellValue(sheetName, fmt.Sprintf("BV%d", i+2), row.F4_15)
		file.SetCellValue(sheetName, fmt.Sprintf("BW%d", i+2), row.F4_16)
		file.SetCellValue(sheetName, fmt.Sprintf("BX%d", i+2), row.F6)
		file.SetCellValue(sheetName, fmt.Sprintf("BY%d", i+2), row.F7)
		file.SetCellValue(sheetName, fmt.Sprintf("BZ%d", i+2), row.F7a)
		file.SetCellValue(sheetName, fmt.Sprintf("CA%d", i+2), row.F10_01)
		file.SetCellValue(sheetName, fmt.Sprintf("CB%d", i+2), row.F10_02)
		file.SetCellValue(sheetName, fmt.Sprintf("CC%d", i+2), row.F16_01)
		file.SetCellValue(sheetName, fmt.Sprintf("CD%d", i+2), row.F16_02)
		file.SetCellValue(sheetName, fmt.Sprintf("CE%d", i+2), row.F16_03)
		file.SetCellValue(sheetName, fmt.Sprintf("CF%d", i+2), row.F16_04)
		file.SetCellValue(sheetName, fmt.Sprintf("CG%d", i+2), row.F16_05)
		file.SetCellValue(sheetName, fmt.Sprintf("CH%d", i+2), row.F16_06)
		file.SetCellValue(sheetName, fmt.Sprintf("CI%d", i+2), row.F16_07)
		file.SetCellValue(sheetName, fmt.Sprintf("CJ%d", i+2), row.F16_08)
		file.SetCellValue(sheetName, fmt.Sprintf("CK%d", i+2), row.F16_09)
		file.SetCellValue(sheetName, fmt.Sprintf("CL%d", i+2), row.F16_10)
		file.SetCellValue(sheetName, fmt.Sprintf("CM%d", i+2), row.F16_11)
		file.SetCellValue(sheetName, fmt.Sprintf("CN%d", i+2), row.F16_12)
		file.SetCellValue(sheetName, fmt.Sprintf("CO%d", i+2), row.F16_13)
		file.SetCellValue(sheetName, fmt.Sprintf("CP%d", i+2), row.F16_14)
		file.SetCellValue(sheetName, fmt.Sprintf("CQ%d", i+2), row.NamaAtasan)
		file.SetCellValue(sheetName, fmt.Sprintf("CR%d", i+2), row.HpAtasan)
		file.SetCellValue(sheetName, fmt.Sprintf("CS%d", i+2), row.EmailAtasan)
	}

	file.SetActiveSheet(index)

	excelFilePath := "pkts_report_" + strconv.FormatInt(time.Now().Unix(), 10) + ".xlsx"
	if err := file.SaveAs(excelFilePath); err != nil {
		log.Println("ERROR: [PKTSService-ExportPKTSReport] Error while save excel file:", err)
		return nil, err
	}

	buff, err := file.WriteToBuffer()
	if err != nil {
		log.Println("ERROR: [PKTSService-ExportPKTSReport] Error while write to buffer:", err)
		return nil, err
	}

	return buff, nil
}

func (svc *PKTSService) FindPKTSRekap(ctx context.Context, kodeprodi string) ([]*entity.PKTSRekap, error) {
	res, err := svc.pktsRepository.FindPKTSRekap(ctx, kodeprodi)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSService-FindPKTSRekap] Error while find pkts rekap:", parseError.Message)
		return nil, err
	}

	return res, nil
}