package auth

import (
	"IndustrialInternetApi/common/middleware"
	"IndustrialInternetApi/model"
	Response "IndustrialInternetApi/model/response"
	jwt "IndustrialInternetApi/service/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

// @Summary 用户登录
// @Description 用户登录接口
// @Tags 鉴权相关接口
// @Accept application/json
// @Produce application/json
// @Param param body Models.LoginInput true "登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /auth/login [post]
func LoginHandler(c *gin.Context) {
	var loginInput model.LoginInput
	err1 := c.ShouldBindJSON(&loginInput)
	if err1 != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(errors.New("数据获取失败")))
	}

	//参数校验
	token, err := model.Login(loginInput.Username, loginInput.Password)
	fmt.Println(loginInput,token,err)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(500,Response.ResponseBody{}.FailRes("该用户不存在"))
			return
		} else if err.Error() == "invalid password" {
			c.JSON(500,Response.ResponseBody{}.FailRes("密码校验不通过"))
			return
		} else {
			c.JSON(500,Response.ResponseBody{}.FailRes("登录错误"))
			return
		}
	}
	c.JSON(200,Response.ResponseBody{}.OKResultWithMsg("登录成功",token))
	return
}

// @Summary 注销登录
// @Description 注销登录接口
// @Tags 鉴权相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /auth/logout [get]
func LogoutHandler(c *gin.Context) {
	c.JSON(200,Response.ResponseBody{}.OKResult("注销成功"))
}

//更新token
func InitiativeExpireHandler(c *gin.Context) {
	j := jwt.NewJWT()
	middleware.AuthToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEsInVzZXJuYW1lIjoicm9vdCIsInBhc3N3b3JkIjoiJDJhJDEwJC40bTYvU0VKRmM2c3A5bjQ5MlhJQ2VCMmlBZWkubmJNeVh1ODBocms4TnhYcW9jS0pCbEQyIiwiZXhwIjoxNjExNDczMjY5LCJpc3MiOiJJbmR1c3RyaWFsSW50ZXJuZXRBcGkiLCJuYmYiOjE2MTE0NjUwNjl9.n4nFkUQ2PuSX_in9jVt21ywrd6fLNl-JGhPQCqU5G3I"
	newToken, err := j.RefreshToken(middleware.AuthToken)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes("token更新失败"))
		return
	}
	c.JSON(200,Response.ResponseBody{}.OKResultWithMsg("更新成功",newToken))
}