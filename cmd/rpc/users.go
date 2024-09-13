package main

import (
	"context"
	"errors"
	"github.com/agung96tm/go-simple-web-application/internal/data"
	"github.com/agung96tm/go-simple-web-application/internal/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	proto.UnimplementedUserServiceServer
	app *Application
}

func NewUserService(app *Application) *UserService {
	return &UserService{
		app: app,
	}
}

func (u UserService) List(_ context.Context, req *proto.UserListRequest) (*proto.UserListResponse, error) {
	users, err := u.app.Models.UserModel.GetAll()
	if err != nil {
		return nil, err
	}

	var usersResp []*proto.User
	for _, u := range *users {
		usersResp = append(usersResp, &proto.User{
			Id:   u.ID,
			Name: u.Name,
		})
	}

	return &proto.UserListResponse{
		Data: usersResp,
	}, nil
}

func (u UserService) Detail(_ context.Context, req *proto.UserDetailRequest) (*proto.UserDetailResponse, error) {
	user, err := u.app.Models.UserModel.GetById(req.Id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &proto.UserDetailResponse{
		User: &proto.User{
			Id:   user.ID,
			Name: user.Name,
		},
	}, nil
}

func (u UserService) Create(_ context.Context, req *proto.UserCreateRequest) (*proto.UserDetailResponse, error) {
	user := &data.User{
		Name: req.User.Name,
	}

	err := u.app.Models.UserModel.Create(user)
	if err != nil {
		return nil, err
	}

	return &proto.UserDetailResponse{
		User: &proto.User{
			Id:   user.ID,
			Name: user.Name,
		},
	}, nil
}

func (u UserService) Update(_ context.Context, req *proto.UserUpdateRequest) (*proto.UserDetailResponse, error) {
	user, err := u.app.Models.UserModel.GetById(req.User.Id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if req.User.Name != user.Name {
		user.Name = req.User.Name
	}
	err = u.app.Models.UserModel.Update(user)
	if err != nil {
		return nil, err
	}

	return &proto.UserDetailResponse{
		User: &proto.User{
			Id:   user.ID,
			Name: user.Name,
		},
	}, nil
}

func (u UserService) Delete(_ context.Context, req *proto.UserDeleteRequest) (*emptypb.Empty, error) {
	user, err := u.app.Models.UserModel.GetById(req.Id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = u.app.Models.UserModel.Delete(user.ID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
