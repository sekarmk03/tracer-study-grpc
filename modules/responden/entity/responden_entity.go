package entity

import (
	"time"
	"tracer-study-grpc/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	RespondenTableName = "responden"
)

type Responden struct {
	ID           uint32         `json:"id"`
	ThnAk        string         `json:"thn_ak"`
	Semester     uint32         `json:"semester"`
	Type         string         `json:"type"`
	StatusUpdate string         `json:"status_update"`
	Jlrmasuk     string         `json:"jlrmasuk"`
	Thnmasuk     string         `json:"thnmasuk"`
	Lamastd      string         `json:"lamastd"`
	Kodefak      string         `json:"kodefak"`
	Namafak      string         `json:"namafak"`
	Nim          string         `json:"nim"`
	Nama         string         `json:"nama"`
	Kodeprodi    string         `json:"kodeprodi"`
	Kodeprodi2   string         `json:"kodeprodi2"`
	Namaprodi    string         `json:"namaprodi"`
	Namaprodi2   string         `json:"namaprodi2"`
	Kodedikti    string         `json:"kodedikti"`
	Jenjang      string         `json:"jenjang"`
	JK           string         `json:"jk"`
	Email        string         `json:"email"`
	Hp           string         `json:"hp"`
	Ipk          string         `json:"ipk"`
	TglSidang    string         `json:"tgl_sidang"`
	ThnSidang    string         `json:"thn_sidang"`
	TglWisuda    string         `json:"tgl_wisuda"`
	Nik          string         `json:"nik"`
	Npwp         string         `json:"npwp"`
	CreatedBy    string         `json:"created_by"`
	UpdatedBy    string         `json:"updated_by"`
	CreatedAt    time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func NewResponden(id uint32, thnAk string, semester uint32, tipe, statusUpdate, jlrMasuk, thnMasuk, lamaStudi, kodeFak, namaFak, nim, nama, kodeProdi, kodeProdi2, namaProdi, namaProdi2, kodeDikti, jenjang, jenisKel, email, hp, ipk, tglSidang, thnSidang, tglWisuda, nik, npwp, createdBy, updatedBy string) *Responden {
	return &Responden{
		ID:           id,
		ThnAk:        thnAk,
		Semester:     semester,
		Type:         tipe,
		StatusUpdate: statusUpdate,
		Jlrmasuk:     jlrMasuk,
		Thnmasuk:     thnMasuk,
		Lamastd:      lamaStudi,
		Kodefak:      kodeFak,
		Namafak:      namaFak,
		Nim:          nim,
		Nama:         nama,
		Kodeprodi:    kodeProdi,
		Kodeprodi2:   kodeProdi2,
		Namaprodi:    namaProdi2,
		Kodedikti:    kodeDikti,
		Jenjang:      jenjang,
		JK:           jenisKel,
		Email:        email,
		Hp:           hp,
		Ipk:          ipk,
		TglSidang:    tglSidang,
		ThnSidang:    thnSidang,
		TglWisuda:    tglWisuda,
		Nik:          nik,
		Npwp:         npwp,
		CreatedBy:    createdBy,
		UpdatedBy:    updatedBy,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (r *Responden) TableName() string {
	return RespondenTableName
}

func ConvertEntityToProto(r *Responden) *pb.Responden {
	return &pb.Responden{
		Id:           r.ID,
		ThnAk:        r.ThnAk,
		Semester:     r.Semester,
		Type:         r.Type,
		StatusUpdate: r.StatusUpdate,
		Jlrmasuk:     r.Jlrmasuk,
		Thnmasuk:     r.Thnmasuk,
		Lamastd:      r.Lamastd,
		Kodefak:      r.Kodefak,
		Namafak:      r.Namafak,
		Nim:          r.Nim,
		Nama:         r.Nama,
		Kodeprodi:    r.Kodeprodi,
		Kodeprodi2:   r.Kodeprodi2,
		Namaprodi:    r.Namaprodi,
		Namaprodi2:   r.Namaprodi2,
		Kodedikti:    r.Kodedikti,
		Jenjang:      r.Jenjang,
		Jk:           r.JK,
		Email:        r.Email,
		Hp:           r.Hp,
		Ipk:          r.Ipk,
		TglSidang:    r.TglSidang,
		ThnSidang:    r.ThnSidang,
		TglWisuda:    r.TglWisuda,
		Nik:          r.Nik,
		Npwp:         r.Npwp,
		CreatedBy:    r.CreatedBy,
		UpdatedBy:    r.UpdatedBy,
		CreatedAt:    timestamppb.New(r.CreatedAt),
		UpdatedAt:    timestamppb.New(r.UpdatedAt),
	}
}
