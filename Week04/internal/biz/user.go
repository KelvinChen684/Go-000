package biz

import (
	"errors"
	"math/rand"
)

//var ServiceProvider = wire.Build()

type UserData interface {
	Query(user *User) (*UserGender, error)
}

type UserGender struct {
	Name string
	Gender string
}

type UserAge struct {
	Name string
	Age int32
}

type UserInfo struct {
	Name   string
	Gender string
	Age    int
	//Passwd string
}

type User struct {
	Name string
}

type UserUsecase struct {
	data UserData
}

func NewUserUsecase(data UserData) *UserUsecase {
	return &UserUsecase{data:data}
}

func (us *UserUsecase) GetUserGender(user *User) (ug *UserGender, err error) {
	// 模拟一个业务逻辑
	// 错误处理？？？
	if user.Name == "abc" {
		return nil, errors.New(user.Name + "don't have the previlige of getting user gender")
	}

	if ug, err = us.data.Query(user); err != nil {
		return nil, errors.New("no user " + user.Name)
	}
	return us.data.Query(user)
}

func (us *UserUsecase) GetUserAge(user *User) (*UserAge, error) {
	// 模拟数据来自其他服务
	// age := otherGrpcService(user)
	age := &UserAge{
		Name: user.Name,
		Age:  rand.Int31n(80),
	}
	return age, nil
}


