package v1

import (
	"IndustrialInternetApi/common/utils"
	Response "IndustrialInternetApi/model/response"
	"IndustrialInternetApi/service/independent"
	"github.com/gin-gonic/gin"
)

func GetAllApplicationInsertsHandler(c *gin.Context)  {
	parkLibs, err := independent.GetAllApplicationInserts(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLibs))
}
func GetApplicationInsertByIdHandler(c *gin.Context)  {
	id := c.Param("id")
	parkLib,err:=independent.GetApplicationInsertById(utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLib))
}
func CreateApplicationInsertHandler(c *gin.Context)  {
	create,err:=independent.CreateApplicationInsert(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(create))
}
func UpdateApplicationInsertHandler(c *gin.Context)  {
	id := c.Param("id")
	update,err:=independent.UpdateApplicationInsert(c,utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(update))
}
func DeleteApplicationInsertHandler(c *gin.Context)  {
	id := c.Param("id")
	affe:=independent.DeleteApplicationInsert(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500,Response.ResponseBody{}.FailRes("删除失败"))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(affe))
}

func GetAllMyInsertsHandler(c *gin.Context)  {
	myInserts, err := independent.GetAllMyInserts(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(myInserts))
}
func GetMyInsertByIdHandler(c *gin.Context)  {
	id := c.Param("id")
	myInsertById,err:=independent.GetMyInsertById(utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(myInsertById))
}
func CreateMyInsertHandler(c *gin.Context)  {
	create,err:=independent.CreateMyInsert(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(create))
}
func UpdateMyInsertHandler(c *gin.Context)  {
	id := c.Param("id")
	update,err:=independent.UpdateMyInsert(c,utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(update))
}
func DeleteMyInsertHandler(c *gin.Context)  {
	id := c.Param("id")
	affe:=independent.DeleteMyInsert(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500,Response.ResponseBody{}.FailRes("删除失败"))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(affe))
}

func GetAllMyDevicesHandler(c *gin.Context)  {
	myDevices, err := independent.GetAllMyDevices()
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(myDevices))
}
func GetMyDeviceByIdHandler(c *gin.Context)  {
	id := c.Param("id")
	myDeviceById,err:=independent.GetMyDeviceById(utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(myDeviceById))
}
func CreateMyDeviceHandler(c *gin.Context)  {
	myDevice,err:=independent.CreateMyDevice(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(myDevice))
}
func UpdateMyDeviceHandler(c *gin.Context)  {
	id := c.Param("id")
	myDevice,err:=independent.UpdateMyDevice(c,utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(myDevice))
}
func DeleteMyDeviceHandler(c *gin.Context)  {
	id := c.Param("id")
	affe:=independent.DeleteMyDevice(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500,Response.ResponseBody{}.FailRes("删除失败"))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(affe))
}
