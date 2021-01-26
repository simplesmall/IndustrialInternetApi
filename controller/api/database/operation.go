package database

import (
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/model"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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
	user :=model.User{}
	if noAdmin :=config.DB.Where("account = ?","admin").First(&user).RecordNotFound();noAdmin{
		SeedDB()
	}
	return
}

func SeedDB()  {
	//初始化插入测试数据
	// 一级目录权限
	per1 := model.Permission{gorm.Model{}, "dataShow", "dataShow", "el-icon-coin", "数据显示", "0", "1", "1"}
	per2 := model.Permission{gorm.Model{}, "appMarket", "appMarket", "el-icon-s-goods", "应用超市", "0", "1", "1"}
	per3 := model.Permission{gorm.Model{}, "operationManage", "operationManage", "el-icon-s-ticket", "运营管理", "0", "1", "1"}
	per4 := model.Permission{gorm.Model{}, "smartConnect", "smartConnect", "dataShow", "智能连接", "0", "1", "1"}
	per5 := model.Permission{gorm.Model{}, "appDevelop", "appDevelop", "el-icon-menu", "应用开发", "0", "1", "1"}
	per6 := model.Permission{gorm.Model{}, "containerManage", "containerManage", "el-icon-s-platform", "容器管理", "0", "1", "1"}
	per7 := model.Permission{gorm.Model{}, "system", "system", "el-icon-s-custom", "系统管理", "0", "1", "1"}
	config.DB.Create(&per1)
	config.DB.Create(&per2)
	config.DB.Create(&per3)
	config.DB.Create(&per4)
	config.DB.Create(&per5)
	config.DB.Create(&per6)
	config.DB.Create(&per7)
	// 系统设置权限
	per71 := model.Permission{gorm.Model{}, "account", "account", "el-icon-coin", "账号管理", "7", "1", "1"}
	per72 := model.Permission{gorm.Model{}, "role", "role", "el-icon-coin", "角色管理", "7", "1", "1"}
	config.DB.Create(&per71)
	config.DB.Create(&per72)
	// 租户用户权限
	per31 := model.Permission{gorm.Model{}, "tenant", "tenant", "el-icon-coin", "租户管理", "3", "1", "1"}
	per32 := model.Permission{gorm.Model{}, "user", "user", "el-icon-coin", "用户管理", "3", "1", "1"}
	config.DB.Create(&per31)
	config.DB.Create(&per32)

	//二级目录权限
	//per11 := model.Permission{gorm.Model{}, "areaMonitor", "areaMonitor", "el-icon-coin", "区域监测", "1", "1", "1"}
	//per12 := model.Permission{gorm.Model{}, "industryMonitor", "industryMonitor", "el-icon-coin", "行业监测", "1", "1", "1"}
	//per13 := model.Permission{gorm.Model{}, "investmentMonitor", "investmentMonitor", "el-icon-coin", "投资监测", "1", "1", "1"}
	//per14 := model.Permission{gorm.Model{}, "enterpriseMonitor", "enterpriseMonitor", "el-icon-coin", "企业监测", "1", "1", "1"}
	//per15 := model.Permission{gorm.Model{}, "parkMonitor", "parkMonitor", "el-icon-coin", "园区监测", "1", "1", "1"}
	//
	//per16 := model.Permission{gorm.Model{}, "productMonitor", "productMonitor", "el-icon-coin", "行业监测", "1", "1", "1"}
	//per17 := model.Permission{gorm.Model{}, "enterpriLibManage", "enterpriLibManage", "el-icon-coin", "行业监测", "1", "1", "1"}
	//per18 := model.Permission{gorm.Model{}, "projectLibManage", "projectLibManage", "el-icon-coin", "行业监测", "1", "1", "1"}
	//per19 := model.Permission{gorm.Model{}, "areaLibManage", "areaLibManage", "el-icon-coin", "行业监测", "1", "1", "1"}
	//per110 := model.Permission{gorm.Model{}, "getData", "getData", "el-icon-coin", "行业监测", "1", "1", "1"}
	//
	//per111 := model.Permission{gorm.Model{}, "getDataCheck", "getDataCheck", "el-icon-coin", "数据采集审核", "1", "1", "1"}
	//per112 := model.Permission{gorm.Model{}, "tableManage", "tableManage", "el-icon-coin", "报表管理", "1", "1", "1"}
	//per113 := model.Permission{gorm.Model{}, "enterpriseSubData", "enterpriseSubData", "el-icon-coin", "企业数据上报详情", "1", "1", "1"}
	//per114 := model.Permission{gorm.Model{}, "mostDataShow", "mostDataShow", "el-icon-coin", "大数据展示", "1", "1", "1"}
	//config.DB.Create(per11)
	//config.DB.Create(per11)
	//config.DB.Create(per11)


	//系统管理员权限
	normalPermissions :=[]model.Permission{per1,per2,per3,per4,per5,per6}
	enterprisePermissions :=[]model.Permission{per1,per2,per3,per4,per5,per6,per31,per32}
	adminPermissions :=[]model.Permission{per1,per2,per3,per4,per5,per6,per7,per71,per72,per31,per32}
	//系统管理员角色
	normalRole:=model.Role{gorm.Model{}, "普通用户", "公用权限", normalPermissions}
	enterpriseRole:=model.Role{gorm.Model{}, "企业用户", "公用权限", enterprisePermissions}
	adminRole:=model.Role{gorm.Model{}, "超级管理", "公用权限", adminPermissions}

	//创建超级管理员+企业用户+普通用户
	admin:=model.User{gorm.Model{}, "系统管理员", "admin", formatPwd("123456"), "0", "1", "1", "", 0, "", "127.0.0.1", "no", "1","13000011234","admin@example.com", "系统管理员",[]model.Role{adminRole}}
	enterprise:=model.User{gorm.Model{}, "企业用户", "enterprise", formatPwd("123456"), "1", "1", "1", "", 0, "", "127.0.0.1", "no", "2","13023311234","enterprise@example.com", "企业用户", []model.Role{enterpriseRole}}
	normal:=model.User{gorm.Model{}, "普通用户", "normal", formatPwd("123456"), "0", "1", "1", "", 0, "", "127.0.0.1", "no", "3", "13067871234","normal@example.com", "普通用户",[]model.Role{normalRole}}
	config.DB.Create(&admin)
	config.DB.Create(&enterprise)
	config.DB.Create(&normal)
}

func formatPwd(pwd string) (formatPwd string) {
	inputPwd := []byte(pwd)
	//生成hash存入数据库
	hashPwd, _ := bcrypt.GenerateFromPassword(inputPwd, bcrypt.DefaultCost) //password为string类型
	formatPwd = string(hashPwd)
	return formatPwd
}
