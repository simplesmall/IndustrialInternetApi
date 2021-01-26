package user

import (
	"IndustrialInternetApi/common/utils"
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/model"
	"github.com/gin-gonic/gin"
)

func GetAllPermissions() (pers []model.Permission, err error) {
	err = config.DB.Model(&model.Permission{}).Find(&pers).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetPermisById(ID uint) (pers []model.Permission, err error) {
	var userId = ID
	// 根据UserID 拿到 Roles :user.Roles
	var user model.User
	err = config.DB.Model(&model.User{}).Preload("Role").Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	// 将user.Roles取出保存一个rolesID => 将roles的ID重构出来
	var roles []model.Role
	roles = user.Role
	rolesID := []uint{}
	for _, v := range roles {
		rolesID = append(rolesID, v.ID)
	}

	//取出permission
	var roleForPermssions []model.Role
	err = config.DB.Model(&model.Role{}).Preload("Permission").Where("id in (?)", rolesID).Find(&roleForPermssions).Error
	if err != nil {
		return nil, err
	}
	// 整合输出对应用户所有权限 Permissions
	var permissions []model.Permission
	for _, v := range roleForPermssions {
		permissions = append(permissions, v.Permission...)
	}
	pers = permissions
	//此处格式化权限表树形输出
	return pers, err
}
func GetUserPermissionTreeById(ID uint) (dataList []interface{}, err error) {
	var userId = ID
	// 根据UserID 拿到 Roles :user.Roles
	var user model.User
	err = config.DB.Model(&model.User{}).Preload("Role").Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	// 将user.Roles取出保存一个rolesID => 将roles的ID重构出来
	var roles []model.Role
	roles = user.Role
	rolesID := []uint{}
	for _, v := range roles {
		rolesID = append(rolesID, v.ID)
	}

	//取出permission
	var roleForPermssions []model.Role
	err = config.DB.Model(&model.Role{}).Preload("Permission").Where("id in (?)", rolesID).Find(&roleForPermssions).Error
	if err != nil {
		return nil, err
	}
	// 整合输出对应用户所有权限 Permissions
	var permissions []model.Permission
	for _, v := range roleForPermssions {
		permissions = append(permissions, v.Permission...)
	}
	for _, v := range permissions {
		// 获取一条数据
		parent := model.PermissionTree{v.Name,v.Url,v.Icon,v.Describe,v.ParentId,v.Status,v.Type, []*model.PermissionTree{}}

		// 将数据对应的pid的所有数据查出放入 childrenList
		var childrenList []model.Permission
		config.DB.Where("parent_id = ?", v.ID).Find(&childrenList)

		for _, c := range childrenList {
			// 对childrenList中的每一项执行查询
			child := model.PermissionTree{c.Name,c.Url,c.Icon,c.Describe,c.ParentId,c.Status,c.Type, []*model.PermissionTree{}}
			parent.Children = append(parent.Children, &child)
		}
		dataList = append(dataList, parent)
	}
	return dataList, err
}

/*
{
	"name":"haoren",
	"url":"http1234",
	"icon":"121323",
	"describe":"deddd",
	"parent_id": "0",
    "status": "1",
    "type": "2"
}
*/
func CreatePermission(c *gin.Context)(per model.Permission,affer int64){
	var permission model.Permission
	_ = c.BindJSON(&permission)
	affected := config.DB.Create(&permission).RowsAffected
	return permission,affected
}

func UpdatePermission(c *gin.Context,ID uint)(per model.Permission,affer int64){
	var permission model.Permission
	_ = c.BindJSON(&permission)
	affer=config.DB.Model(&permission).Where("id = ?", ID).Update(&permission).RowsAffected
	return permission,affer
}

func DeletePermission(ID uint)(affe int64){
	var permission model.Permission
	affe = config.DB.Model(&permission).Where("id = ?", ID).Delete(&permission).RowsAffected
	return
}

//构建树形结构
func GetPermissionTree() (dataList []interface{}, err error) {
	var parentList []model.Permission
	//获取父节点
	config.DB.Where("parent_id = ?", uint(0)).Find(&parentList)

	for _, v := range parentList {
		// 获取一条数据
		parent := model.PermissionTree{v.Name,v.Url,v.Icon,v.Describe,v.ParentId,v.Status,v.Type, []*model.PermissionTree{}}

		// 将数据对应的pid的所有数据查出放入 childrenList
		var childrenList []model.Permission
		config.DB.Where("parent_id = ?", v.ID).Find(&childrenList)

		for _, c := range childrenList {
			// 对childrenList中的每一项执行查询
			child := model.PermissionTree{c.Name,c.Url,c.Icon,c.Describe,c.ParentId,c.Status,c.Type, []*model.PermissionTree{}}
			// 将数据对应的pid的所有数据查出放入 childrenList========
			//var thirdList []Model.Permission
			//MySQL.DB.Where("p_id = ?", v.ID).Find(&thirdList)
			//	for _,j :=range thirdList{
			//		third := Model.PermissionTree{j.Desc,j.PID, []*Model.PermissionTree{}}
			//		child.Children = append(child.Children, &third)
			//	}
			// 将数据对应的pid的所有数据查出放入 childrenList========
			parent.Children = append(parent.Children, &child)
		}
		dataList = append(dataList, parent)
	}
	return dataList, nil
}

func GetPermissionFamily(ID uint)(list []model.Permission,err error)  {
	var per model.Permission
	var pers []model.Permission

	config.DB.Model(&per).Where("id = ?",ID).First(&per)
	pers=append(pers, per)
	pid := per.ParentId
	for{
		if utils.StrToUInt(pid) == uint(0) {
			break
		}
		var tempPer model.Permission
		config.DB.Model(&per).Where("id = ?",pid).First(&tempPer)
		pid = tempPer.ParentId
		pers=append(pers, tempPer)
	}
	return pers,err
}
