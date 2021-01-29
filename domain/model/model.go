package model

type User struct {
	// 主键ID
	Id int64 `gorm:"PRIMARY_KEY;NOT_NULL;AUTO_INCREMENT"`
	// 用户名称
	UserName string `gorm:"NOT_NULL;UNIQUE_INDEX"`
	// 添加需要的字段
	FirstName string
	// 密码
	HashPassword string
}
