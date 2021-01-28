package model

import "github.com/jinzhu/gorm"

type EnterpriseLib struct {
	gorm.Model
	Name string `json:"name"`
	Type string `json:"type"`
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