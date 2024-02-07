package handler

import (
	"context"
	"log"
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
		if mhs == nil {
			log.Println("WARNING: [AuthHandler-Login] Mhs resource not found")
			return nil, status.Errorf(codes.NotFound, "mhs resource not found")
		}
		parseError := errors.ParseError(err)
		log.Println("ERROR: [AuthHandler-Login] Error while fetching mhs biodata:", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	if mhs == nil {
		log.Println("WARNING: [AuthHandler-Login] Mhs resource not found")
		return nil, status.Errorf(codes.NotFound, "mhs resource not found")
	}

	// generate token with role 2
	token, err := ah.jwtManager.GenerateToken(mhs.NIM, 2)
	if err != nil {
		log.Println("ERROR: [AuthHandler-Login] Error while generating token:", err)
		return nil, status.Errorf(codes.Internal, "token failed to generate: %v", err)
	}

	return &pb.LoginResponse{
		Code:    uint32(http.StatusOK),
		Message: "login success",
		Token:   token,
	}, nil
}
