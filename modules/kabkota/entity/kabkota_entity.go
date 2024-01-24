package entity

import (
	"time"
	"tracer-study-grpc/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	KabKotaTableName = "ref_kabkota"
)

type KabKota struct {
	IdWil          string         `json:"id_wil"`
	Nama           string         `json:"nama"`
	IdIndukWilayah string         `json:"id_induk_wilayah"`
	CreatedAt      time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func NewKabKota(idWil, nama, idIndukWilayah string) *KabKota {
	return &KabKota{
		IdWil:          idWil,
		Nama:           nama,
		IdIndukWilayah: idIndukWilayah,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func (k *KabKota) TableName() string {
	return KabKotaTableName
}

func ConvertEntityToProto(k *KabKota) *pb.KabKota {
	return &pb.KabKota{
		IdWil:          k.IdWil,
		Nama:           k.Nama,
		IdIndukWilayah: k.IdIndukWilayah,
		CreatedAt:      timestamppb.New(k.CreatedAt),
		UpdatedAt:      timestamppb.New(k.UpdatedAt),
	}
}
