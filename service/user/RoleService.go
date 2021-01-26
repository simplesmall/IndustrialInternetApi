package user

import (
	"IndustrialInternetApi/common/middleware"
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/model"
	"IndustrialInternetApi/model/paginate"
	"IndustrialInternetApi/service/jwt"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
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
	err = config.DB.Model(&model.Role{}).Preload("Permission").Where("id = ?", userId).First(&user).Error
	found := config.DB.Model(&model.Role{}).Preload("Permission").Where("id = ?", userId).First(&user).RecordNotFound()
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


// 用户
type UserForm struct {
	model.User
	Roles []int `json:"Roles"`
}

func GetUserItem(userId int) (user UserForm, err error) {
	var userInfo UserForm
	result := config.DB.Preload("Role").First(&userInfo.User, userId)

	var roleIds []int
	for i := 0; i < len(userInfo.Role); i++ {
		roleIds = append(roleIds, int(userInfo.Role[i].ID))
	}

	userInfo.Roles = roleIds

	if result.Error != nil {
		return userInfo, errors.New("user not find")
	}
	return userInfo, nil
}

//返回用户及权限
type UserAndPermission struct {
	User        model.User     `json:"user"`
	Permissions []string `json:"permissions"`
}

func GetLoginUser() (UserAndPermission, bool) {
	j := jwt.NewJWT()
	jwtClaims, _ := j.ParseToken(middleware.AuthToken)

	var resultData UserAndPermission

	result := config.DB.Preload("Role").First(&resultData.User, jwtClaims.Id)
	if result.Error != nil {
		return resultData, false
	}

	if resultData.User.Type == "1" {
		var permissions []model.Permission
		config.DB.Find(&permissions)
		for _, value := range permissions {
			resultData.Permissions = append(resultData.Permissions, strconv.Itoa(int(value.ID))) // 追加1个元素
		}
	} else {
		var roles []model.Role
		config.DB.Model(&resultData.User).Association("Role").Find(&roles)

		for _, item := range roles { //item是值拷贝
			var permissions []model.Permission
			config.DB.Model(&item).Association("Permission").Find(&permissions)
			for _, value := range permissions {
				resultData.Permissions = append(resultData.Permissions, strconv.Itoa(int(value.ID))) // 追加1个元素
			}
		}
		resultData.Permissions = removeDuplicateElement(resultData.Permissions)
	}
	return resultData, true
}

//数组去重
func removeDuplicateElement(data []string) []string {
	result := make([]string, 0, len(data))
	temp := map[string]struct{}{}
	for _, item := range data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}