package api

import (
	"context"
	"log_user/dao"
	"log_user/service"
	"log_user/utils/errmsg"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user dao.User
	resp = new(service.UserDetailResponse)
	resp.Code = errmsg.SUCCESS
	err = user.ShowUserInfo(req)
	if err != nil {
		resp.Code = errmsg.ERROR
		return resp, err
	}
	resp.UserDetail = dao.BuildUser(user)
	return resp, nil
}

func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user dao.User
	resp = new(service.UserDetailResponse)
	resp.Code = errmsg.SUCCESS
	err = user.Create(req)
	if err != nil {
		resp.Code = errmsg.ERROR
		return resp, err
	}
	resp.UserDetail = dao.BuildUser(user)
	return resp, nil
}

func (*UserService) UserLogout(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	resp = new(service.UserDetailResponse)
	return resp, nil
}
