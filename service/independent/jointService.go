package independent

import (
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/model"
	"IndustrialInternetApi/model/paginate"
	"github.com/gin-gonic/gin"
)

func GetAllJoints(c *gin.Context) (jointComposer paginate.JointComposer,err error) {
	page := c.Param("page")
	pagesize := c.Param("pageSize")
	var parkLibs []model.Joint
	SQL:=config.DB.Model(&model.Joint{})
	jointComposer,err = paginate.JointPaginator(SQL,page,pagesize, parkLibs)
	return
}
func GetJointById(ID uint) (joint model.Joint,err error) {
	if err =config.DB.Where("id = ?",ID).First(&joint).Error;err != nil{
		return model.Joint{}, err
	}
	return
}
func CreateJoint(c *gin.Context)(joint model.Joint,err error){
	_ = c.BindJSON(&joint)
	if err = config.DB.Model(&joint).Create(&joint).Error;err!=nil{
		return model.Joint{}, err
	}
	return
}
func UpdateJoint(c *gin.Context,ID uint)(joint model.Joint,err error){
	_ = c.BindJSON(&joint)
	if err = config.DB.Model(&joint).Where("id = ?",ID).Update(&joint).Error;err !=nil{
		return model.Joint{}, err
	}
	return
}
func DeleteJoint(ID uint)(affe int64){
	affe =config.DB.Where("id = ?",ID).First(&model.Joint{}).Delete(&model.Joint{}).RowsAffected
	return
}
