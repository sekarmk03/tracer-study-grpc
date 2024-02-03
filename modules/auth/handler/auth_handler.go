package handler

import (
	"context"
	"net/http"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/common/errors"
	commonJwt "tracer-study-grpc/common/jwt"

	"tracer-study-grpc/modules/mhsbiodata/service"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		parseError := errors.ParseError(err)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	if mhs == nil {
		return nil, status.Errorf(codes.NotFound, "mhs resource not found")
	}

	token, err := ah.jwtManager.GenerateToken(mhs.NIM, 2)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "token failed to generate")
	}

	return &pb.LoginResponse{
		Code:    uint32(http.StatusOK),
		Message: "login success",
		Token:   token,
	}, nil
}
