package model

import (
	"database/sql/driver"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type UserLogin struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}
// 1. 创建 time.Time 类型的副本 XTime；
type XTime struct {
	time.Time
}

const TimeFormat = "2006-01-02 15:04:05"

//MyTime 自定义时间
type MyTime time.Time

func (t *XTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = XTime{now}
	return err
}

// 2. 为 Xtime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t XTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format(TimeFormat))
	return []byte(output), nil
}

// 3. 为 Xtime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 4. 为 Xtime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type BaseModel struct {
	ID        uint `gorm:"primary_key,AUTO_INCREMENT"`
	CreatedAt XTime
	UpdatedAt XTime
	DeletedAt *XTime `sql:"index"`
}

// 用户
type User struct {
	gorm.Model
	Name     string `gorm:"size:100;column(name)" json:"name" form:"name"`
	Account string `gorm:"size:100;column(account)" json:"account" form:"account"`
	Password string `gorm:"size:255;column(password)" json:"password" form:"password"`
	BelongTo string `json:"belong_to" form:"belong_to"`     							// 所属企业
	Status   string `gorm:"size:20;column(status)" json:"status" form:"status"`			//是否有效状态 有效,过期
	IsLogin   string `gorm:"size:20;column(is_login)" json:"is_login" form:"is_login"`	//是否在线状态 在线,离线
	ExpireTime string `json:"expire_time" form:"expire_time"`							//到期时间
	LoginCount	uint `json:"login_count" form:"login_count"`							//登录次数
	UsingTime	string `json:"using_time" form:"using_time"`							//使用时长
	IP string `json:"ip" form:"ip"`														//用户IP
	Token string `gorm:"size:500;column:token" json:"token" form:"token"`				//用户Token
	Type  string `gorm:"size:10;column:type" json:"type" form:"type"`					//用户类型
	Phone  string `gorm:"size:20;column:phone" json:"phone" form:"phone"`				//用户手机
	Email  string `gorm:"size:100;column:email" json:"email" form:"email"`				//用户邮箱
	Remark  string `gorm:"size:200;column:remark" json:"remark" form:"remark"`			//用户备注
	Role     []Role `gorm:"many2many:user_role" json:"role"`
}

//角色
type Role struct {
	gorm.Model
	Name       string       `gorm:"size:100;column(name)" json:"name" form:"name"`
	Desc       string       `gorm:"size:500;column(desc)" json:"desc" form:"desc"`
	Permission []Permission `gorm:"many2many:role_permission"`
}

//权限
type Permission struct {
	gorm.Model
	Name     string `gorm:"size:100;column(name)" json:"name" form:"name"`
	Url      string `gorm:"size:500;column(url)" json:"url" form:"url"`
	Icon     string `gorm:"size:500;column(icon)" json:"icon" form:"icon"`
	Describe string `gorm:"size:200;column(describe)" json:"describe" form:"describe"`
	ParentId   string `gorm:"size:10;column(parent_id)" json:"parent_id" form:"parent_id"`
	Status   string `gorm:"size:10;column(status)" json:"status" form:"status"`			//是否有效状态 有效,过期
	Type     string `gorm:"size:10;column(type)" json:"type" form:"type"`
}

//权限树桩结构
type PermissionTree struct {
	Name     string `gorm:"size:100;column(name)" json:"name" form:"name"`
	Url      string `gorm:"size:500;column(url)" json:"url" form:"url"`
	Icon     string `gorm:"size:500;column(icon)" json:"icon" form:"icon"`
	Describe string `gorm:"size:200;column(describe)" json:"describe" form:"describe"`
	ParentId   string `gorm:"size:10;column(parent_id)" json:"parent_id" form:"parent_id"`
	Status   string `gorm:"size:10;column(status)" json:"status" form:"status"`			//是否有效状态 有效,过期
	Type     string `gorm:"size:10;column(type)" json:"type" form:"type"`
	Children []*PermissionTree
}

type IdList struct {
	Ids []uint `json:"ids"`
}

type RoleIdList struct {
	Role
	IdList
}

type UserIdList struct {
	User
	IdList
}

type DeliveRole struct {
	Names []string `json:"names"`
}