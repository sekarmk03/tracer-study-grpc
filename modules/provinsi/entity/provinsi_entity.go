package entity

import (
	"time"
	"tracer-study-grpc/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	ProvinsiTableName = "ref_provinsi"
)

type Provinsi struct {
	IdWil     string         `json:"id_wil"`
	Nama      string         `json:"nama"`
	UMP12     string         `json:"ump12"`
	CreatedAt time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func NewProvinsi(idWil, nama, ump12 string) *Provinsi {
	return &Provinsi{
		IdWil:     idWil,
		Nama:      nama,
		UMP12:     ump12,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (p *Provinsi) TableName() string {
	return ProvinsiTableName
}

func ConvertEntityToProto(p *Provinsi) *pb.Provinsi {
	return &pb.Provinsi{
		IdWil:     p.IdWil,
		Nama:      p.Nama,
		Ump12:     p.UMP12,
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt),
	}
}
