package independent

import (
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/model"
	"IndustrialInternetApi/model/paginate"
	"IndustrialInternetApi/service/user"
	"github.com/gin-gonic/gin"
)

func GetAllModels(c *gin.Context) (composer paginate.ModelComposer,err error) {
	page := c.Param("page")
	pagesize := c.Param("pageSize")
	var userList []model.Model
	SQL:=config.DB.Model(&model.Model{})

	composer,err = paginate.ModelPaginator(SQL,page,pagesize,userList)
	return composer,err
}
func GetModelById(ID uint) (lib model.Model,err error) {
	if err =config.DB.Where("id = ?",ID).First(&lib).Error;err != nil{
		return model.Model{}, err
	}
	return
}
func CreateModel(c *gin.Context)(model model.Model,err error){
	_ = c.BindJSON(&model)
	loginUser,_:=user.GetLoginUser(c)
	model.Owner = loginUser.User.Account
	if err = config.DB.Model(&model).Create(&model).Error;err!=nil{
		return model, err
	}
	return
}
func UpdateModel(c *gin.Context,ID uint)(lib model.Model,err error){
	_ = c.BindJSON(&lib)
	if err = config.DB.Model(&lib).Where("id = ?",ID).Update(&lib).Error;err !=nil{
		return model.Model{}, err
	}
	return
}
func DeleteModel(ID uint)(affe int64){
	affe =config.DB.Where("id = ?",ID).First(&model.Model{}).Delete(&model.Model{}).RowsAffected
	return
}

func GetAllModules(c *gin.Context) (composer paginate.ModuleComposer,err error) {
	page := c.Param("page")
	pagesize := c.Param("pageSize")
	var modules []model.Module
	SQL:=config.DB.Model(&model.Module{}).Preload("Models")
	composer,err = paginate.ModulePaginator(SQL,page,pagesize, modules)
	return
}
func GetModuleById(ID uint) (parkLib model.Module,err error) {
	if err =config.DB.Preload("Models").Where("id = ?",ID).First(&parkLib).Error;err != nil{
		return model.Module{}, err
	}
	return
}
func CreateModule(c *gin.Context)(Module model.Module,err error){
	var module model.Module
	var moduleIdList model.ModuleIdList
	_ = c.BindJSON(&moduleIdList)
	modelId := moduleIdList.Ids
	module = moduleIdList.Module
	var models []model.Model
	config.DB.Model(&model.Model{}).Where("id in (?)",modelId).Find(&models)
	loginUser,_:=user.GetLoginUser(c)
	module.Owner = loginUser.User.Account
	module.Models = models

	if err = config.DB.Model(&module).Create(&module).Error;err!=nil{
		return model.Module{}, err
	}
	return module,err
}
func UpdateModule(c *gin.Context,ID uint)(module model.Module,err error){
	var moduleIdList model.ModuleIdList
	_ = c.BindJSON(&moduleIdList)
	modelId := moduleIdList.Ids

	var models []model.Model
	config.DB.Model(&model.Model{}).Where("id in (?)",modelId).Find(&models)

	moduleIdList.Module.ID = ID
	moduleIdList.Module.Models = models

	if err = config.DB.Model(&model.Module{}).Where("id = ?",ID).Update(&moduleIdList.Module).Error;err !=nil{
		return model.Module{}, err
	}

	config.DB.Model(&moduleIdList.Module.Model).Update(&moduleIdList.Module.Model)
	config.DB.Model(&moduleIdList.Module).Association("Models").Replace(&models)
	return	moduleIdList.Module,err
}
func DeleteModule(ID uint)(affe int64){
	affe =config.DB.Where("id = ?",ID).First(&model.Module{}).Delete(&model.Module{}).RowsAffected
	return
}

func GetAllTemplates(c *gin.Context) (templateComposer paginate.TemplateComposer,err error) {
	page := c.Param("page")
	pagesize := c.Param("pageSize")
	var applications []model.Template
	SQL:=config.DB.Model(&model.Template{}).Preload("Modules")
	templateComposer,err = paginate.TemplatePaginator(SQL,page,pagesize, applications)
	return
}
func GetTemplateById(ID uint) (parkLib model.Template,err error) {
	if err =config.DB.Preload("Modules").Preload("Modules.Models").Where("id = ?",ID).First(&parkLib).Error;err != nil{
		return model.Template{}, err
	}
	return
}
func CreateTemplate(c *gin.Context)(Template model.Template,err error){
	var template model.Template
	var templateIdList model.TemplateIdList
	_ = c.BindJSON(&templateIdList)
	moduleId := templateIdList.Ids
	template = templateIdList.Template
	var modules []model.Module
	config.DB.Model(&model.Module{}).Where("id in (?)",moduleId).Find(&modules)
	template.Modules = modules
	loginUser,_:=user.GetLoginUser(c)
	template.Owner = loginUser.User.Account

	if err = config.DB.Model(&template).Create(&template).Error;err!=nil{
		return model.Template{}, err
	}
	return template,err
}
func UpdateTemplate(c *gin.Context,ID uint)(Template model.Template,err error){
	var templateIdList model.TemplateIdList
	_ = c.BindJSON(&templateIdList)

	moduleId := templateIdList.Ids

	var modules []model.Module
	config.DB.Model(&model.Module{}).Where("id in (?)",moduleId).Find(&modules)

	templateIdList.Template.ID = ID
	templateIdList.Template.Modules = modules

	if err = config.DB.Model(&model.Template{}).Where("id = ?",ID).Update(&templateIdList.Template).Error;err !=nil{
		return model.Template{}, err
	}

	config.DB.Model(&templateIdList.Template.Modules).Update(&templateIdList.Template.Modules)
	config.DB.Model(&templateIdList.Template).Association("Modules").Replace(&modules)
	return templateIdList.Template,err
}
func DeleteTemplate(ID uint)(affe int64){
	affe =config.DB.Where("id = ?",ID).First(&model.Template{}).Delete(&model.Template{}).RowsAffected
	return
}