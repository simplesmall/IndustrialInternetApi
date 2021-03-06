package independent

import (
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/model"
	"IndustrialInternetApi/model/paginate"
	"github.com/gin-gonic/gin"
)

func GetAllApplicationInserts(c *gin.Context) (applicationInsertComposer paginate.ApplicationInsertComposer, err error) {
	page := c.Param("page")
	pagesize := c.Param("pageSize")

	var inserts []model.ApplicationInsert
	SQL := config.DB.Model(&model.ApplicationInsert{})
	applicationInsertComposer, err = paginate.ApplicationInsertPaginator(SQL, page, pagesize, inserts)
	if err = config.DB.Find(&inserts).Error; err != nil {
		return applicationInsertComposer, err
	}
	return
}
func GetApplicationInsertById(ID uint) (applicationInsert model.ApplicationInsert, err error) {
	if err = config.DB.Where("id = ?", ID).First(&applicationInsert).Error; err != nil {
		return model.ApplicationInsert{}, err
	}
	return
}
func CreateApplicationInsert(c *gin.Context) (applicationInsert model.ApplicationInsert, err error) {
	_ = c.BindJSON(&applicationInsert)
	if err = config.DB.Model(&applicationInsert).Create(&applicationInsert).Error; err != nil {
		return model.ApplicationInsert{}, err
	}
	return
}
func UpdateApplicationInsert(c *gin.Context, ID uint) (applicationInsert model.ApplicationInsert, err error) {
	_ = c.BindJSON(&applicationInsert)
	if err = config.DB.Model(&applicationInsert).Where("id = ?", ID).Update(&applicationInsert).Error; err != nil {
		return model.ApplicationInsert{}, err
	}
	return
}
func DeleteApplicationInsert(ID uint) (affe int64) {
	affe = config.DB.Where("id = ?", ID).First(&model.ApplicationInsert{}).Delete(&model.ApplicationInsert{}).RowsAffected
	return
}

func GetAllMyInserts(c *gin.Context) (myInsertComposer paginate.MyInsertComposer, err error) {
	page := c.Param("page")
	pageSize := c.Param("pageSize")

	var inserts []model.MyInsert
	SQL := config.DB.Model(&model.MyInsert{})
	myInsertComposer, err = paginate.MyInsertPaginator(SQL, page, pageSize, inserts)
	if err = config.DB.Find(&inserts).Error; err != nil {
		return myInsertComposer, err
	}
	return
}
func GetMyInsertById(ID uint) (myInsert model.MyInsert, err error) {
	if err = config.DB.Where("id = ?", ID).First(&myInsert).Error; err != nil {
		return model.MyInsert{}, err
	}
	return
}
func CreateMyInsert(c *gin.Context) (myInsert model.MyInsert, err error) {
	_ = c.BindJSON(&myInsert)
	if err = config.DB.Model(&myInsert).Create(&myInsert).Error; err != nil {
		return model.MyInsert{}, err
	}
	return
}
func UpdateMyInsert(c *gin.Context, ID uint) (myInsert model.MyInsert, err error) {
	_ = c.BindJSON(&myInsert)
	if err = config.DB.Model(&myInsert).Where("id = ?", ID).Update(&myInsert).Error; err != nil {
		return model.MyInsert{}, err
	}
	return
}
func DeleteMyInsert(ID uint) (affe int64) {
	affe = config.DB.Where("id = ?", ID).First(&model.MyInsert{}).Delete(&model.MyInsert{}).RowsAffected
	return
}

func GetAllMyDevices(c *gin.Context) (myDeviceComposer paginate.MyDeviceComposer, err error) {
	page := c.Param("page")
	pagesize := c.Param("pageSize")
	var devices []model.MyDevice
	SQL := config.DB.Model(&model.MyDevice{}).Preload("MyInserts")
	myDeviceComposer, err = paginate.MyDevicePaginator(SQL, page, pagesize, devices)
	if err = config.DB.Preload("MyInserts").Find(&devices).Error; err != nil {
		return myDeviceComposer, err
	}
	return
}
func GetMyDeviceById(ID uint) (myDevice model.MyDevice, err error) {
	if err = config.DB.Where("id = ?", ID).Preload("MyInserts").First(&myDevice).Error; err != nil {
		return model.MyDevice{}, err
	}
	return
}
func CreateMyDevice(c *gin.Context) (myDevice model.MyDevice, err error) {
	var mydeviceIdList model.MyDeviceIdList
	_ = c.BindJSON(&mydeviceIdList)
	myinsertId := mydeviceIdList.Ids
	var myinserts []model.MyInsert
	config.DB.Model(&model.MyInsert{}).Where("id in (?)", myinsertId).Find(&myinserts)
	var mydevice = mydeviceIdList.MyDevice
	mydevice.MyInserts = myinserts
	if err = config.DB.Model(&model.MyDevice{}).Create(&mydevice).Error; err != nil {
		return model.MyDevice{}, err
	}
	return mydevice, err
}
func UpdateMyDevice(c *gin.Context, ID uint) (myDevice model.MyDevice, err error) {
	var mydeviceIdList model.MyDeviceIdList
	_ = c.BindJSON(&mydeviceIdList)
	myinsertId := mydeviceIdList.Ids

	var myInserts []model.MyInsert
	config.DB.Model(&model.MyInsert{}).Where("id in (?)", myinsertId).Find(&myInserts)
	mydeviceIdList.MyDevice.ID = ID
	mydeviceIdList.MyDevice.MyInserts = myInserts
	if err = config.DB.Model(&model.MyDevice{}).Where("id = ?", ID).Update(&mydeviceIdList.MyDevice).Error; err != nil {
		return model.MyDevice{}, err
	}
	config.DB.Model(&mydeviceIdList.MyDevice.MyInserts).Update(&mydeviceIdList.MyDevice.MyInserts)
	config.DB.Model(&mydeviceIdList.MyDevice).Association("MyInserts").Replace(&myInserts)
	return mydeviceIdList.MyDevice, err
}
func DeleteMyDevice(ID uint) (affe int64) {
	affe = config.DB.Where("id = ?", ID).First(&model.MyDevice{}).Delete(&model.MyDevice{}).RowsAffected
	return
}
