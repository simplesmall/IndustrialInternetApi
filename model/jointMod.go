package model

import "github.com/jinzhu/gorm"

type Joint struct {
	gorm.Model
	Name string `json:"name"`
	Desc string `json:"desc"`
	Type string `json:"type"`
	Url string `json:"url"`
}
