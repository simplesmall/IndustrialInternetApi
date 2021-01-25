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