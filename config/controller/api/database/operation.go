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
	//per32 := model.Permission{gorm.Model{}, "user", "user", "el-icon-coin", "用户管理", "3", "1", "1"}
	config.DB.Create(&per31)
	//config.DB.Create(&per32)

	// 数据展示 二级目录权限
	per11 := model.Permission{gorm.Model{}, "areaMonitor", "areaMonitor", "el-icon-coin", "区域监测", "1", "1", "1"}
	per12 := model.Permission{gorm.Model{}, "industryMonitor", "industryMonitor", "el-icon-coin", "行业监测", "1", "1", "1"}
	per13 := model.Permission{gorm.Model{}, "investmentMonitor", "investmentMonitor", "el-icon-coin", "投资监测", "1", "1", "1"}
	per14 := model.Permission{gorm.Model{}, "enterpriseMonitor", "enterpriseMonitor", "el-icon-coin", "企业监测", "1", "1", "1"}
	per15 := model.Permission{gorm.Model{}, "parkMonitor", "parkMonitor", "el-icon-coin", "园区监测", "1", "1", "1"}

	per16 := model.Permission{gorm.Model{}, "productMonitor", "productMonitor", "el-icon-coin", "行业监测", "1", "1", "1"}
	per17 := model.Permission{gorm.Model{}, "enterpriLibManage", "enterpriLibManage", "el-icon-coin", "行业监测", "1", "1", "1"}
	per18 := model.Permission{gorm.Model{}, "projectLibManage", "projectLibManage", "el-icon-coin", "行业监测", "1", "1", "1"}
	per19 := model.Permission{gorm.Model{}, "areaLibManage", "areaLibManage", "el-icon-coin", "行业监测", "1", "1", "1"}
	per110 := model.Permission{gorm.Model{}, "getData", "getData", "el-icon-coin", "行业监测", "1", "1", "1"}

	per111 := model.Permission{gorm.Model{}, "getDataCheck", "getDataCheck", "el-icon-coin", "数据采集审核", "1", "1", "1"}
	per112 := model.Permission{gorm.Model{}, "tableManage", "tableManage", "el-icon-coin", "报表管理", "1", "1", "1"}
	per113 := model.Permission{gorm.Model{}, "enterpriseSubData", "enterpriseSubData", "el-icon-coin", "企业数据上报详情", "1", "1", "1"}
	per114 := model.Permission{gorm.Model{}, "mostDataShow", "mostDataShow", "el-icon-coin", "大数据展示", "1", "1", "1"}
	config.DB.Create(&per11)
	config.DB.Create(&per12)
	config.DB.Create(&per13)
	config.DB.Create(&per14)
	config.DB.Create(&per15)
	config.DB.Create(&per16)
	config.DB.Create(&per17)
	config.DB.Create(&per18)
	config.DB.Create(&per19)
	config.DB.Create(&per110)
	config.DB.Create(&per111)
	config.DB.Create(&per112)
	config.DB.Create(&per113)
	config.DB.Create(&per114)

	// 应用超市 二级权限目录
	per21 := model.Permission{gorm.Model{}, "market", "market", "el-icon-coin", "我的超市", "2", "1", "1"}
	per22 := model.Permission{gorm.Model{}, "appDetail", "appDetail", "el-icon-coin", "应用详情", "2", "1", "1"}
	per23 := model.Permission{gorm.Model{}, "appDock", "appDock", "el-icon-coin", "应用对接", "2", "1", "1"}
	per24 := model.Permission{gorm.Model{}, "appOpen", "appOpen", "el-icon-coin", "在线开通", "2", "1", "1"}
	config.DB.Create(&per21)
	config.DB.Create(&per22)
	config.DB.Create(&per23)
	config.DB.Create(&per24)

	// 运营管理 二级权限目录
	per301 := model.Permission{gorm.Model{}, "modelType", "modelType", "el-icon-coin", "模块类型", "3", "1", "1"}
	per302 := model.Permission{gorm.Model{}, "modelManage", "modelManage", "el-icon-coin", "模块管理", "3", "1", "1"}
	per303 := model.Permission{gorm.Model{}, "template", "template", "el-icon-coin", "模板管理", "3", "1", "1"}
	per304 := model.Permission{gorm.Model{}, "enterUser", "enterUser", "el-icon-coin", "用户管理", "3", "1", "1"}
	per35 := model.Permission{gorm.Model{}, "enterprise", "enterprise", "el-icon-coin", "企业认证审核", "3", "1", "1"}

	per36 := model.Permission{gorm.Model{}, "order", "order", "el-icon-coin", "订单管理", "3", "1", "1"}
	per37 := model.Permission{gorm.Model{}, "appAudit", "appAudit", "el-icon-coin", "应用审核", "3", "1", "1"}
	per38 := model.Permission{gorm.Model{}, "personal", "personal", "el-icon-coin", "个人认证审核", "3", "1", "1"}
	per39 := model.Permission{gorm.Model{}, "content", "content", "el-icon-coin", "内容认证审核", "3", "1", "1"}
	per310 := model.Permission{gorm.Model{}, "statistics", "statistics", "el-icon-coin", "运营统计分析", "3", "1", "1"}

	per311 := model.Permission{gorm.Model{}, "appMonitor", "appMonitor", "el-icon-coin", "应用监控与预警", "3", "1", "1"}
	per312 := model.Permission{gorm.Model{}, "docker", "docker", "el-icon-coin", "Docker监控", "3", "1", "1"}
	config.DB.Create(&per301)
	config.DB.Create(&per302)
	config.DB.Create(&per303)
	config.DB.Create(&per304)
	config.DB.Create(&per35)
	config.DB.Create(&per36)
	config.DB.Create(&per37)
	config.DB.Create(&per38)
	config.DB.Create(&per39)
	config.DB.Create(&per310)
	config.DB.Create(&per311)
	config.DB.Create(&per312)

	// 智能连接 二级目录权限
	per41 := model.Permission{gorm.Model{}, "access", "access", "el-icon-coin", "接入管理", "4", "1", "1"}
	per42 := model.Permission{gorm.Model{}, "device", "device", "el-icon-coin", "设备管理", "4", "1", "1"}
	per43 := model.Permission{gorm.Model{}, "configuration", "configuration", "el-icon-coin", "组态管理", "4", "1", "1"}
	per44 := model.Permission{gorm.Model{}, "mechanism", "mechanism", "el-icon-coin", "机构管理", "4", "1", "1"}
	per45 := model.Permission{gorm.Model{}, "assembly", "assembly", "el-icon-coin", "组件管理", "4", "1", "1"}
	config.DB.Create(&per41)
	config.DB.Create(&per42)
	config.DB.Create(&per43)
	config.DB.Create(&per44)
	config.DB.Create(&per45)

	// 应用开发 二级目录权限
	per51 := model.Permission{gorm.Model{}, "myApp", "myApp", "el-icon-coin", "我的应用", "5", "1", "1"}
	per52 := model.Permission{gorm.Model{}, "developDom", "developDom", "el-icon-coin", "开发文档", "5", "1", "1"}
	per53 := model.Permission{gorm.Model{}, "appAccess", "appAccess", "el-icon-coin", "应用接入", "5", "1", "1"}
	config.DB.Create(&per51)
	config.DB.Create(&per52)
	config.DB.Create(&per53)

	// 容器管理 二级目录权限
	per61 := model.Permission{gorm.Model{}, "image", "image", "el-icon-coin", "镜像管理", "6", "1", "1"}
	per62 := model.Permission{gorm.Model{}, "container", "container", "el-icon-coin", "容器列表", "6", "1", "1"}
	config.DB.Create(&per61)
	config.DB.Create(&per62)


	//系统管理员权限
	normalPermissions :=[]model.Permission{per1,per2,per3,per4,per5,per6,
		per11,per12,per13,per14,per14,per15,per16,per16,per17,per18,per19,per110,per111,per112,per113,per114,
		per21,per22,per23,per24,
		per301,per302,per303,per35,per36,per37,per38,per39,per310,per311,per312,
		per41,per42,per43,per44,per45,
		per51,per52,per53,
		per61,per62,
	}
	enterprisePermissions :=[]model.Permission{per1,per2,per3,per4,per5,per6,
		per11,per12,per13,per14,per14,per15,per16,per16,per17,per18,per19,per110,per111,per112,per113,per114,
		per21,per22,per23,per24,
		per301,per302,per303,per304,per35,per36,per37,per38,per39,per310,per311,per312,
		per41,per42,per43,per44,per45,
		per51,per52,per53,
		per61,per62,
		}
	adminPermissions :=[]model.Permission{per1,per2,per3,per4,per5,per6,per7,per71,per72,per31,
		per11,per12,per13,per14,per14,per15,per16,per16,per17,per18,per19,per110,per111,per112,per113,per114,
		per21,per22,per23,per24,
		per301,per302,per303,per304,per35,per36,per37,per38,per39,per310,per311,per312,
		per41,per42,per43,per44,per45,
		per51,per52,per53,
		per61,per62,
		}
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
