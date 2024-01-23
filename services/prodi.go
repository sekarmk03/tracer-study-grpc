package services

import (
	"tracer-study-grpc/pb"

	"gorm.io/gorm"
)

type ProdiService struct {
	pb.UnimplementedProdiServiceServer
	db *gorm.DB
}

func NewProdiService(db *gorm.DB) *ProdiService {
	return &ProdiService{
		pb.UnimplementedProdiServiceServer{},
		db,
	}
}