package paginate

import (
	"IndustrialInternetApi/common/utils"
	"IndustrialInternetApi/model"
	"github.com/jinzhu/gorm"
)

type Pagination struct {
	Ok bool 		  `json:"ok"`
	Size  int        `form:"size" json:"size"`
	Page  int        `form:"page" json:"page"`
	Total int64        `json:"total"`
}

type UserComposer struct {
	Pagination
	Data []model.User `json:"data"`
}

func UserPaginator(sql *gorm.DB,page,size string, users []model.User) (data UserComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Offset((intPage-1)*intSize).Find(&users)
	if err = composeSQL.Error; err != nil {
		data.Ok = false
		return data,err
	}
	affected := composeSQL.RowsAffected
	data.Data = users
	data.Ok = true
	data.Size = intSize
	data.Page = intPage
	data.Total = affected
	return data,err
}

type RoleComposer struct {
	Pagination
	Data []model.Role `json:"data"`
}

func RolePaginator(sql *gorm.DB,page,size string, roles []model.Role) (data RoleComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Offset((intPage-1)*intSize).Find(&roles)
	if err = composeSQL.Error; err != nil {
		data.Ok = false
		return data,err
	}
	affected := composeSQL.RowsAffected
	data.Data = roles
	data.Ok = true
	data.Size = intSize
	data.Page = intPage
	data.Total = affected
	return data,err
}

type MirrorComposer struct {
	Pagination
	Data []model.Mirror `json:"data"`
}

func MirrorPaginator(sql *gorm.DB,page,size string, roles []model.Mirror) (data MirrorComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Offset((intPage-1)*intSize).Find(&roles)
	if err = composeSQL.Error; err != nil {
		data.Ok = false
		return data,err
	}
	affected := composeSQL.RowsAffected
	data.Data = roles
	data.Ok = true
	data.Size = intSize
	data.Page = intPage
	data.Total = affected
	return data,err
}

type ApplicationComposer struct {
	Pagination
	Data []model.Application `json:"data"`
}

func ApplicationPaginator(sql *gorm.DB,page,size string, applications []model.Application) (applicationComposer ApplicationComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Preload("Mirrors").Offset((intPage-1)*intSize).Find(&applications)

	if err = composeSQL.Error; err != nil {
		applicationComposer.Ok = false
		return applicationComposer,err
	}
	affected := composeSQL.RowsAffected
	applicationComposer.Data = applications
	applicationComposer.Ok = true
	applicationComposer.Size = intSize
	applicationComposer.Page = intPage
	applicationComposer.Total = affected
	return applicationComposer,err
}

type ApplicationInsertComposer struct {
	Pagination
	Data []model.ApplicationInsert `json:"data"`
}

func ApplicationInsertPaginator(sql *gorm.DB,page,size string, applications []model.ApplicationInsert) (data ApplicationInsertComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Offset((intPage-1)*intSize).Find(&applications)
	if err = composeSQL.Error; err != nil {
		data.Ok = false
		return data,err
	}
	affected := composeSQL.RowsAffected
	data.Data = applications
	data.Ok = true
	data.Size = intSize
	data.Page = intPage
	data.Total = affected
	return data,err
}

type MyInsertComposer struct {
	Pagination
	Data []model.MyInsert `json:"data"`
}

func MyInsertPaginator(sql *gorm.DB,page,size string, applications []model.MyInsert) (data MyInsertComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Offset((intPage-1)*intSize).Find(&applications)
	if err = composeSQL.Error; err != nil {
		data.Ok = false
		return data,err
	}
	affected := composeSQL.RowsAffected
	data.Data = applications
	data.Ok = true
	data.Size = intSize
	data.Page = intPage
	data.Total = affected
	return data,err
}

type MyDeviceComposer struct {
	Pagination
	Data []model.MyDevice `json:"data"`
}

func MyDevicePaginator(sql *gorm.DB,page,size string, applications []model.MyDevice) (data MyDeviceComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Offset((intPage-1)*intSize).Find(&applications)
	if err = composeSQL.Error; err != nil {
		data.Ok = false
		return data,err
	}
	affected := composeSQL.RowsAffected
	data.Data = applications
	data.Ok = true
	data.Size = intSize
	data.Page = intPage
	data.Total = affected
	return data,err
}

type ModelComposer struct {
	Pagination
	Data []model.Model `json:"data"`
}

func ModelPaginator(sql *gorm.DB,page,size string, applications []model.Model) (data ModelComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Offset((intPage-1)*intSize).Find(&applications)
	if err = composeSQL.Error; err != nil {
		data.Ok = false
		return data,err
	}
	affected := composeSQL.RowsAffected
	data.Data = applications
	data.Ok = true
	data.Size = intSize
	data.Page = intPage
	data.Total = affected
	return data,err
}

