package repository

import (
	"github.com/itgamerskyHuang/user/domain/model"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	// 初始化数据库
	InitTable() error
	// 查找用户
	FindUserByName(string) (*model.User, error)
	// 按用户ID查找拥挤
	FindUserByID(int64) (*model.User, error)
	// 创建用户
	CreateUser(*model.User) (int64, error)
	// 根据用户ID删除用户
	DelectUserByID(int64) error
	// 更新用户
	UpdataUser(*model.User) error
	// 查找所有用户
	FindAllUser() ([]*model.User, error)
}

type UserRepository struct {
	mysqldb *gorm.DB
}

//
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		mysqldb: db,
	}
}

// 初始化表
func (u *UserRepository) InitTable() error {
	// return u.mysqldb.CreateTable(&model.User{}).Error
	return u.mysqldb.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.User{}).Error
}

// 根据用户名称查找用户信息
func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqldb.Where("user_name = ?", name).Find(user).Error
}

// 按用户ID查找拥挤
func (u *UserRepository) FindUserByID(userId int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqldb.First(user, userId).Error
}

// 创建用户
func (u *UserRepository) CreateUser(user *model.User) (userId int64, err error) {
	return user.Id, u.mysqldb.Create(user).Error
}

// 根据用户ID删除用户
func (u *UserRepository) DelectUserByID(userId int64) error {
	return u.mysqldb.Where("id = ?", userId).Delete(&model.User{}).Error
}

// 更新用户
func (u *UserRepository) UpdataUser(user *model.User) error {
	return u.mysqldb.Model(user).Update(&user).Error
}

// 查找所有用户
func (u *UserRepository) FindAllUser() (userall []*model.User, err error) {
	return userall, u.mysqldb.Find(&userall).Error
}
