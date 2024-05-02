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
	Id            uint32         `json:"id"`
	Kode          string         `json:"kode"`
	KodeDikti     string         `json:"kode_dikti"`
	KodeIntegrasi string         `json:"kode_integrasi"`
	Nama          string         `json:"nama"`
	Jenjang       string         `json:"jenjang"`
	KodeFakultas  string         `json:"kode_fakultas"`
	CreatedAt     time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func NewProdi(
	kode, kodeDikti, kodeFakultas, kodeIntegrasi, nama, jenjang string,
) *Prodi {
	return &Prodi{
		Kode:          kode,
		KodeDikti:     kodeDikti,
		KodeIntegrasi: kodeIntegrasi,
		Nama:          nama,
		Jenjang:       jenjang,
		KodeFakultas:  kodeFakultas,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

func (r *Prodi) TableName() string {
	return ProdiTableName
}

func ConvertEntityToProto(p *Prodi) *pb.Prodi {
	return &pb.Prodi{
		Kode:          p.Kode,
		KodeDikti:     p.KodeDikti,
		KodeIntegrasi: p.KodeIntegrasi,
		Nama:          p.Nama,
		Jenjang:       p.Jenjang,
		KodeFakultas:  p.KodeFakultas,
		CreatedAt:     timestamppb.New(p.CreatedAt),
		UpdatedAt:     timestamppb.New(p.CreatedAt),
	}
}