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
		log.Println("ERROR: [ProvinsiHandler - GetAllProvinsi] Error while get all provinsi: ", parseError.Message)
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

func (ph *ProvinsiHandler) GetProvinsiByIdWil(ctx context.Context, req *pb.GetProvinsiByIdWilRequest) (*pb.GetProvinsiResponse, error) {
	provinsi, err := ph.provinsiSvc.FindByIdWil(ctx, req.GetIdWil())
	if err != nil {
		if provinsi == nil {
			log.Println("WARNING: [ProvinsiHandler - GetProvinsiByIdWil] Resource provinsi not found for idWil:", req.GetIdWil())
			return nil, status.Errorf(status.Code(err), "provinsi not found")
		}
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProvinsiHandler - GetProvinsiByIdWil] Internal server error:", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	provinsiProto := entity.ConvertEntityToProto(provinsi)

	return &pb.GetProvinsiResponse{
		Code:    uint32(http.StatusOK),
		Message: "get provinsi by idWil success",
		Data:    provinsiProto,
	}, nil
}

func (ph *ProvinsiHandler) CreateProvinsi(ctx context.Context, req *pb.Provinsi) (*pb.GetProvinsiResponse, error) {
	provinsi, err := ph.provinsiSvc.Create(ctx, req.GetIdWil(), req.GetNama(), req.GetUmp())

	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProvinsiHandler - CreateProvinsi] Error while create provinsi: ", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	provinsiProto := entity.ConvertEntityToProto(provinsi)

	return &pb.GetProvinsiResponse{
		Code:    uint32(http.StatusOK),
		Message: "create provinsi success",
		Data:    provinsiProto,
	}, nil
}

func (ph *ProvinsiHandler) UpdateProvinsi(ctx context.Context, req *pb.Provinsi) (*pb.GetProvinsiResponse, error) {
	provinsiDataUpdate := &entity.Provinsi{
		Nama: req.GetNama(),
		Ump:  req.GetUmp(),
	}

	provinsi, err := ph.provinsiSvc.Update(ctx, req.GetIdWil(), provinsiDataUpdate)

	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProvinsiHandler - UpdateProvinsi] Error while update provinsi: ", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	provinsiProto := entity.ConvertEntityToProto(provinsi)

	return &pb.GetProvinsiResponse{
		Code:    uint32(http.StatusOK),
		Message: "update provinsi success",
		Data:    provinsiProto,
	}, nil
}

func (ph *ProvinsiHandler) DeleteProvinsi(ctx context.Context, req *pb.GetProvinsiByIdWilRequest) (*pb.DeleteProvinsiResponse, error) {
	err := ph.provinsiSvc.Delete(ctx, req.GetIdWil())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProvinsiHandler - DeleteProvinsi] Error while delete provinsi: ", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.DeleteProvinsiResponse{
		Code:    uint32(http.StatusOK),
		Message: "delete provinsi success",
	}, nil
}
