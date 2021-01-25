package routers

import (
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/controller/api/auth"
	"IndustrialInternetApi/controller/api/database"
	version1 "IndustrialInternetApi/controller/api/v1"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.Default()
	//数据库初始化
	config.InitConnect()
	defer config.CloseDB()
	//创建数据库表
	_ = database.AutoMigrate()
	//token中间件 JWT 放行路由
	//swagger 配置
	// token 测试
	auther :=r.Group("auth")
	// 使用token鉴权中间件 下面这几个API放入白名单
	{
		auther.POST("/login", auth.LoginHandler)
		auther.GET("/logout",auth.LogoutHandler)
		auther.GET("/updateToken",auth.InitiativeExpireHandler)
	}
	api := r.Group("api")
	//api.Use(jwt.JWTAuth())
	{
		v1 := api.Group("v1")

		// RBAC
		{
			// 获取全部  用户  权限  角色
			v1.GET("/permissions", version1.PermissionsHandler)
			v1.GET("/permissionsTrees", version1.PermissionsTreeHandler)
			v1.GET("/permissionTree/:pid", version1.GetPermissionFamilyHandler)
			v1.GET("/roles", version1.RolesHandler)
			v1.GET("/users", version1.UsersHandler)

			// 根据用户ID获得用户User Roles Permissions
			v1.GET("/permission/:id", version1.PermissionByIdHandler)
			v1.GET("/role/:id", version1.RoleByIdHandler)
			v1.GET("/user/:id", version1.UserByIdHandler)

			// Permission 的创建,更改,删除
			v1.POST("/permission",version1.CreatePermissionHandler)
			v1.PUT("/permission/:id", version1.UpdatePerHandler)
			v1.DELETE("/permission/:id", version1.DeletePermissionHandler)

			// Role 的创建,更新,删除
			v1.POST("/role",version1.CreateRoleHandler)
			v1.PUT("/role/:id", version1.UpdateRoleHandler)
			v1.DELETE("/role/:id", version1.DeleteRoleHandler)

			// User 的创建,更新,删除
			v1.POST("/user",version1.CreateUserHandler)
			v1.PUT("/user/:id", version1.UpdateUserHandler)
			v1.DELETE("/user/:id", version1.DeleteUserHandler)
		}

		// 镜像表+应用表  mirror  application
		{
			v1.GET("/mirrors",version1.GetAllMirrorsHandler)
			v1.GET("/applications",version1.GetAllApplicationsHandler)
			v1.GET("/mirror/:id",version1.GetMirrorByIdHandler)
			v1.GET("/application/:id",version1.GetApplicationByIdHandler)

			v1.POST("/mirror",version1.CreateMirrorHandler)
			v1.PUT("/mirror/:id",version1.UpdateMirrorHandler)
			v1.DELETE("/mirror/:id",version1.DeleteMirrorHandler)

			v1.POST("/application",version1.CreateApplicationHandler)
			v1.PUT("/application/:id",version1.UpdateApplicationHandler)
			v1.DELETE("/application/:id",version1.DeleteApplicationHandler)
		}

		// 应用接入 + 我的接入  + 我的设备 applicetionInsert myInsert  myDevice
		{
			v1.GET("/applicationInserts",version1.GetAllApplicationInsertsHandler)
			v1.GET("/myInserts",version1.GetAllMyInsertsHandler)
			v1.GET("/myDevices",version1.GetAllMyDevicesHandler)
			v1.GET("/applicationInsert/:id",version1.GetApplicationInsertByIdHandler)
			v1.GET("/myInsert/:id",version1.GetMyInsertByIdHandler)
			v1.GET("/myDevice/:id",version1.GetMyDeviceByIdHandler)

			v1.POST("/applicationInsert",version1.CreateApplicationInsertHandler)
			v1.PUT("/applicationInsert/:id",version1.UpdateApplicationInsertHandler)
			v1.DELETE("/applicationInsert/:id",version1.DeleteApplicationInsertHandler)

			v1.POST("/myInsert",version1.CreateMyInsertHandler)
			v1.PUT("/myInsert/:id",version1.UpdateMyInsertHandler)
			v1.DELETE("/myInsert/:id",version1.DeleteMyInsertHandler)

			v1.POST("/myDevice",version1.CreateMyDeviceHandler)
			v1.PUT("/myDevice/:id",version1.UpdateMyDeviceHandler)
			v1.DELETE("/myDevice/:id",version1.DeleteMyDeviceHandler)
		}

		// 企业库 + 项目库 + 园区库  enterprise + project + park  库 library
		{
			v1.GET("/enterpriseLibs",version1.GetAllEnterPrisesHandler)
			v1.GET("/projectLibs",version1.GetAllProjectsHandler)
			v1.GET("/parkLibs",version1.GetAllParksHandler)
			v1.GET("/enterpriseLib/:id",version1.GetEnterPriseByIdHandler)
			v1.GET("/projectLib/:id",version1.GetProjectByIdHandler)
			v1.GET("/parkLib/:id",version1.GetParkByIdHandler)

			v1.POST("/enterpriseLib",version1.CreateEnterPriseHandler)
			v1.PUT("/enterpriseLib/:id",version1.UpdateEnterPriseHandler)
			v1.DELETE("/enterpriseLib/:id",version1.DeleteEnterPriseHandler)

			v1.POST("/projectLib",version1.CreateProjectHandler)
			v1.PUT("/projectLib/:id",version1.UpdateProjectHandler)
			v1.DELETE("/projectLib/:id",version1.DeleteProjectHandler)

			v1.POST("/parkLib",version1.CreateParkHandler)
			v1.PUT("/parkLib/:id",version1.UpdateParkHandler)
			v1.DELETE("/parkLib/:id",version1.DeleteParkHandler)
		}

		// 月报报表 + 统计报表 Monthly report + statistics report
		{
			v1.GET("/monthlyReports",version1.GetAllMonthlyReportsHandler)
			v1.GET("/statisticReports",version1.GetAllStatisticReportsHandler)
			v1.GET("/monthlyReport/:id",version1.GetMonthlyReportByIdHandler)
			v1.GET("/statisticReport/:id",version1.GetStatisticReportByIdHandler)

			v1.POST("/monthlyReport",version1.CreateMonthlyReportHandler)
			v1.PUT("/monthlyReport/:id",version1.UpdateMonthlyReportHandler)
			v1.DELETE("/monthlyReport/:id",version1.DeleteMonthlyReportHandler)

			v1.POST("/statisticReport",version1.CreateStatisticReportHandler)
			v1.PUT("/statisticReport/:id",version1.UpdateStatisticReportHandler)
			v1.DELETE("/statisticReport/:id",version1.DeleteStatisticReportHandler)
		}

		// 模型类型管理 + 模块 + 模板   model + module + Template
		{
			v1.GET("/models",version1.GetAllModelsHandler)
			v1.GET("/modules",version1.GetAllModulesHandler)
			v1.GET("/templates",version1.GetAllTemplatesHandler)
			v1.GET("/model/:id",version1.GetModelByIdHandler)
			v1.GET("/module/:id",version1.GetModuleByIdHandler)
			v1.GET("/template/:id",version1.GetTemplateByIdHandler)

			v1.POST("/model",version1.CreateModelHandler)
			v1.PUT("/model/:id",version1.UpdateModelHandler)
			v1.DELETE("/model/:id",version1.DeleteModelHandler)

			v1.POST("/module",version1.CreateModuleHandler)
			v1.PUT("/module/:id",version1.UpdateModuleHandler)
			v1.DELETE("/module/:id",version1.DeleteModuleHandler)

			v1.POST("/template",version1.CreateTemplateHandler)
			v1.PUT("/template/:id",version1.UpdateTemplateHandler)
			v1.DELETE("/template/:id",version1.DeleteTemplateHandler)
		}
	}
	_ = r.Run(":8090")
}
