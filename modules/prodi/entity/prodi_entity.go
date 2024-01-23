package entity

import (
	"time"
	"tracer-study-grpc/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	ProdiTableName = "ref_prodi"
)

// Prodi represents the entity for the Prodi model.
type Prodi struct {
	Kode          string         `json:"kode"`
	KodeDikti     string         `json:"kode_dikti"`
	KodeFak       string         `json:"kode_fak"`
	KodeIntegrasi string         `json:"kode_integrasi"`
	Nama          string         `json:"nama"`
	Jenjang       string         `json:"jenjang"`
	NamaFak       string         `json:"nama_fak"`
	CreatedAt     time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func NewProdi(
	kode, kodeDikti, kodeFak, kodeIntegrasi, nama, jenjang, namaFak string,
) *Prodi {
	return &Prodi{
		Kode:          kode,
		KodeDikti:     kodeDikti,
		KodeFak:       kodeFak,
		KodeIntegrasi: kodeIntegrasi,
		Nama:          nama,
		Jenjang:       jenjang,
		NamaFak:       namaFak,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

func (p *Prodi) TableName() string {
	return ProdiTableName
}

func ConvertEntityToProto(p *Prodi) *pb.Prodi {
	return &pb.Prodi{
		Kode:          p.Kode,
		KodeDikti:     p.KodeDikti,
		KodeFak:       p.KodeFak,
		KodeIntegrasi: p.KodeIntegrasi,
		Nama:          p.Nama,
		Jenjang:       p.Jenjang,
		NamaFak:       p.NamaFak,
		CreatedAt:     timestamppb.New(p.CreatedAt),
		UpdatedAt:     timestamppb.New(p.CreatedAt),
	}
}
