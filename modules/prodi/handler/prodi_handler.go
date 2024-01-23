package handler

import (
	"context"
	"net/http"
	"tracer-study-grpc/common/config"

	"tracer-study-grpc/common/errors"
	"tracer-study-grpc/modules/prodi/entity"
	"tracer-study-grpc/modules/prodi/service"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc/status"
)

type ProdiHandler struct {
	pb.UnimplementedProdiServiceServer
	config   config.Config
	prodiSvc service.ProdiServiceUseCase
}

func NewProdiHandler(config config.Config, prodiService service.ProdiServiceUseCase) *ProdiHandler {
	return &ProdiHandler{
		config:   config,
		prodiSvc: prodiService,
	}
}

func (ph *ProdiHandler) GetAllProdi(ctx context.Context, req *pb.EmptyProdiRequest) (*pb.GetAllProdiResponse, error) {
	prodi, err := ph.prodiSvc.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	var prodiArr []*pb.Prodi
	for _, p := range prodi {
		// convert each p to pb.Prodi
		prodiProto := entity.ConvertEntityToProto(p)
		prodiArr = append(prodiArr, prodiProto)
	}

	code := uint32(http.StatusOK)
	message := "get all prodi success"

	return &pb.GetAllProdiResponse{
		Code:    code,
		Message: message,
		Data:    prodiArr,
	}, nil
}
