package model

import "github.com/jinzhu/gorm"

type Model struct {
	gorm.Model
	Type string `json:"type"`
	Name string `json:"name"`
	Owner string `json:"owner"`
	CreateTime string `json:"create_time"`
}
type Module struct {
	gorm.Model
	Version string `json:"version"`
	Name string `json:"name"`
	Status string `gorm:"default:'1'" json:"status"`
	Owner string `json:"owner"`
	CreateTime string `json:"create_time"`
	Models []Model `gorm:"many2many:module_model" json:"models"`
}
type Template struct {
	gorm.Model
	Name string `json:"name"`
	Status string `gorm:"default:'1'" json:"status"`
	Owner string `json:"owner"`
	CreateTime string `json:"create_time"`
	Modules []Module `gorm:"many2many:template_module" json:"modules"`
}

type ModuleIdList struct {
	Module
	IdList
}

type TemplateIdList struct {
	Template
	IdList
}
