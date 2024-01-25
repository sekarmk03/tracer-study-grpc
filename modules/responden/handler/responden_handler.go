package handler

import (
	"context"
	"net/http"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/responden/entity"
	resSvc "tracer-study-grpc/modules/responden/service"
	mhsbSvc "tracer-study-grpc/modules/mhsbiodata/service"
	"tracer-study-grpc/pb"
)

type RespondenHandler struct {
	pb.UnimplementedRespondenServiceServer
	config       config.Config
	respondenSvc resSvc.RespondenServiceUseCase
	mhsbiodataSvc mhsbSvc.MhsBiodataServiceUseCase
}

func NewRespondenHandler(config config.Config, respondenService resSvc.RespondenServiceUseCase, mhsbiodataService mhsbSvc.MhsBiodataServiceUseCase) *RespondenHandler {
	return &RespondenHandler{
		config:       config,
		respondenSvc: respondenService,
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

func (rh *RespondenHandler) UpdateRespondenFromSiak(ctx context.Context, req *pb.UpdateRespondenFromSiakRequest) (*pb.UpdateRespondenFromSiakResponse, error) {
	mhsbiodata, err := rh.mhsbiodataSvc.FetchMhsBiodataByNimFromSiakApi(req.GetNim())
	if err != nil {
		return nil, err
	}

	responden, err := rh.respondenSvc.Update(ctx, req.GetNim(), mhsbiodata)
	if err != nil {
		return nil, err
	}

	respondenProto := entity.ConvertEntityToProto(responden)

	return &pb.UpdateRespondenFromSiakResponse{
		Code:    uint32(http.StatusOK),
		Message: "update responden from siak success",
		Data:    respondenProto,
	}, nil
}
