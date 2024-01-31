package handler

import (
	"context"
	"net/http"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/pkts/entity"
	"tracer-study-grpc/modules/pkts/service"
	"tracer-study-grpc/pb"
)

type PKTSHandler struct {
	pb.UnimplementedPKTSServiceServer
	config  config.Config
	PKTSSvc service.PKTSServiceUseCase
}

func NewPKTSHandler(config config.Config, pktsService service.PKTSServiceUseCase) *PKTSHandler {
	return &PKTSHandler{
		config:  config,
		PKTSSvc: pktsService,
	}
}

func (ph *PKTSHandler) GetAllPKTS(ctx context.Context, req *pb.EmptyPKTSRequest) (*pb.GetAllPKTSResponse, error) {
	pkts, err := ph.PKTSSvc.FindAll(ctx, req)
	if err != nil {
		return nil, err
	}

	var pktsArr []*pb.PKTS
	for _, p := range pkts {
		pktsProto := entity.ConvertEntityToProto(p)
		pktsArr = append(pktsArr, pktsProto)
	}

	return &pb.GetAllPKTSResponse{
		Code:    uint32(http.StatusOK),
		Message: "get all pkts success",
		Data:    pktsArr,
	}, nil
}