package handler

import (
	"context"
	"net/http"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/common/errors"
	"tracer-study-grpc/modules/kabkota/entity"
	"tracer-study-grpc/modules/kabkota/service"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc/status"
)

type KabKotaHandler struct {
	pb.UnimplementedKabKotaServiceServer
	config     config.Config
	kabkotaSvc service.KabKotaServiceUseCase
}

func NewKabKotaHandler(config config.Config, kabkotaService service.KabKotaServiceUseCase) *KabKotaHandler {
	return &KabKotaHandler{
		config:     config,
		kabkotaSvc: kabkotaService,
	}
}

func (kh *KabKotaHandler) GetAllKabKota(ctx context.Context, req *pb.EmptyKabKotaRequest) (*pb.GetAllKabKotaResponse, error) {
	kabkota, err := kh.kabkotaSvc.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	var kabkotaArr []*pb.KabKota
	for _, k := range kabkota {
		kabkotaProto := entity.ConvertEntityToProto(k)
		kabkotaArr = append(kabkotaArr, kabkotaProto)
	}

	code := uint32(http.StatusOK)
	message := "get all kabkota success"

	return &pb.GetAllKabKotaResponse{
		Code:    code,
		Message: message,
		Data:    kabkotaArr,
	}, nil
}