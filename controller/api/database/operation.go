package database

import (
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/model"
)

func AutoMigrate()(err error){
	Migrate := config.DB.AutoMigrate(&model.User{},&model.Role{},&model.Permission{},
	// 镜像 + 应用表
	&model.Mirror{},&model.Application{},
	// 应用接入 + 我的接入  + 我的设备 applicetionInsert myInsert  myDevice
	&model.ApplicationInsert{},&model.MyInsert{},&model.MyDevice{},
	// 企业库 + 项目库 + 园区库  enterprise + project + park  库 library
	&model.EnterpriseLib{},&model.ProjectLib{},&model.ParkLib{},
	// 月报报表 + 统计报表 Monthly report + statistics report
	&model.MonthlyReport{},&model.StatisticReport{},
	// 模型类型管理 + 模块 + 模板   model + module + Template
	&model.Model{},&model.Module{},&model.Template{},
	)
	if !config.DB.HasTable(&model.Role{}){
		// 删除已有数据库保证操作干净
		if err= Migrate.Error;err!=nil{
			//创建表出错
			panic(err)
		}
	}
	return
}

