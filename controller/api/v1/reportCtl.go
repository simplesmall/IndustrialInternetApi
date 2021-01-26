package v1

import (
	"IndustrialInternetApi/common/utils"
	Response "IndustrialInternetApi/model/response"
	"IndustrialInternetApi/service/independent"
	"github.com/gin-gonic/gin"
)

func GetAllMonthlyReportsHandler(c *gin.Context)  {
	parkLibs, err := independent.GetAllMonthlyReports()
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLibs))
}
func GetMonthlyReportByIdHandler(c *gin.Context)  {
	id := c.Param("id")
	parkLib,err:=independent.GetMonthlyReportById(utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLib))
}
func CreateMonthlyReportHandler(c *gin.Context)  {
	create,err:=independent.CreateMonthlyReport(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(create))
}
func UpdateMonthlyReportHandler(c *gin.Context)  {
	id := c.Param("id")
	update,err:=independent.UpdateMonthlyReport(c,utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(update))
}
func DeleteMonthlyReportHandler(c *gin.Context)  {
	id := c.Param("id")
	affe:=independent.DeleteMonthlyReport(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500,Response.ResponseBody{}.FailRes("删除失败"))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(affe))
}

func GetAllStatisticReportsHandler(c *gin.Context)  {
	parkLibs, err := independent.GetAllStatisticReports()
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLibs))
}
func GetStatisticReportByIdHandler(c *gin.Context)  {
	id := c.Param("id")
	parkLib,err:=independent.GetStatisticReportById(utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLib))
}
func CreateStatisticReportHandler(c *gin.Context)  {
	create,err:=independent.CreateStatisticReport(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(create))
}
func UpdateStatisticReportHandler(c *gin.Context)  {
	id := c.Param("id")
	update,err:=independent.UpdateStatisticReport(c,utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(update))
}
func DeleteStatisticReportHandler(c *gin.Context)  {
	id := c.Param("id")
	affe:=independent.DeleteStatisticReport(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500,Response.ResponseBody{}.FailRes("删除失败"))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(affe))
}