package independent

import (
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/model"
	"github.com/gin-gonic/gin"
)

func GetAllEnterPrises() (enterpriseLibs []model.EnterpriseLib,err error) {
	if err = config.DB.Find(&enterpriseLibs).Error;err !=nil{
		return nil, err
	}
	return
}
func GetEnterPriseById(ID uint) (enterpriseLib model.EnterpriseLib,err error) {
	if err =config.DB.Where("id = ?",ID).First(&enterpriseLib).Error;err != nil{
		return model.EnterpriseLib{}, err
	}
	return
}

func GetAllProjects() (projectLibs []model.ProjectLib,err error) {
	if err = config.DB.Find(&projectLibs).Error;err !=nil{
		return nil, err
	}
	return
}
func GetProjectById(ID uint) (projectLib model.ProjectLib,err error) {
	if err =config.DB.Where("id = ?",ID).First(&projectLib).Error;err != nil{
		return model.ProjectLib{}, err
	}
	return
}

func CreateEnterPrise(c *gin.Context)(enterprise model.EnterpriseLib,err error){
	_ = c.BindJSON(&enterprise)
	if err = config.DB.Model(&enterprise).Create(&enterprise).Error;err!=nil{
		return model.EnterpriseLib{}, err
	}
	return
}
func UpdateEnterPrise(c *gin.Context,ID uint)(enterprise model.EnterpriseLib,err error){
	_ = c.BindJSON(&enterprise)
	if err = config.DB.Model(&enterprise).Where("id = ?",ID).Update(&enterprise).Error;err !=nil{
		return model.EnterpriseLib{}, err
	}
	return
}
func DeleteEnterPrise(ID uint)(affe int64){
	affe =config.DB.Model(&model.EnterpriseLib{}).Where("id = ?",ID).First(&model.EnterpriseLib{}).Delete(&model.EnterpriseLib{}).RowsAffected
	return
}

func CreateProject(c *gin.Context)(projectLib model.ProjectLib,err error){
	_ = c.BindJSON(&projectLib)
	if err = config.DB.Model(&projectLib).Create(&projectLib).Error;err!=nil{
		return model.ProjectLib{}, err
	}
	return
}
func UpdateProject(c *gin.Context,ID uint)(projectLib model.ProjectLib,err error){
	_ = c.BindJSON(&projectLib)
	if err = config.DB.Model(&projectLib).Where("id = ?",ID).Update(&projectLib).Error;err !=nil{
		return model.ProjectLib{}, err
	}
	return
}
func DeleteProject(ID uint)(affe int64){
	affe =config.DB.Where("id = ?",ID).First(&model.ProjectLib{}).Delete(&model.ProjectLib{}).RowsAffected
	return
}

func GetAllParks() (parkLibs []model.ParkLib,err error) {
	if err = config.DB.Find(&parkLibs).Error;err !=nil{
		return nil, err
	}
	return
}
func GetParkById(ID uint) (parkLib model.ParkLib,err error) {
	if err =config.DB.Where("id = ?",ID).First(&parkLib).Error;err != nil{
		return model.ParkLib{}, err
	}
	return
}
func CreatePark(c *gin.Context)(parkLib model.ParkLib,err error){
	_ = c.BindJSON(&parkLib)
	if err = config.DB.Model(&parkLib).Create(&parkLib).Error;err!=nil{
		return model.ParkLib{}, err
	}
	return
}
func UpdatePark(c *gin.Context,ID uint)(parkLib model.ParkLib,err error){
	_ = c.BindJSON(&parkLib)
	if err = config.DB.Model(&parkLib).Where("id = ?",ID).Update(&parkLib).Error;err !=nil{
		return model.ParkLib{}, err
	}
	return
}
func DeletePark(ID uint)(affe int64){
	affe =config.DB.Where("id = ?",ID).First(&model.ParkLib{}).Delete(&model.ParkLib{}).RowsAffected
	return
}