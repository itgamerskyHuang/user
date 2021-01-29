package service

import (
	"errors"

	"github.com/itgamerskyHuang/user/domain/model"
	"github.com/itgamerskyHuang/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	// 添加用户
	AddUser(*model.User) (int64, error)
	// 删除用户
	DelectUser(int64) error
	// 更新用户
	UpdateUser(*model.User, bool) error
	// 查找用户
	FindUserByName(string) (*model.User, error)
	// 验证账号密码
	CheckPwd(userName string, pwd string) (bool, error)
}

func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

// 添加用户
func (u *UserDataService) AddUser(user *model.User) (userId int64, err error) {
	pwdByte, err := GeneratePwd(user.HashPassword)
	if err != nil {
		return user.Id, err
	}
	user.HashPassword = string(pwdByte)

	return u.UserRepository.CreateUser(user)
}

// 删除用户
func (u *UserDataService) DelectUser(UserId int64) error {
	return u.UserRepository.DelectUserByID(UserId)
}

// 更新用户
func (u *UserDataService) UpdateUser(user *model.User, isChangePwd bool) error {
	// 判断是否更新了密码
	if isChangePwd {
		pwdByte, err := GeneratePwd(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	return u.UserRepository.UpdataUser(user)
}

// 查找用户
func (u *UserDataService) FindUserByName(userName string) (*model.User, error) {
	return u.UserRepository.FindUserByName(userName)
}

// 验证账号密码
func (u *UserDataService) CheckPwd(userName string, pwd string) (bool, error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}
	return ValidatePwd(pwd, user.HashPassword)
}

// 用户密码加密

func GeneratePwd(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

// 验证密码
func ValidatePwd(userPwd string, hashed string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPwd)); err != nil {
		return false, errors.New("密码错误")
	}
	return true, nil
}
