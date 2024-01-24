package handler

import (
	"context"
	"net/http"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/responden/entity"
	"tracer-study-grpc/modules/responden/service"
	"tracer-study-grpc/pb"
)

type RespondenHandler struct {
	pb.UnimplementedRespondenServiceServer
	config config.Config
	respondenSvc service.RespondenServiceUseCase
}

func NewRespondenHandler(config config.Config, respondenService service.RespondenServiceUseCase) *RespondenHandler {
	return &RespondenHandler{
		config: config,
		respondenSvc: respondenService,
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