package handler

import (
	"context"
	"log"
	"net/http"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/common/errors"
	"tracer-study-grpc/modules/mhsbiodata/entity"
	"tracer-study-grpc/modules/mhsbiodata/service"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		if apiResponse == nil {
			log.Println("WARNING: [MhsBiodataHandler - FetchMhsBiodataByNim] Resource not found: nim ", nim)
			return nil, status.Errorf(codes.NotFound, "resource not found")
		}

		parseError := errors.ParseError(err)
		log.Println("ERROR: [MhsBiodataHandler - FetchMhsBiodataByNim] Error while fetching mhs biodata:", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	var mhsBiodata = entity.ConvertEntityToProto(apiResponse)

	return &pb.MhsBiodataResponse{
		Code:    uint32(http.StatusOK),
		Message: "get mhs biodata success",
		Data:    mhsBiodata,
	}, nil
}
