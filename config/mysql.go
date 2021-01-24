package config

import (
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
}

func CloseDB() {
	_ = DB.Close()
}

