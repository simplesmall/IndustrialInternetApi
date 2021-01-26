package user

import (
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/model"
	"IndustrialInternetApi/model/paginate"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(c *gin.Context) (Composer paginate.UserComposer, err error) {
	page := c.Param("page")
	pagesize := c.Param("pageSize")

	var userList []model.User
	SQL:=config.DB.Model(&model.User{}).Preload("Role").Preload("Role.Permission")
	Composer,err = paginate.UserPaginator(SQL,page,pagesize,userList)
	return Composer,err
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

/*
	"name": "",
	"account": "hhudsuhi",
	"password": "hoiuh",
	"belong_to": "oi",
	"status": "boiibo",
	"is_login": "oiboi",
	"expire_time": "oiboi",
	"login_count": 0,
	"using_time": "",
	"ip": "",
	"ids":[1,2,3]
*/
func CreateUser(c *gin.Context) (user model.User, affe int64) {
	var userIdList model.UserIdList
	_ = c.BindJSON(&userIdList)
	var idList = userIdList.Ids
	//处理密码
	inputPwd := []byte(userIdList.User.Password)
	//生成hash存入数据库   bcrypt.GenerateFromPassword(originPwd, bcrypt.DefaultCost)    bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	hashPwd, _ := bcrypt.GenerateFromPassword(inputPwd, bcrypt.DefaultCost) //password为string类型
	userIdList.User.Password = string(hashPwd)
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
	inputPwd := []byte(userIdList.User.Password)
	//生成hash存入数据库
	hashPwd, _ := bcrypt.GenerateFromPassword(inputPwd, bcrypt.DefaultCost) //password为string类型
	//根据permission ID 取出权限列表
	var roles []model.Role

	config.DB.Model(&model.Role{}).Where("id in (?)", roleIdList).Find(&roles)
	//// 将列表关系数据绑定到新建角色上
	userIdList.User.Password = string(hashPwd)
	userIdList.User.Role = roles
	affe = config.DB.Model(&model.User{}).Where("id = ?",ID).Update(&userIdList.User).RowsAffected

	// 更新关系
	config.DB.Model(&userIdList.User.Role).Update(&userIdList.User.Role)
	config.DB.Model(&userIdList.User).Association("Role").Replace(&roles)

	return userIdList.User,affe
}

func DeleteUser(ID uint)(affe int64){
	var user model.User
	affe = config.DB.Model(&user).Where("id = ?", ID).First(&user).Delete(&user).RowsAffected
	return
}