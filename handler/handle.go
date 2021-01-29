package handler

import (
	"context"

	"github.com/itgamerskyHuang/user/domain/model"
	"github.com/itgamerskyHuang/user/domain/service"
	userpb "github.com/itgamerskyHuang/user/proto/user"
)

type User struct {
	UserDataService service.IUserDataService
}

// 注册用户
// 注册
func (u *User) Register(ctx context.Context, r *userpb.RegisterRequest, w *userpb.RegisterResponse) error {
	userRegister := &model.User{
		UserName:     r.UserName,
		FirstName:    r.FirstName,
		HashPassword: r.Pwd,
	}

	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}

	w.Message = "添加成功"

	return nil
}

// 登录
func (u *User) Login(ctx context.Context, r *userpb.LoginRequest, w *userpb.LoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(r.UserName, r.Pwd)
	if err != nil {
		return err
	}
	w.IsSucceed = isOk
	return nil
}

// 查询用户信息
func (u *User) GetUserInfo(ctx context.Context, r *userpb.GetUserInfoRequest, w *userpb.GetUserInfoResponse) error {
	user, err := u.UserDataService.FindUserByName(r.UserName)
	if err != nil {
		return err
	}
	w.UserId = user.Id
	w.UserName = user.UserName
	w.FirstName = user.FirstName
	return nil
}
