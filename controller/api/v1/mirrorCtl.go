package v1

import (
	"IndustrialInternetApi/common/utils"
	Response "IndustrialInternetApi/model/response"
	"IndustrialInternetApi/service/independent"
	"github.com/gin-gonic/gin"
)

func GetAllMirrorsHandler(c *gin.Context)  {
	mirrors, err := independent.GetAllMirrors(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(mirrors))
}
func GetMirrorByIdHandler(c *gin.Context)  {
	id := c.Param("id")
	parkLib,err:=independent.GetMirrorById(utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLib))
}
func CreateMirrorHandler(c *gin.Context)  {
	create,err:=independent.CreateMirror(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(create))
}
func UpdateMirrorHandler(c *gin.Context)  {
	id := c.Param("id")
	update,err:=independent.UpdateMirror(c,utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(update))
}
func DeleteMirrorHandler(c *gin.Context)  {
	id := c.Param("id")
	affe:=independent.DeleteMirror(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500,Response.ResponseBody{}.FailRes("删除失败"))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(affe))
}

func GetAllApplicationsHandler(c *gin.Context)  {
	parkLibs, err := independent.GetAllApplications(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLibs))
}
func GetApplicationByIdHandler(c *gin.Context)  {
	id := c.Param("id")
	parkLib,err:=independent.GetApplicationById(utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLib))
}
func CreateApplicationHandler(c *gin.Context)  {
	create,err:=independent.CreateApplication(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(create))
}
func UpdateApplicationHandler(c *gin.Context)  {
	id := c.Param("id")
	update,err:=independent.UpdateApplication(c,utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(update))
}
func DeleteApplicationHandler(c *gin.Context)  {
	id := c.Param("id")
	affe:=independent.DeleteApplication(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500,Response.ResponseBody{}.FailRes("删除失败"))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(affe))
}

