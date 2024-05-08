package service

import (
	"context"
	"log"
	"time"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/common/errors"
	"tracer-study-grpc/common/utils"
	"tracer-study-grpc/modules/user/entity"
	"tracer-study-grpc/modules/user/repository"
)

type UserService struct {
	cfg            config.Config
	userRepository repository.UserRepositoryUseCase
}

type UserServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.User, error)
	FindById(ctx context.Context, id uint64) (*entity.User, error)
	Create(ctx context.Context, name, username, email, password string, roleId uint32) (*entity.User, error)
	Update(ctx context.Context, id uint64, fields *entity.User) (*entity.User, error)
	Delete(ctx context.Context, id uint64) error
}

func NewUserService(cfg config.Config, userRepository repository.UserRepositoryUseCase) *UserService {
	return &UserService{
		cfg:            cfg,
		userRepository: userRepository,
	}
}

func (svc *UserService) FindAll(ctx context.Context, req any) ([]*entity.User, error) {
	res, err := svc.userRepository.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserService - FindAll] Error while find all user: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *UserService) FindById(ctx context.Context, id uint64) (*entity.User, error) {
	res, err := svc.userRepository.FindById(ctx, id)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserService - FindById] Error while find user by ID: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *UserService) Create(ctx context.Context, name, username, email, password string, roleId uint32) (*entity.User, error) {
	user := &entity.User{
		Name:      name,
		Username:  username,
		Email:     email,
		Password:  utils.HashPassword(password),
		RoleId:    roleId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res, err := svc.userRepository.Create(ctx, user)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserService - Create] Error while create user: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *UserService) Update(ctx context.Context, id uint64, fields *entity.User) (*entity.User, error) {
	user, err := svc.userRepository.FindById(ctx, id)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserService - Update] Error while find user by ID: ", parseError.Message)
		return nil, err
	}

	updatedMap := make(map[string]interface{})

	utils.AddItemToMap(updatedMap, "name", fields.Name)
	utils.AddItemToMap(updatedMap, "username", fields.Username)
	utils.AddItemToMap(updatedMap, "email", fields.Email)
	utils.AddItemToMap(updatedMap, "role_id", fields.RoleId)

	res, err := svc.userRepository.Update(ctx, user, updatedMap)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserService - Update] Error while update user: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *UserService) Delete(ctx context.Context, id uint64) error {
	err := svc.userRepository.Delete(ctx, id)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserService - Delete] Error while delete user: ", parseError.Message)
		return err
	}

	return nil
}
