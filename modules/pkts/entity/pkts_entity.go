package entity

import (
	"time"
	"tracer-study-grpc/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	PKTSTableName = "pkts"
)

type PKTS struct {
	ID                  uint32         `json:"id"`
	Nim                 string         `json:"nim"`
	Kodeprodi           string         `json:"kodeprodi"`
	ThnSidang           string         `json:"thn_sidang"`
	F8                  uint32         `json:"f8"`
	F5_04               uint32         `json:"f5_04"`
	F5_02               string         `json:"f5_02"`
	F5_06               string         `json:"f5_06"`
	F5_05               string         `json:"f5_05"`
	F5a1                string         `json:"f5a1"`
	F5a2                string         `json:"f5a2"`
	F11_01              uint32         `json:"f11_01"`
	F11_02              string         `json:"f11_02"`
	F5b                 string         `json:"f5b"`
	F5c                 uint32         `json:"f5c"`
	F5d                 uint32         `json:"f5d"`
	F18a                uint32         `json:"f18a"`
	F18b                string         `json:"f18b"`
	F18c                string         `json:"f18c"`
	F18d                string         `json:"f18d"`
	F12_01              uint32         `json:"f12_01"`
	F12_02              string         `json:"f12_02"`
	F14                 uint32         `json:"f14"`
	F15                 uint32         `json:"f15"`
	F1761               uint32         `json:"f1761"`
	F1762               uint32         `json:"f1762"`
	F1763               uint32         `json:"f1763"`
	F1764               uint32         `json:"f1764"`
	F1765               uint32         `json:"f1765"`
	F1766               uint32         `json:"f1766"`
	F1767               uint32         `json:"f1767"`
	F1768               uint32         `json:"f1768"`
	F1769               uint32         `json:"f1769"`
	F1770               uint32         `json:"f1770"`
	F1771               uint32         `json:"f1771"`
	F1772               uint32         `json:"f1772"`
	F1773               uint32         `json:"f1773"`
	F1774               uint32         `json:"f1774"`
	F21                 uint32         `json:"f21"`
	F22                 uint32         `json:"f22"`
	F23                 uint32         `json:"f23"`
	F24                 uint32         `json:"f24"`
	F25                 uint32         `json:"f25"`
	F26                 uint32         `json:"f26"`
	F27                 uint32         `json:"f27"`
	F301                uint32         `json:"f301"`
	F302                string         `json:"f302"`
	F303                uint32         `json:"f303"`
	F4_01               string         `json:"f4_01"`
	F4_02               string         `json:"f4_02"`
	F4_03               string         `json:"f4_03"`
	F4_04               string         `json:"f4_04"`
	F4_05               string         `json:"f4_05"`
	F4_06               string         `json:"f4_06"`
	F4_07               string         `json:"f4_07"`
	F4_08               string         `json:"f4_08"`
	F4_09               string         `json:"f4_09"`
	F4_10               string         `json:"f4_10"`
	F4_11               string         `json:"f4_11"`
	F4_12               string         `json:"f4_12"`
	F4_13               string         `json:"f4_13"`
	F4_14               string         `json:"f4_14"`
	F4_15               string         `json:"f4_15"`
	F4_16               string         `json:"f4_16"`
	F6                  uint32         `json:"f6"`
	F7                  uint32         `json:"f7"`
	F7a                 uint32         `json:"f7a"`
	F10_01              uint32         `json:"f10_01"`
	F10_02              string         `json:"f10_02"`
	F16_01              string         `json:"f16_01"`
	F16_02              string         `json:"f16_02"`
	F16_03              string         `json:"f16_03"`
	F16_04              string         `json:"f16_04"`
	F16_05              string         `json:"f16_05"`
	F16_06              string         `json:"f16_06"`
	F16_07              string         `json:"f16_07"`
	F16_08              string         `json:"f16_08"`
	F16_09              string         `json:"f16_09"`
	F16_10              string         `json:"f16_10"`
	F16_11              string         `json:"f16_11"`
	F16_12              string         `json:"f16_12"`
	F16_13              string         `json:"f16_13"`
	F16_14              string         `json:"f16_14"`
	NamaAtasan          string         `json:"nama_atasan"`
	HpAtasan            string         `json:"hp_atasan"`
	EmailAtasan         string         `json:"email_atasan"`
	TinggalSelamaKuliah string         `json:"tinggal_selama_kuliah"`
	Code                string         `json:"code"`
	MailSent            string         `json:"mail_sent"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (p *PKTS) TableName() string {
	return PKTSTableName
}

func ConvertEntityToProto(p *PKTS) *pb.PKTS {
	return &pb.PKTS{
		Id:                  p.ID,
		Nim:                 p.Nim,
		Kodeprodi:           p.Kodeprodi,
		ThnSidang:           p.ThnSidang,
		F8:                  p.F8,
		F5_04:               p.F5_04,
		F5_02:               p.F5_02,
		F5_06:               p.F5_06,
		F5_05:               p.F5_05,
		F5A1:                p.F5a1,
		F5A2:                p.F5a2,
		F11_01:              p.F11_01,
		F11_02:              p.F11_02,
		F5B:                 p.F5b,
		F5C:                 p.F5c,
		F5D:                 p.F5d,
		F18A:                p.F18a,
		F18B:                p.F18b,
		F18C:                p.F18c,
		F18D:                p.F18d,
		F12_01:              p.F12_01,
		F12_02:              p.F12_02,
		F14:                 p.F14,
		F15:                 p.F15,
		F1761:               p.F1761,
		F1762:               p.F1762,
		F1763:               p.F1763,
		F1764:               p.F1764,
		F1765:               p.F1765,
		F1766:               p.F1766,
		F1767:               p.F1767,
		F1768:               p.F1768,
		F1769:               p.F1769,
		F1770:               p.F1770,
		F1771:               p.F1771,
		F1772:               p.F1772,
		F1773:               p.F1773,
		F1774:               p.F1774,
		F21:                 p.F21,
		F22:                 p.F22,
		F23:                 p.F23,
		F24:                 p.F24,
		F25:                 p.F25,
		F26:                 p.F26,
		F27:                 p.F27,
		F301:                p.F301,
		F302:                p.F302,
		F303:                p.F303,
		F4_01:               p.F4_01,
		F4_02:               p.F4_02,
		F4_03:               p.F4_03,
		F4_04:               p.F4_04,
		F4_05:               p.F4_05,
		F4_06:               p.F4_06,
		F4_07:               p.F4_07,
		F4_08:               p.F4_08,
		F4_09:               p.F4_09,
		F4_10:               p.F4_10,
		F4_11:               p.F4_11,
		F4_12:               p.F4_12,
		F4_13:               p.F4_13,
		F4_14:               p.F4_14,
		F4_15:               p.F4_15,
		F4_16:               p.F4_16,
		F6:                  p.F6,
		F7:                  p.F7,
		F7A:                 p.F7a,
		F10_01:              p.F10_01,
		F10_02:              p.F10_02,
		F16_01:              p.F16_01,
		F16_02:              p.F16_02,
		F16_03:              p.F16_03,
		F16_04:              p.F16_04,
		F16_05:              p.F16_05,
		F16_06:              p.F16_06,
		F16_07:              p.F16_07,
		F16_08:              p.F16_08,
		F16_09:              p.F16_09,
		F16_10:              p.F16_10,
		F16_11:              p.F16_11,
		F16_12:              p.F16_12,
		F16_13:              p.F16_13,
		F16_14:              p.F16_14,
		NamaAtasan:          p.NamaAtasan,
		HpAtasan:            p.HpAtasan,
		EmailAtasan:         p.EmailAtasan,
		TinggalSelamaKuliah: p.TinggalSelamaKuliah,
		Code:                p.Code,
		MailSent:            p.MailSent,
		CreatedAt:           timestamppb.New(p.CreatedAt),
		UpdatedAt:           timestamppb.New(p.UpdatedAt),
	}
}

func ConvertProtoToEntity(p *pb.PKTS) *PKTS {
	return &PKTS{
		ID:                  p.Id,
		Nim:                 p.Nim,
		Kodeprodi:           p.Kodeprodi,
		ThnSidang:           p.ThnSidang,
		F8:                  p.F8,
		F5_04:               p.F5_04,
		F5_02:               p.F5_02,
		F5_06:               p.F5_06,
		F5_05:               p.F5_05,
		F5a1:                p.F5A1,
		F5a2:                p.F5A2,
		F11_01:              p.F11_01,
		F11_02:              p.F11_02,
		F5b:                 p.F5B,
		F5c:                 p.F5C,
		F5d:                 p.F5D,
		F18a:                p.F18A,
		F18b:                p.F18B,
		F18c:                p.F18C,
		F18d:                p.F18D,
		F12_01:              p.F12_01,
		F12_02:              p.F12_02,
		F14:                 p.F14,
		F15:                 p.F15,
		F1761:               p.F1761,
		F1762:               p.F1762,
		F1763:               p.F1763,
		F1764:               p.F1764,
		F1765:               p.F1765,
		F1766:               p.F1766,
		F1767:               p.F1767,
		F1768:               p.F1768,
		F1769:               p.F1769,
		F1770:               p.F1770,
		F1771:               p.F1771,
		F1772:               p.F1772,
		F1773:               p.F1773,
		F1774:               p.F1774,
		F21:                 p.F21,
		F22:                 p.F22,
		F23:                 p.F23,
		F24:                 p.F24,
		F25:                 p.F25,
		F26:                 p.F26,
		F27:                 p.F27,
		F301:                p.F301,
		F302:                p.F302,
		F303:                p.F303,
		F4_01:               p.F4_01,
		F4_02:               p.F4_02,
		F4_03:               p.F4_03,
		F4_04:               p.F4_04,
		F4_05:               p.F4_05,
		F4_06:               p.F4_06,
		F4_07:               p.F4_07,
		F4_08:               p.F4_08,
		F4_09:               p.F4_09,
		F4_10:               p.F4_10,
		F4_11:               p.F4_11,
		F4_12:               p.F4_12,
		F4_13:               p.F4_13,
		F4_14:               p.F4_14,
		F4_15:               p.F4_15,
		F4_16:               p.F4_16,
		F6:                  p.F6,
		F7:                  p.F7,
		F7a:                 p.F7A,
		F10_01:              p.F10_01,
		F10_02:              p.F10_02,
		F16_01:              p.F16_01,
		F16_02:              p.F16_02,
		F16_03:              p.F16_03,
		F16_04:              p.F16_04,
		F16_05:              p.F16_05,
		F16_06:              p.F16_06,
		F16_07:              p.F16_07,
		F16_08:              p.F16_08,
		F16_09:              p.F16_09,
		F16_10:              p.F16_10,
		F16_11:              p.F16_11,
		F16_12:              p.F16_12,
		F16_13:              p.F16_13,
		F16_14:              p.F16_14,
		NamaAtasan:          p.NamaAtasan,
		HpAtasan:            p.HpAtasan,
		EmailAtasan:         p.EmailAtasan,
		TinggalSelamaKuliah: p.TinggalSelamaKuliah,
		Code:                p.Code,
		MailSent:            p.MailSent,
		CreatedAt:           p.CreatedAt.AsTime(),
		UpdatedAt:           p.UpdatedAt.AsTime(),
	}
}
