package model

import "github.com/jinzhu/gorm"

type ApplicationInsert struct {
	gorm.Model
	Name string `json:"name"`
	NickName string `json:"nick_name"`
	Url string `json:"url"`
	Status string `json:"status"`
	CreateTime string `json:"create_time"`
}
type MyInsert struct {
	gorm.Model
	InsertName string `json:"insert_name"`
	Protocal string `json:"protocal"`
	TransferPlugin string `json:"transfer_plugin"`
	CreateTime string `json:"create_time"`
}

type MyDevice struct {
	gorm.Model
	Name string `json:"name"`
	NetType string `json:"net_type"`
	Type string `json:"type"`
	CreateTime string `json:"create_time"`
	MyInserts []*MyInsert `json:"my_inserts"`
}
