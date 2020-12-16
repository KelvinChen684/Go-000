package main

import (
	"github.com/KelvinChen684/Go-000/tree/main/Week04/internal/biz"
	"github.com/KelvinChen684/Go-000/tree/main/Week04/internal/data"
	grpcservice "github.com/KelvinChen684/Go-000/tree/main/Week04/internal/service/grpc"
	"github.com/google/wire"
)

//go:generate wire
func InitializeUserService() (*grpcservice.UserService, error){
	wire.Build(grpcservice.NewUserService, biz.NewUserUsecase, data.NewData, data.NewDB, data.NewRedis)
	return &grpcservice.UserService{}, nil
}