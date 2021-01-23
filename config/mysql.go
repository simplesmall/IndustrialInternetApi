package config

import (
	"IndustrialInternetApi/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitConnect()  {
	var err error
	DB,err=gorm.Open("mysql","root:root@(127.0.0.1:3306)/dustapi?charset=utf8&parseTime=true")
	if err != nil {
		// 连接数据出错
		panic(err)
	}
	DB.SingularTable(true)
	if !DB.HasTable(&model.Role{}){
		// 删除已有数据库保证操作干净
		if err=AutoMigrate();err!=nil{
			//创建表出错
			panic(err)
		}
	}
}

func CloseDB() {
	_ = DB.Close()
}

func AutoMigrate()(err error){
	err = DB.AutoMigrate(&model.User{},&model.Role{},&model.Permission{}).Error
	return
}
