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
	c.JSON(200, gin.H{
		"role": role,
		"affe": affe,
	})
}

func UpdateRoleHandler(c *gin.Context) {
	id := c.Param("id")
	role, affe := user.UpdateRole(c, utils.StrToUInt(id))
	c.JSON(200, gin.H{
		"role": role,
		"affe": affe,
	})
}

func DeleteRoleHandler(c *gin.Context) {
	id := c.Param("id")
	affe := user.DeleteRole(utils.StrToUInt(id))
	c.JSON(200, gin.H{
		"affe": affe,
	})
}
