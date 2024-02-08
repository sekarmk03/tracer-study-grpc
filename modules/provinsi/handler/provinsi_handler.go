package handler

import (
	"context"
	"log"
	"net/http"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/common/errors"
	"tracer-study-grpc/modules/provinsi/entity"
	"tracer-study-grpc/modules/provinsi/service"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProvinsiHandler struct {
	pb.UnimplementedProvinsiServiceServer
	config      config.Config
	provinsiSvc service.ProvinsiServiceUseCase
}

func NewProvinsiHandler(config config.Config, provinsiService service.ProvinsiServiceUseCase) *ProvinsiHandler {
	return &ProvinsiHandler{
		config:      config,
		provinsiSvc: provinsiService,
	}
}

func (ph *ProvinsiHandler) GetAllProvinsi(ctx context.Context, req *emptypb.Empty) (*pb.GetAllProvinsiResponse, error) {
	provinsi, err := ph.provinsiSvc.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProvinsiHandler-GetAllProvinsi] Error while get all provinsi: ", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	var provinsiArr []*pb.Provinsi
	for _, p := range provinsi {
		provinsiProto := entity.ConvertEntityToProto(p)
		provinsiArr = append(provinsiArr, provinsiProto)
	}

	return &pb.GetAllProvinsiResponse{
		Code:    uint32(http.StatusOK),
		Message: "get all provinsi success",
		Data:    provinsiArr,
	}, nil
}
