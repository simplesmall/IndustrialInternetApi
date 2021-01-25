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
	userById, _ := user.GetUserById(utils.StrToUInt(ID))
	//errorResult(c, err)
	c.JSON(http.StatusOK, Response.ResponseBody{}.OKResult(userById))
}

func CreateUserHandler(c *gin.Context) {
	createUser, affe := user.CreateUser(c)
	c.JSON(200,gin.H{
		"createUser":createUser,
		"affe":affe,
	})
}

func UpdateUserHandler(c *gin.Context) {
	id:=c.Param("id")
	createUser, affe := user.UpdateUser(c,utils.StrToUInt(id))
	c.JSON(200,gin.H{
		"updateUser":createUser,
		"affe":affe,
	})
}

func DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")
	affe := user.DeleteUser(utils.StrToUInt(id))
	c.JSON(200, gin.H{
		"affe": affe,
	})
}

