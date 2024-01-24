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
	ThnAngkatan  string         `json:"thn_ak"`
	Semester     uint32         `json:"semester"`
	Type         string         `json:"type"`
	StatusUpdate string         `json:"status_update"`
	JalurMasuk   string         `json:"jlrmasuk"`
	TahunMasuk   string         `json:"thnmasuk"`
	LamaStudi    string         `json:"lamastd"`
	KodeFak      string         `json:"kodefak"`
	NamaFak      string         `json:"namafak"`
	Nim          string         `json:"nim"`
	Nama         string         `json:"nama"`
	KodeProdi    string         `json:"kodeprodi"`
	KodeProdi2   string         `json:"kodeprodi2"`
	NamaProdi    string         `json:"namaprodi"`
	NamaProdi2   string         `json:"namaprodi2"`
	KodeDikti    string         `json:"kodedikti"`
	Jenjang      string         `json:"jenjang"`
	JenisKel     string         `json:"jk"`
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
		ThnAngkatan:  thnAk,
		Semester:     semester,
		Type:         tipe,
		StatusUpdate: statusUpdate,
		JalurMasuk:   jlrMasuk,
		TahunMasuk:   thnMasuk,
		LamaStudi:    lamaStudi,
		KodeFak:      kodeFak,
		NamaFak:      namaFak,
		Nim:          nim,
		Nama:         nama,
		KodeProdi:    kodeProdi,
		KodeProdi2:   kodeProdi2,
		NamaProdi:    namaProdi2,
		KodeDikti:    kodeDikti,
		Jenjang:      jenjang,
		JenisKel:     jenisKel,
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
		Id:    r.ID,
		ThnAk: r.ThnAngkatan,
		Semester: r.Semester,
		Type: r.Type,
		StatusUpdate: r.StatusUpdate,
		Jlrmasuk: r.JalurMasuk,
		Thnmasuk: r.TahunMasuk,
		Lamastd: r.LamaStudi,
		Kodefak: r.KodeFak,
		Namafak: r.NamaFak,
		Nim: r.Nim,
		Nama: r.Nama,
		Kodeprodi: r.KodeProdi,
		Kodeprodi2: r.KodeProdi2,
		Namaprodi: r.NamaProdi,
		Namaprodi2: r.NamaProdi2,
		Kodedikti: r.KodeDikti,
		Jenjang: r.Jenjang,
		Jk: r.JenisKel,
		Email: r.Email,
		Hp: r.Hp,
		TglSidang: r.TglSidang,
		ThnSidang: r.ThnSidang,
		TglWisuda: r.TglWisuda,
		Nik: r.Nik,
		Npwp: r.Npwp,
		CreatedBy: r.CreatedBy,
		UpdatedBy: r.UpdatedBy,
		CreatedAt: timestamppb.New(r.CreatedAt),
		UpdatedAt: timestamppb.New(r.UpdatedAt),
	}
}
