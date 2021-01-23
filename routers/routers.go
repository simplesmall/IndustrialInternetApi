package routers

import (
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/controller/api/auth"
	version1 "IndustrialInternetApi/controller/api/v1"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.Default()
	//数据库初始化
	config.InitConnect()
	//创建数据库表
	//token中间件 JWT 放行路由
	//swagger 配置
	api := r.Group("api")
	{
		// token 测试
		auther :=api.Group("auther")
		{
			auther.POST("/login", auth.LoginHandler)
			auther.GET("/logout",auth.LogoutHandler)
			auther.GET("/updateToken",auth.InitiativeExpireHandler)
		}

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

		// 应用接入 + 我的接入 applicetionInsert myInsert

		// 我的设备  myDevice

		// 企业库 + 项目库 + 园区库  enterprise + project + park  库 library

		// 月报报表	 Monthly report

		// 模型类型管理 + 模块 + 模板   model + module + Template

		// 统计报表	statistics report
	}
	_ = r.Run(":8090")
}
