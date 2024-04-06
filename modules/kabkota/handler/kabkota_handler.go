package handler

import (
	"context"
	"log"
	"net/http"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/common/errors"
	"tracer-study-grpc/modules/kabkota/entity"
	"tracer-study-grpc/modules/kabkota/service"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (kh *KabKotaHandler) GetAllKabKota(ctx context.Context, req *emptypb.Empty) (*pb.GetAllKabKotaResponse, error) {
	kabkota, err := kh.kabkotaSvc.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [KabKotaHandler-GetAllKabKota] Error:", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	var kabkotaArr []*pb.KabKota
	for _, k := range kabkota {
		kabkotaProto := entity.ConvertEntityToProto(k)
		kabkotaArr = append(kabkotaArr, kabkotaProto)
	}

	return &pb.GetAllKabKotaResponse{
		Code:    uint32(http.StatusOK),
		Message: "get all kabkota success",
		Data:    kabkotaArr,
	}, nil
}

func (kh *KabKotaHandler) GetKabKotaByIdWil(ctx context.Context, req *pb.GetKabKotaByIdWilRequest) (*pb.GetKabKotaByIdWilResponse, error) {
	kabkota, err := kh.kabkotaSvc.FindByIdWil(ctx, req.IdWil)
	if err != nil {
		if kabkota == nil {
			log.Println("WARNING: [KabKotaHandler-GetKabKotaByIdWil] Resource kabkota not found for idWil:", req.IdWil)
			return nil, status.Errorf(codes.NotFound, "kabkota not found")
		}
		parseError := errors.ParseError(err)
		log.Println("ERROR: [KabKotaHandler-GetKabKotaByIdWil] Internal server error:", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	kabkotaProto := entity.ConvertEntityToProto(kabkota)

	return &pb.GetKabKotaByIdWilResponse{
		Code:   uint32(http.StatusOK),
		Message: "get kabkota by idWil success",
		Data:   kabkotaProto,
	}, nil
}