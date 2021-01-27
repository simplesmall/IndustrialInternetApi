package model

import "github.com/jinzhu/gorm"

type Mirror struct {
	gorm.Model
	Name string `json:"name"`
	Size int `json:"size"`
	Desc string `json:"desc"`
	UpdateTime string `json:"update_time"`
}
type Application struct {
	gorm.Model
	Name string `json:"name"`
	NickName string `json:"nick_name"`
	GitUrl string `json:"git_url"`
	Status string `gorm:"default:'1'" json:"status"`
	CreateTime string `json:"create_time"`
	Mirrors []Mirror `gorm:"many2many:application_mirror" json:"mirrors"`
}

type ApplicationIdList struct {
	Application
	IdList
}