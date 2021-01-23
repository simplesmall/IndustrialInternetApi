package v1

import (
	"IndustrialInternetApi/common/utils"
	Response "IndustrialInternetApi/model/response"
	"IndustrialInternetApi/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PermissionsHandler(c *gin.Context) {
	pers, err := user.GetAllPermissions()
	errorResult(c, err)
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(pers))
}

func PermissionByIdHandler(c *gin.Context) {
	ID := c.Param("id")
	pers, err := user.GetPermisById(utils.StrToUInt(ID))
	if err != nil {
		c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult("请求正常,数据错误"))
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(pers))
}

func CreatePermissionHandler(c *gin.Context) {
	per,affect :=user.CreatePermission(c)
	c.JSON(200,gin.H{
		"permission":per,
		"affect":affect,
	})
}

func UpdatePerHandler(c *gin.Context) {
	var id = c.Param("id")
	per,affe:=user.UpdatePermission(c, utils.StrToUInt(id))
	c.JSON(200,gin.H{
		"per":per,
		"aff":affe,
	})
}

func DeletePermissionHandler(c *gin.Context) {
	var id = c.Param("id")
	affe:=user.DeletePermission(utils.StrToUInt(id))
	c.JSON(200,gin.H{
		"affe":affe,
	})
}

func PermissionsTreeHandler(c *gin.Context){
	treeList, _ :=user.GetPermissionTree()
	c.JSON(200,Response.ResponseBody{}.OKResult(treeList))
}

func GetPermissionFamilyHandler(c *gin.Context){
	id := c.Param("pid")
	treeList, _ :=user.GetPermissionFamily(utils.StrToUInt(id))
	c.JSON(200,Response.ResponseBody{}.OKResult(treeList))
}