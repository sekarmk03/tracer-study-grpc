package jwt

import (
	"github.com/golang-jwt/jwt"
)

const (
	prodiServiceGetAllProdi = "/tracer_study_grpc.ProdiService/GetAllProdi"
)

var ignoreMethod = []string{
	prodiServiceGetAllProdi,
}

type CustomClaims struct {
	jwt.StandardClaims
	Nim string `json:"nim"`
	Role string `json:"role"`
}