type ModuleComposer struct {
	Pagination
	Data []model.Module `json:"data"`
}

func ModulePaginator(sql *gorm.DB,page,size string, applications []model.Module) (data ModuleComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Offset((intPage-1)*intSize).Find(&applications)
	if err = composeSQL.Error; err != nil {
		data.Ok = false
		return data,err
	}
	affected := composeSQL.RowsAffected
	data.Data = applications
	data.Ok = true
	data.Size = intSize
	data.Page = intPage
	data.Total = affected
	return data,err
}

type TemplateComposer struct {
	Pagination
	Data []model.Template `json:"data"`
}

func TemplatePaginator(sql *gorm.DB,page,size string, applications []model.Template) (data TemplateComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Offset((intPage-1)*intSize).Find(&applications)
	if err = composeSQL.Error; err != nil {
		data.Ok = false
		return data,err
	}
	affected := composeSQL.RowsAffected
	data.Data = applications
	data.Ok = true
	data.Size = intSize
	data.Page = intPage
	data.Total = affected
	return data,err
}

type EnterPriseComposer struct {
	Pagination
	Data []model.EnterpriseLib `json:"data"`
}

func EnterPrisePaginator(sql *gorm.DB,page,size string, applications []model.EnterpriseLib) (data EnterPriseComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Offset((intPage-1)*intSize).Find(&applications)
	if err = composeSQL.Error; err != nil {
		data.Ok = false
		return data,err
	}
	affected := composeSQL.RowsAffected
	data.Data = applications
	data.Ok = true
	data.Size = intSize
	data.Page = intPage
	data.Total = affected
	return data,err
}

type ProjectComposer struct {
	Pagination
	Data []model.ProjectLib `json:"data"`
}

func ProjectPaginator(sql *gorm.DB,page,size string, applications []model.ProjectLib) (data ProjectComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Offset((intPage-1)*intSize).Find(&applications)
	if err = composeSQL.Error; err != nil {
		data.Ok = false
		return data,err
	}
	affected := composeSQL.RowsAffected
	data.Data = applications
	data.Ok = true
	data.Size = intSize
	data.Page = intPage
	data.Total = affected
	return data,err
}

type ParkComposer struct {
	Pagination
	Data []model.ParkLib `json:"data"`
}

func ParkPaginator(sql *gorm.DB,page,size string, applications []model.ParkLib) (data ParkComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Offset((intPage-1)*intSize).Find(&applications)
	if err = composeSQL.Error; err != nil {
		data.Ok = false
		return data,err
	}
	affected := composeSQL.RowsAffected
	data.Data = applications
	data.Ok = true
	data.Size = intSize
	data.Page = intPage
	data.Total = affected
	return data,err
}

type MonthlyReportComposer struct {
	Pagination
	Data []model.MonthlyReport `json:"data"`
}

func MonthlyReportPaginator(sql *gorm.DB,page,size string, applications []model.MonthlyReport) (data MonthlyReportComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Offset((intPage-1)*intSize).Find(&applications)
	if err = composeSQL.Error; err != nil {
		data.Ok = false
		return data,err
	}
	affected := composeSQL.RowsAffected
	data.Data = applications
	data.Ok = true
	data.Size = intSize
	data.Page = intPage
	data.Total = affected
	return data,err
}

type JointComposer struct {
	Pagination
	Data []model.Joint `json:"data"`
}

func JointPaginator(sql *gorm.DB,page,size string, applications []model.Joint) (data JointComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Offset((intPage-1)*intSize).Find(&applications)
	if err = composeSQL.Error; err != nil {
		data.Ok = false
		return data,err
	}
	affected := composeSQL.RowsAffected
	data.Data = applications
	data.Ok = true
	data.Size = intSize
	data.Page = intPage
	data.Total = affected
	return data,err
}

type LibComposer struct {
	Pagination
	Data []model.Lib `json:"data"`
}

func LibPaginator(sql *gorm.DB,page,size string, libs []model.Lib) (data LibComposer,err error) {
	intPage := utils.StrToInt(page)
	intSize := utils.StrToInt(size)
	if intPage<1 {
		intPage = 1
	}
	if intSize<1 {
		intSize = 1
	}else if intSize >100 {
		intSize = 100
	}
	composeSQL := sql.Offset((intPage-1)*intSize).Find(&libs)
	if err = composeSQL.Error; err != nil {
		data.Ok = false
		return data,err
	}
	affected := composeSQL.RowsAffected
	data.Data = libs
	data.Ok = true
	data.Size = intSize
	data.Page = intPage
	data.Total = affected
	return data,err
}
