package independent

import (
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/model"
	"IndustrialInternetApi/model/paginate"
	"github.com/gin-gonic/gin"
)

func GetAllMirrors(c *gin.Context) (mirrors paginate.MirrorComposer,err error) {
	page := c.Param("page")
	pagesize := c.Param("pageSize")
	var userList []model.Mirror
	SQL:=config.DB.Model(&model.Mirror{})
	mirrors,err = paginate.MirrorPaginator(SQL,page,pagesize,userList)
	return	mirrors,err
}
func GetMirrorById(ID uint) (mirror model.Mirror,err error) {
	if err =config.DB.Where("id = ?",ID).First(&mirror).Error;err != nil{
		return model.Mirror{}, err
	}
	return
}
func CreateMirror(c *gin.Context)(parkLib model.Mirror,err error){
	_ = c.BindJSON(&parkLib)
	if err = config.DB.Model(&parkLib).Create(&parkLib).Error;err!=nil{
		return model.Mirror{}, err
	}
	return
}
func UpdateMirror(c *gin.Context,ID uint)(parkLib model.Mirror,err error){
	_ = c.BindJSON(&parkLib)
	if err = config.DB.Model(&parkLib).Where("id = ?",ID).Update(&parkLib).Error;err !=nil{
		return model.Mirror{}, err
	}
	return
}
func DeleteMirror(ID uint)(affe int64){
	affe =config.DB.Where("id = ?",ID).First(&model.Mirror{}).Delete(&model.Mirror{}).RowsAffected
	return
}

func GetAllApplications(c *gin.Context) (applicationComposer paginate.ApplicationComposer,err error) {
	page := c.Param("page")
	pagesize := c.Param("pageSize")
	var applications []model.Application
	SQL:=config.DB.Model(&model.Application{}).Preload("Mirrors")
	applicationComposer,err = paginate.ApplicationPaginator(SQL,page,pagesize, applications)
	return
}
func GetApplicationById(ID uint) (parkLib model.Application,err error) {
	if err =config.DB.Preload("Mirrors").Where("id = ?",ID).First(&parkLib).Error;err != nil{
		return model.Application{}, err
	}
	return
}
func CreateApplication(c *gin.Context)(application model.Application,err error){
	var applicationIds model.ApplicationIdList
	_ = c.BindJSON(&applicationIds)
	var mirrors []model.Mirror
	config.DB.Model(&model.Mirror{}).Where("id in (?)",applicationIds.Ids).Find(&mirrors)
	application.Mirrors = mirrors
	if err = config.DB.Model(&application).Create(&application).Error;err!=nil{
		return model.Application{}, err
	}
	return
}

func UpdateApplication(c *gin.Context,ID uint)(application model.Application,err error){
	//先创建一个变量准备接受自己定义的结构数据
	var applicationIdList model.ApplicationIdList
	//绑定请求参数
	_ = c.BindJSON(&applicationIdList)
	// 取出 mirror ID list
	var mirrorIdList = applicationIdList.Ids
	// 定义子属性接收数组
	var mirrors []model.Mirror

	// 取出子属性数组
	config.DB.Model(&model.Mirror{}).Where("id in (?)",mirrorIdList).Find(&mirrors)
	// 将列表关系数据绑定到新建application上
	applicationIdList.Application.ID = ID
	applicationIdList.Application.Mirrors = mirrors
	if err = config.DB.Model(&model.Application{}).Where("id = ?",ID).Update(&applicationIdList.Application).Error;err !=nil{
		return model.Application{}, err
	}
	// 更新关系
	config.DB.Model(&applicationIdList.Application.Mirrors).Update(&applicationIdList.Application.Mirrors)
	config.DB.Model(&applicationIdList.Application).Association("Mirrors").Replace(&mirrors)
	return applicationIdList.Application,err
}
func DeleteApplication(ID uint)(affe int64){
	affe =config.DB.Where("id = ?",ID).First(&model.Application{}).Delete(&model.Application{}).RowsAffected
	return
}