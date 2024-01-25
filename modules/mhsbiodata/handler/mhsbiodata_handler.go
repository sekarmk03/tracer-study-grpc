package handler

import (
	"context"
	"net/http"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/mhsbiodata/entity"
	"tracer-study-grpc/modules/mhsbiodata/service"
	"tracer-study-grpc/pb"
)

type MhsBiodataHandler struct {
	pb.UnimplementedMhsBiodataServiceServer
	config        config.Config
	mhsbiodataSvc service.MhsBiodataServiceUseCase
}

func NewMhsBiodataHandler(config config.Config, mhsbiodataService service.MhsBiodataServiceUseCase) *MhsBiodataHandler {
	return &MhsBiodataHandler{
		config:        config,
		mhsbiodataSvc: mhsbiodataService,
	}
}

func (mbh *MhsBiodataHandler) FetchMhsBiodataByNim(ctx context.Context, req *pb.MhsBiodataRequest) (*pb.MhsBiodataResponse, error) {
	nim := req.GetNim()

	var apiResponse *entity.MhsBiodata
	apiResponse, err := mbh.mhsbiodataSvc.FetchMhsBiodataByNimFromSiakApi(nim)
	if err != nil {
		return nil, err
	}

	var mhsBiodata = entity.ConvertEntityToProto(apiResponse)

	code := uint32(http.StatusOK)
	message := "get mhs biodata success"

	return &pb.MhsBiodataResponse{
		Code:    code,
		Message: message,
		Data:    mhsBiodata,
	}, nil
}
