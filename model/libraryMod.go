package model

import "github.com/jinzhu/gorm"

type Lib struct {
	gorm.Model
	Name string `json:"name"`     // 名称
	Type string `json:"type"`	  // 类别
	City string `json:"city"`	  // 区县
	Category string `json:"category"` // 分类 企业库 项目库 园区库
	Owner string `json:"owner"`		  // 上传者
}

type EnterpriseLib struct {
	gorm.Model
	Name string `json:"name"`     // 名称
	Type string `json:"type"`	  // 类别
}
type ProjectLib struct {
	gorm.Model
	Name string `json:"name"`
	Type string `json:"type"`
}
type ParkLib struct {
	gorm.Model
	Name string `json:"name"`
	Type string `json:"type"`
}