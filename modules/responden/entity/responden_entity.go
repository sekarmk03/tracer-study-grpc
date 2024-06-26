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
	Id            uint32         `json:"id"`
	Nim           string         `json:"nim"`
	Nama          string         `json:"nama"`
	JalurMasuk    string         `json:"jalur_masuk"`
	TahunMasuk    string         `json:"tahun_masuk"`
	LamaStudi     uint32         `json:"lama_studi"`
	KodeFakultas  string         `json:"kode_fakultas"`
	KodeProdi     string         `json:"kode_prodi"`
	JenisKelamin  string         `json:"jenis_kelamin"`
	Email         string         `json:"email"`
	Hp            string         `json:"hp"`
	Ipk           string         `json:"ipk"`
	TanggalSidang string         `json:"tanggal_sidang"`
	TahunSidang   string         `json:"tahun_sidang"`
	TanggalWisuda string         `json:"tanggal_wisuda"`
	Nik           string         `json:"nik"`
	Npwp          string         `json:"npwp"`
	CreatedAt     time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func NewResponden(id uint32, nim, nama, jlrMasuk, thnMasuk string, lamaStudi uint32, kodeFak, kodeProdi, jenisKel, email, hp, ipk, tglSidang, thnSidang, tglWisuda, nik, npwp string) *Responden {
	return &Responden{
		Id:            id,
		Nim:           nim,
		Nama:          nama,
		JalurMasuk:    jlrMasuk,
		TahunMasuk:    thnMasuk,
		LamaStudi:     lamaStudi,
		KodeFakultas:  kodeFak,
		KodeProdi:     kodeProdi,
		JenisKelamin:  jenisKel,
		Email:         email,
		Hp:            hp,
		Ipk:           ipk,
		TanggalSidang: tglSidang,
		TahunSidang:   thnSidang,
		TanggalWisuda: tglWisuda,
		Nik:           nik,
		Npwp:          npwp,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

func (r *Responden) TableName() string {
	return RespondenTableName
}

func ConvertEntityToProto(r *Responden) *pb.Responden {
	return &pb.Responden{
		Id:            r.Id,
		Nim:           r.Nim,
		Nama:          r.Nama,
		JalurMasuk:    r.JalurMasuk,
		TahunMasuk:    r.TahunMasuk,
		LamaStudi:     r.LamaStudi,
		KodeFakultas:  r.KodeFakultas,
		KodeProdi:     r.KodeProdi,
		JenisKelamin:  r.JenisKelamin,
		Email:         r.Email,
		Hp:            r.Hp,
		Ipk:           r.Ipk,
		TanggalSidang: r.TanggalSidang,
		TahunSidang:   r.TahunSidang,
		TanggalWisuda: r.TanggalWisuda,
		Nik:           r.Nik,
		Npwp:          r.Npwp,
		CreatedAt:     timestamppb.New(r.CreatedAt),
		UpdatedAt:     timestamppb.New(r.UpdatedAt),
	}
}

func ConvertProtoToEntity(r *pb.Responden) *Responden {
	return &Responden{
		Id:            r.GetId(),
		Nim:           r.GetNim(),
		Nama:          r.GetNama(),
		JalurMasuk:    r.GetJalurMasuk(),
		TahunMasuk:    r.GetTahunMasuk(),
		LamaStudi:     r.GetLamaStudi(),
		KodeFakultas:  r.GetKodeFakultas(),
		KodeProdi:     r.GetKodeProdi(),
		JenisKelamin:  r.GetJenisKelamin(),
		Email:         r.GetEmail(),
		Hp:            r.GetHp(),
		Ipk:           r.GetIpk(),
		TanggalSidang: r.GetTanggalSidang(),
		TahunSidang:   r.GetTahunSidang(),
		TanggalWisuda: r.GetTanggalWisuda(),
		Nik:           r.GetNik(),
		Npwp:          r.GetNpwp(),
		CreatedAt:     r.GetCreatedAt().AsTime(),
		UpdatedAt:     r.GetUpdatedAt().AsTime(),
	}
}
