package user

import (
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/model"
	"IndustrialInternetApi/model/paginate"
	"github.com/gin-gonic/gin"
)

func GetAllRoles(c *gin.Context) (composer paginate.RoleComposer, err error) {
	page := c.Param("page")
	pagesize := c.Param("pageSize")

	var userList []model.Role
	SQL:=config.DB.Preload("Permission").Model(&model.Role{})
	composer,err = paginate.RolePaginator(SQL,page,pagesize,userList)
	if err != nil {
		return paginate.RoleComposer{}, err
	}
	return
}

func GetRoleById(ID uint) (role model.Role, err error) {
	var userId = ID
	var user model.Role
	err = config.DB.Model(&model.Role{}).Preload("Role").Where("id = ?", userId).First(&user).Error
	found := config.DB.Model(&model.User{}).Preload("Role").Where("id = ?", userId).First(&user).RecordNotFound()
	if found || err != nil {
		return model.Role{}, err
	}
	return user, err
}

func CreateRole(c *gin.Context) (role model.Role,affe int64) {
	var roleIdList model.RoleIdList
	_ = c.BindJSON(&roleIdList)

	var permissionIdList = roleIdList.Ids
	//根据permission ID 取出权限列表
	var permissions []model.Permission
	config.DB.Model(&model.Permission{}).Where("id in (?)", permissionIdList).Find(&permissions)
	// 将列表关系数据绑定到新建角色上
	roleIdList.Role.Permission = permissions
	affe = config.DB.Create(&roleIdList.Role).RowsAffected

	return roleIdList.Role,affe
}

func UpdateRole(c *gin.Context,ID uint) (role model.Role,affe int64) {
	var roleIdList model.RoleIdList
	_ = c.BindJSON(&roleIdList)

	var permissionIdList = roleIdList.Ids
	//根据permission ID 取出权限列表
	var permissions []model.Permission

	config.DB.Model(&model.Permission{}).Where("id in (?)", permissionIdList).Find(&permissions)
	// 将列表关系数据绑定到新建角色上
	roleIdList.Role.Permission = permissions
	affe = config.DB.Model(&model.Role{}).Where("id = ?",ID).Update(&roleIdList.Role).RowsAffected

	// 更新关系
	config.DB.Model(&roleIdList.Role.Permission).Update(&roleIdList.Role.Permission)
	config.DB.Model(&roleIdList.Role).Association("Permission").Replace(&permissions)
	return roleIdList.Role,affe
}

func DeleteRole(ID uint)(affe int64){
	var role model.Role
	affe = config.DB.Model(&role).Where("id = ?", ID).First(&role).Delete(&role).RowsAffected
	return
}