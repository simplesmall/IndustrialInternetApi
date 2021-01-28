package model

import "github.com/jinzhu/gorm"

type ApplicationInsert struct {
	gorm.Model
	Name string `json:"name"`
	NickName string `json:"nick_name"`
	PublicKey string `gorm:"default:'UIGIiuuibdsugfsdiue5785873564uuyvuyv876vv'" json:"public_key"'`
	Url string `json:"url"`
	Status string `gorm:"default:'1'" json:"status"`
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
	MyInserts []MyInsert `gorm:"many2many:mydevice_myinsert" json:"my_inserts"`
}

type MyDeviceIdList struct {
	MyDevice
	IdList
}