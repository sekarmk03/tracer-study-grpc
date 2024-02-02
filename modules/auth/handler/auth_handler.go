package handler

import (
	"context"
	"net/http"
	"tracer-study-grpc/common/config"
	commonJwt "tracer-study-grpc/common/jwt"
	"tracer-study-grpc/modules/mhsbiodata/service"
	"tracer-study-grpc/pb"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	config     config.Config
	mhsSvc     service.MhsBiodataServiceUseCase
	jwtManager *commonJwt.JWT
}

func NewAuthHandler(config config.Config, mhsService service.MhsBiodataServiceUseCase, jwtManager *commonJwt.JWT) *AuthHandler {
	return &AuthHandler{
		config:     config,
		mhsSvc:     mhsService,
		jwtManager: jwtManager,
	}
}

func (ah *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	mhs, err := ah.mhsSvc.FetchMhsBiodataByNimFromSiakApi(req.GetNim())
	if err != nil {
		return nil, err
	}

	if mhs == nil {
		return &pb.LoginResponse{
			Code:    uint32(http.StatusNotFound),
			Message: "mahasiswa tidak ditemukan",
			Token:   "",
		}, nil
	}

	token, err := ah.jwtManager.GenerateToken(mhs.NIM, 2)
	if err != nil {
		return &pb.LoginResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: "token failed to generate",
			Token:   "",
		}, nil
	}

	return &pb.LoginResponse{
		Code:    uint32(http.StatusOK),
		Message: "login success",
		Token:   token,
	}, nil
}
