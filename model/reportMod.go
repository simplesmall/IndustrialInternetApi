package model

import "github.com/jinzhu/gorm"

type MonthlyReport struct {
	gorm.Model
	TotalValue string `json:"total_value"`			// 总产值
	IndustAddValue string `json:"indust_add_value"` // 工业附加值
	BusinessIncome string `json:"business_income"`	// 营业收入
	PaidTax string `json:"paid_tax"`				// 实缴税金
	CapitalInvest string `json:"capital_invest"`	//基建投资
}
type StatisticReport struct {
	gorm.Model
	Type string `json:"type"`
	Name string `json:"name"`
	Url string `json:"url"`
}