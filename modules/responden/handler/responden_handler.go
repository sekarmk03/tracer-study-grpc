package handler

import (
	"context"
	"net/http"
	"tracer-study-grpc/common/config"
	mhsbSvc "tracer-study-grpc/modules/mhsbiodata/service"
	"tracer-study-grpc/modules/responden/entity"
	resSvc "tracer-study-grpc/modules/responden/service"
	"tracer-study-grpc/pb"
)

type RespondenHandler struct {
	pb.UnimplementedRespondenServiceServer
	config        config.Config
	respondenSvc  resSvc.RespondenServiceUseCase
	mhsbiodataSvc mhsbSvc.MhsBiodataServiceUseCase
}

func NewRespondenHandler(config config.Config, respondenService resSvc.RespondenServiceUseCase, mhsbiodataService mhsbSvc.MhsBiodataServiceUseCase) *RespondenHandler {
	return &RespondenHandler{
		config:        config,
		respondenSvc:  respondenService,
		mhsbiodataSvc: mhsbiodataService,
	}
}

func (rh *RespondenHandler) GetAllResponden(ctx context.Context, req *pb.EmptyRespondenRequest) (*pb.GetAllRespondenResponse, error) {
	responden, err := rh.respondenSvc.FindAll(ctx, req)
	if err != nil {
		return nil, err
	}

	var respondenArr []*pb.Responden
	for _, r := range responden {
		respondenProto := entity.ConvertEntityToProto(r)
		respondenArr = append(respondenArr, respondenProto)
	}

	return &pb.GetAllRespondenResponse{
		Code:    uint32(http.StatusOK),
		Message: "get all responden success",
		Data:    respondenArr,
	}, nil
}

func (rh *RespondenHandler) GetRespondenByNim(ctx context.Context, req *pb.GetRespondenByNimRequest) (*pb.GetRespondenByNimResponse, error) {
	responden, err := rh.respondenSvc.FindByNim(ctx, req.GetNim())
	if err != nil {
		return nil, err
	}

	respondenProto := entity.ConvertEntityToProto(responden)

	return &pb.GetRespondenByNimResponse{
		Code:    uint32(http.StatusOK),
		Message: "get responden by nim success",
		Data:    respondenProto,
	}, nil
}

func (rh *RespondenHandler) UpdateRespondenFromSiak(ctx context.Context, req *pb.UpdateRespondenFromSiakRequest) (*pb.UpdateRespondenResponse, error) {
	mhsbiodata, err := rh.mhsbiodataSvc.FetchMhsBiodataByNimFromSiakApi(req.GetNim())
	if err != nil {
		return nil, err
	}

	convertEntity := &entity.Responden{
		Ipk:        mhsbiodata.IPK,
		Kodedikti:  mhsbiodata.KODEPSTD,
		Jenjang:    mhsbiodata.JENJANG,
		Namaprodi:  mhsbiodata.PRODI,
		Namaprodi2: mhsbiodata.NAMAPST,
		Kodeprodi:  mhsbiodata.KODEPST,
		Kodeprodi2: mhsbiodata.KODEPST[:4],
		Kodefak:    mhsbiodata.KODEFAK,
		Namafak:    mhsbiodata.NAMAFAK,
		Jlrmasuk:   mhsbiodata.JLRMASUK,
		Thnmasuk:   mhsbiodata.THNMASUK,
		ThnAk:      mhsbiodata.THNMASUK,
		Lamastd:    mhsbiodata.LAMASTD,
	}

	responden, err := rh.respondenSvc.Update(ctx, req.GetNim(), convertEntity)
	if err != nil {
		return nil, err
	}

	respondenProto := entity.ConvertEntityToProto(responden)

	return &pb.UpdateRespondenResponse{
		Code:    uint32(http.StatusOK),
		Message: "update responden from siak success",
		Data:    respondenProto,
	}, nil
}

func (rh *RespondenHandler) CreateResponden(ctx context.Context, req *pb.CreateRespondenRequest) (*pb.CreateRespondenResponse, error) {
	responden, err := rh.respondenSvc.Create(ctx, req.GetNim(), req.GetSemester(), req.GetType(), req.GetNama(), req.GetKodeprodi(), req.GetJk(), req.GetTglWisuda(), req.GetTglSidang(), req.GetThnSidang())

	if err != nil {
		return nil, err
	}

	respondenProto := entity.ConvertEntityToProto(responden)

	return &pb.CreateRespondenResponse{
		Code:    uint32(http.StatusOK),
		Message: "create responden success",
		Data:    respondenProto,
	}, nil
}

func (rh *RespondenHandler) UpdateResponden(ctx context.Context, req *pb.Responden) (*pb.UpdateRespondenResponse, error) {
	respDataUpdate := &entity.Responden{
		Email: req.GetEmail(),
		Hp:    req.GetHp(),
		Nik:   req.GetNik(),
		Npwp:  req.GetNpwp(),
	}

	responden, err := rh.respondenSvc.Update(ctx, req.GetNim(), respDataUpdate)
	if err != nil {
		return nil, err
	}

	respondenProto := entity.ConvertEntityToProto(responden)

	return &pb.UpdateRespondenResponse{
		Code:    uint32(http.StatusOK),
		Message: "update responden success",
		Data:    respondenProto,
	}, nil
}
