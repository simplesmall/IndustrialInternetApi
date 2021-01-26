package v1

import (
	"IndustrialInternetApi/common/utils"
	Response "IndustrialInternetApi/model/response"
	"IndustrialInternetApi/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RolesHandler(c *gin.Context) {
	roles, err := user.GetAllRoles(c)
	if err != nil {
		c.JSON(500, Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(roles))
}

func RoleByIdHandler(c *gin.Context) {
	ID := c.Param("id")
	roles, err := user.GetRoleById(utils.StrToUInt(ID))
	if err != nil {
		c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(err))
		return
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(roles))
}

func CreateRoleHandler(c *gin.Context) {
	role, affe := user.CreateRole(c)
	if affe == 0 {
		c.JSON(500, Response.ResponseBody{}.FailRes("创建成功"))
		return
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(role))
}

func UpdateRoleHandler(c *gin.Context) {
	id := c.Param("id")
	role, affe := user.UpdateRole(c, utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500, Response.ResponseBody{}.FailRes("更新失败"))
		return
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(role))
}

func DeleteRoleHandler(c *gin.Context) {
	id := c.Param("id")
	affe := user.DeleteRole(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500, Response.ResponseBody{}.FailRes("删除失败"))
		return
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult("删除成功"))
}
