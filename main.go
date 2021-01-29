package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/itgamerskyHuang/user/domain/repository"
	userservice "github.com/itgamerskyHuang/user/domain/service"
	"github.com/itgamerskyHuang/user/handler"
	pb "github.com/itgamerskyHuang/user/proto/user"
	"github.com/jinzhu/gorm"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
)

func main() {
	// 创建V2版本的micro实例
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
		micro.Address(":8999"),
	)
	// 初始化服务
	srv.Init()

	// 创建数据库
	db, err := gorm.Open("mysql", "root:123456@tcp(192.168.0.201:3306)/micro?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//一个坑，不设置这个参数，gorm会把表名转义后加个s，导致找不到数据库的表
	db.SingularTable(true)
	// 创建数据库操作实例
	rp := repository.NewUserRepository(db)
	// 创建数据库表 执行一次后注释
	// rp.InitTable()
	//  将数据库操作实例传到user_data_serivce中
	userdataservice := userservice.NewUserDataService(rp)

	// Register handler
	err = pb.RegisterUserHandler(srv.Server(), &handler.User{
		UserDataService: userdataservice,
	})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
