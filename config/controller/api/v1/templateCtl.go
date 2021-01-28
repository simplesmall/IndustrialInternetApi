package v1

import (
	"IndustrialInternetApi/common/utils"
	Response "IndustrialInternetApi/model/response"
	"IndustrialInternetApi/service/independent"
	"github.com/gin-gonic/gin"
)

func GetAllModelsHandler(c *gin.Context)  {
	parkLibs, err := independent.GetAllModels(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLibs))
}
func GetModelByIdHandler(c *gin.Context)  {
	id := c.Param("id")
	parkLib,err:=independent.GetModelById(utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLib))
}
func CreateModelHandler(c *gin.Context)  {
	create,err:=independent.CreateModel(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(create))
}
func UpdateModelHandler(c *gin.Context)  {
	id := c.Param("id")
	update,err:=independent.UpdateModel(c,utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(update))
}
func DeleteModelHandler(c *gin.Context)  {
	id := c.Param("id")
	affe:=independent.DeleteModel(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500,Response.ResponseBody{}.FailRes("删除失败"))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(affe))
}

func GetAllModulesHandler(c *gin.Context)  {
	parkLibs, err := independent.GetAllModules(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLibs))
}
func GetModuleByIdHandler(c *gin.Context)  {
	id := c.Param("id")
	parkLib,err:=independent.GetModuleById(utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(parkLib))
}
func CreateModuleHandler(c *gin.Context)  {
	create,err:=independent.CreateModule(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(create))
}
func UpdateModuleHandler(c *gin.Context)  {
	id := c.Param("id")
	update,err:=independent.UpdateModule(c,utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(update))
}
func DeleteModuleHandler(c *gin.Context)  {
	id := c.Param("id")
	affe:=independent.DeleteModule(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500,Response.ResponseBody{}.FailRes("删除失败"))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(affe))
}

func GetAllTemplatesHandler(c *gin.Context)  {
	templates, err := independent.GetAllTemplates(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(templates))
}
func GetTemplateByIdHandler(c *gin.Context)  {
	id := c.Param("id")
	templateById,err:=independent.GetTemplateById(utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(templateById))
}
func CreateTemplateHandler(c *gin.Context)  {
	create,err:=independent.CreateTemplate(c)
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(create))
}
func UpdateTemplateHandler(c *gin.Context)  {
	id := c.Param("id")
	update,err:=independent.UpdateTemplate(c,utils.StrToUInt(id))
	if err != nil {
		c.JSON(500,Response.ResponseBody{}.FailRes(err))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(update))
}
func DeleteTemplateHandler(c *gin.Context)  {
	id := c.Param("id")
	affe:=independent.DeleteTemplate(utils.StrToUInt(id))
	if affe == 0 {
		c.JSON(500,Response.ResponseBody{}.FailRes("删除失败"))
	}
	c.JSON(200,Response.ResponseBody{}.OKResult(affe))
}
