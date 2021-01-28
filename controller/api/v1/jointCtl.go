package v1

import (
	"IndustrialInternetApi/common/utils"
	Response "IndustrialInternetApi/model/response"
	"IndustrialInternetApi/service/independent"
	"github.com/gin-gonic/gin"
)

func GetAllJointsHandler(c *gin.Context)  {
	joints, err := independent.GetAllJoints(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(joints))
}
func GetJointByIdHandler(c *gin.Context)  {
	id := c.Param("id")
	parkLib,err:=independent.GetJointById(utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLib))
}
func CreateJointHandler(c *gin.Context)  {
	create,err:=independent.CreateJoint(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(create))
}
func UpdateJointHandler(c *gin.Context)  {
	id := c.Param("id")
	update,err:=independent.UpdateJoint(c,utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(update))
}
func DeleteJointHandler(c *gin.Context)  {
	id := c.Param("id")
	affe:=independent.DeleteJoint(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500,Response.ResponseBody{}.FailRes("删除失败"))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(affe))
}