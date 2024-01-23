package handler

import (
	"context"
	"net/http"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/common/errors"
	"tracer-study-grpc/modules/provinsi/entity"
	"tracer-study-grpc/modules/provinsi/service"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc/status"
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

func (ph *ProvinsiHandler) GetAllProvinsi(ctx context.Context, req *pb.EmptyProvinsiRequest) (*pb.GetAllProvinsiResponse, error) {
	provinsi, err := ph.provinsiSvc.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	var provinsiArr []*pb.Provinsi
	for _, p := range provinsi {
		provinsiProto := entity.ConvertEntityToProto(p)
		provinsiArr = append(provinsiArr, provinsiProto)
	}

	code := uint32(http.StatusOK)
	message := "get all provinsi success"

	return &pb.GetAllProvinsiResponse{
		Code:    code,
		Message: message,
		Data:    provinsiArr,
	}, nil
}
