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
	if err != nil {
		c.JSON(500, Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(pers))
}

func PermissionByIdHandler(c *gin.Context) {
	ID := c.Param("id")
	pers, err := user.GetPermisById(utils.StrToUInt(ID))
	if err != nil {
		c.JSON(500, Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(pers))
}

func CreatePermissionHandler(c *gin.Context) {
	per,affe :=user.CreatePermission(c)
	if affe == 0 {
		c.JSON(500, Response.ResponseBody{}.FailRes("创建失败"))
		return
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(per))
}

func UpdatePerHandler(c *gin.Context) {
	var id = c.Param("id")
	per,affe:=user.UpdatePermission(c, utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500, Response.ResponseBody{}.FailRes("创建失败"))
		return
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(per))
}

func DeletePermissionHandler(c *gin.Context) {
	var id = c.Param("id")
	affe:=user.DeletePermission(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500, Response.ResponseBody{}.FailRes("删除失败"))
		return
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(affe))
}

func PermissionsTreeHandler(c *gin.Context){
	treeList, err :=user.GetPermissionTree()
	if err != nil {
		c.JSON(500, Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(treeList))
}

func GetPermissionFamilyHandler(c *gin.Context){
	id := c.Param("id")
	treeList, err :=user.GetPermissionFamily(utils.StrToUInt(id))
	if err != nil {
		c.JSON(500, Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(treeList))
}
// 获得用户权限树形结构
func GetUserPermissionTreeHandler(c *gin.Context){
	id := c.Param("id")
	treeList, err :=user.GetUserPermissionTreeById(utils.StrToUInt(id))
	if err != nil {
		c.JSON(404, Response.ResponseBody{}.NotFound())
		return
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(treeList))
}