package user

import (
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/model"
	"github.com/gin-gonic/gin"
)

func GetAllUsers() (users []model.User, err error) {
	err = config.DB.Preload("Role").Preload("Role.Permission").Model(&model.User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetUserById(ID uint) (user model.User, err error) {
	var tempUser model.User
	err=config.DB.Model(&model.User{}).Preload("Role").Where("id = ?", ID).First(&tempUser).Error
	user = tempUser
	if err != nil {
		return tempUser, err
	}
	return user,err
}

func CreateUser(c *gin.Context) (user model.User, affe int64) {
	var userIdList model.UserIdList
	_ = c.BindJSON(&userIdList)
	var idList = userIdList.Ids
	//根据idList ID 取出角色列表
	var roles []model.Role
	config.DB.Model(&model.Role{}).Where("id in (?)", idList).Find(&roles)
	// 将列表关系数据绑定到新建用户上
	userIdList.User.Role = roles
	affe = config.DB.Create(&userIdList.User).RowsAffected
	return userIdList.User,affe
}

func UpdateUser(c *gin.Context,ID uint) (role model.User,affe int64) {
	var userIdList model.UserIdList
	_ = c.BindJSON(&userIdList)

	var roleIdList = userIdList.Ids
	//根据permission ID 取出权限列表
	var roles []model.Role
	config.DB.Where("id = ?",ID).First(&userIdList.User)

	config.DB.Model(&model.Role{}).Where("id in (?)", roleIdList).Find(&roles)
	// 将列表关系数据绑定到新建角色上
	userIdList.User.Role = roles
	affe = config.DB.Model(&userIdList.User).Where("id = ?",ID).Update(&userIdList.User).RowsAffected

	// 更新关系
	config.DB.Model(&userIdList.User.Role).Update(&userIdList.User.Role)

	config.DB.Model(&userIdList.User).Association("Role").Replace(roles)

	return userIdList.User,affe
}

func DeleteUser(ID uint)(affe int64){
	var user model.User
	affe = config.DB.Model(&user).Where("id = ?", ID).First(&user).Delete(&user).RowsAffected
	return
}