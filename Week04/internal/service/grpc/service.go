package grpcservice

import (
	"context"
	pb "github.com/KelvinChen684/Go-000/tree/main/Week04/api/v1"
	"github.com/KelvinChen684/Go-000/tree/main/Week04/internal/biz"
)

type UserService struct {
	us *biz.UserUsecase
	pb.UnimplementedUserServer
}

func NewUserService(us *biz.UserUsecase) *UserService {
	return &UserService{us:us}
	/*
	lis, _ := net.Listen("tcp", ":8000")
	s := grpc.NewServer()
	pb.RegisterUserServer(s, userServ)
	go s.Serve(lis)
	 */
}

func (userSrv *UserService) GetUserInfo(ctx context.Context, in *pb.UserReq) (out *pb.UserInfo, err error)  {
	// DTO-->DO
	user := &biz.User{Name: in.Name}
	// 模拟多个biz业务组装数据
	var userGender *biz.UserGender
	var userAge    *biz.UserAge

	if userGender, err = userSrv.us.GetUserGender(user); err != nil {
		return nil, err
	}
	if userAge, err = userSrv.us.GetUserAge(user); err != nil {
		return nil, err
	}

	out = &pb.UserInfo{
		Name:   in.Name,
		Gender: userGender.Gender,
		Age:    userAge.Age,
	}
	return
}
