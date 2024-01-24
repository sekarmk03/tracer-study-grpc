package entity

import "tracer-study-grpc/pb"

type MhsBiodata struct {
	NIM      string `json:"nim"`
	NAMA     string `json:"nama"`
	IPK      string `json:"ipk"`
	HP       string `json:"hp"`
	EMAIL    string `json:"email"`
	KODEPSTD string `json:"kodepstd"`
	JENJANG  string `json:"jenjang"`
	PRODI    string `json:"prodi"`
	NAMAPST  string `json:"namapst"`
	KODEFAK  string `json:"kodefak"`
	NAMAFAK  string `json:"namafak"`
	JLRMASUK string `json:"jlrmasuk"`
	THNMASUK string `json:"thnmasuk"`
	LAMASTD  string `json:"lamastd"`
}

func NewMhsBiodata(nim, nama, ipk, hp, email, kodepstd, jenjang, prodi, namapst, kodefak, namafak, jlrmasuk, thnmasuk, lamastd string) *MhsBiodata {
	return &MhsBiodata{
		NIM:      nim,
		NAMA:     nama,
		IPK:      ipk,
		HP:       hp,
		EMAIL:    email,
		KODEPSTD: kodepstd,
		JENJANG:  jenjang,
		PRODI:    prodi,
		NAMAPST:  namapst,
		KODEFAK:  kodefak,
		NAMAFAK:  namafak,
		JLRMASUK: jlrmasuk,
		THNMASUK: thnmasuk,
		LAMASTD:  lamastd,
	}
}

func ConvertEntityToProto(m *MhsBiodata) *pb.MhsBiodata {
	return &pb.MhsBiodata{
		NIM:      m.NIM,
		NAMA:     m.NAMA,
		IPK:      m.IPK,
		HP:       m.HP,
		EMAIL:    m.EMAIL,
		KODEPSTD: m.KODEPSTD,
		JENJANG:  m.JENJANG,
		PRODI:    m.PRODI,
		NAMAPST:  m.NAMAPST,
		KODEFAK:  m.KODEFAK,
		JLRMASUK: m.JLRMASUK,
		THNMASUK: m.THNMASUK,
		LAMASTD:  m.LAMASTD,
	}
}
