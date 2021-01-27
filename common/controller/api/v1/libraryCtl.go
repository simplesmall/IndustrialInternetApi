package v1

import (
	"IndustrialInternetApi/common/utils"
	Response "IndustrialInternetApi/model/response"
	"IndustrialInternetApi/service/independent"
	"github.com/gin-gonic/gin"
)

func GetAllEnterPrisesHandler(c *gin.Context)  {
	enterpriseLibs, err := independent.GetAllEnterPrises()
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(enterpriseLibs))
}
func GetEnterPriseByIdHandler(c *gin.Context)  {
	id := c.Param("id")
	enterpriseLib,err:=independent.GetEnterPriseById(utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(enterpriseLib))
}
func CreateEnterPriseHandler(c *gin.Context)  {
	create,err:=independent.CreateEnterPrise(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(create))
}
func UpdateEnterPriseHandler(c *gin.Context)  {
	id := c.Param("id")
	update,err:=independent.UpdateEnterPrise(c,utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(update))
}
func DeleteEnterPriseHandler(c *gin.Context)  {
	id := c.Param("id")
	affe:=independent.DeleteEnterPrise(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500,Response.ResponseBody{}.FailRes("删除失败"))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(affe))
}

func GetAllProjectsHandler(c *gin.Context)  {
	projects, err := independent.GetAllProjects()
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(projects))
}
func GetProjectByIdHandler(c *gin.Context)  {
	id := c.Param("id")
	project,err:=independent.GetProjectById(utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(project))
}
func CreateProjectHandler(c *gin.Context)  {
	create,err:=independent.CreateProject(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(create))
}
func UpdateProjectHandler(c *gin.Context)  {
	id := c.Param("id")
	update,err:=independent.UpdateProject(c,utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(update))
}
func DeleteProjectHandler(c *gin.Context)  {
	id := c.Param("id")
	affe:=independent.DeleteProject(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500,Response.ResponseBody{}.FailRes("删除失败"))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(affe))
}


func GetAllParksHandler(c *gin.Context)  {
	parkLibs, err := independent.GetAllParks()
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLibs))
}
func GetParkByIdHandler(c *gin.Context)  {
	id := c.Param("id")
	parkLib,err:=independent.GetParkById(utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLib))
}
func CreateParkHandler(c *gin.Context)  {
	create,err:=independent.CreatePark(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(create))
}
func UpdateParkHandler(c *gin.Context)  {
	id := c.Param("id")
	update,err:=independent.UpdatePark(c,utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(update))
}
func DeleteParkHandler(c *gin.Context)  {
	id := c.Param("id")
	affe:=independent.DeletePark(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500,Response.ResponseBody{}.FailRes("删除失败"))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(affe))
}
