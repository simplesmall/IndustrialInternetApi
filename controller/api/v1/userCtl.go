package v1

import (
	"IndustrialInternetApi/common/utils"
	Response "IndustrialInternetApi/model/response"
	"IndustrialInternetApi/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func errorResult(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response.ResponseBody{}.FailRes(err))
		return
	}
}

func UsersHandler(c *gin.Context) {
	users, err := user.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(users))
}

func UserByIdHandler(c *gin.Context) {
	ID := c.Param("id")
	userById, err := user.GetUserById(utils.StrToUInt(ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response.ResponseBody{}.FailRes(err))
		return
	}
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(userById))
}

func CreateUserHandler(c *gin.Context) {
	createUser, affe := user.CreateUser(c)
	if affe!=1{
		c.JSON(500,Response.ResponseBody{}.FailRes("创建失败"))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(createUser))
}

func UpdateUserHandler(c *gin.Context) {
	id:=c.Param("id")
	createUser, affe := user.UpdateUser(c,utils.StrToUInt(id))
	if affe!=1{
		c.JSON(500,Response.ResponseBody{}.FailRes("更新失败"))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(createUser))
}

func DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")
	affe := user.DeleteUser(utils.StrToUInt(id))
	if affe!=1{
		c.JSON(500,Response.ResponseBody{}.FailRes("删除失败"))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(affe))
}

