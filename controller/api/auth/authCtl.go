package auth

import (
	"IndustrialInternetApi/common/middleware"
	"IndustrialInternetApi/model"
	Response "IndustrialInternetApi/model/response"
	"errors"
	"github.com/gin-gonic/gin"
	jwt "IndustrialInternetApi/service/jwt"
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
	err := c.ShouldBindJSON(&loginInput)
	if err != nil {
		Response.ResponseBody{}.FailRes(errors.New("数据获取失败"))
	}

	//参数校验
	token, err := model.Login(loginInput.Username, loginInput.Password)
	if err != nil {
		if err.Error() == "record not found" {
			Response.ResponseBody{}.FailRes("该用户不存在")
			return
		} else if err.Error() == "invalid password" {
			Response.ResponseBody{}.FailRes("密码校验不通过")
			return
		} else {
			Response.ResponseBody{}.FailRes("登录错误")
			return
		}
	}
	Response.ResponseBody{}.OKResultWithMsg(token,"登录成功")
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
	Response.ResponseBody{}.OKResult("注销成功")
}

//更新token
func InitiativeExpireHandler(c *gin.Context) {
	j := jwt.NewJWT()
	newToken, err := j.RefreshToken(middleware.AuthToken)
	if err != nil {
		Response.ResponseBody{}.FailRes("token更新失败")
		return
	}
	Response.ResponseBody{}.OKResultWithMsg(newToken,"更新成功")
}