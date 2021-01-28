package independent

import (
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/model"
	"IndustrialInternetApi/model/paginate"
	"github.com/gin-gonic/gin"
)

func GetAllMonthlyReports(c *gin.Context) (monthlyReportComposer paginate.MonthlyReportComposer,err error) {
	page := c.Param("page")
	pagesize := c.Param("pageSize")
	var userList []model.MonthlyReport
	SQL:=config.DB.Model(&model.MonthlyReport{})
	monthlyReportComposer,err = paginate.MonthlyReportPaginator(SQL,page,pagesize,userList)
	return
}
func GetMonthlyReportById(ID uint) (monthlyReport model.MonthlyReport,err error) {
	if err =config.DB.Where("id = ?",ID).First(&monthlyReport).Error;err != nil{
		return model.MonthlyReport{}, err
	}
	return
}
func CreateMonthlyReport(c *gin.Context)(monthlyReport model.MonthlyReport,err error){
	_ = c.BindJSON(&monthlyReport)
	if err = config.DB.Model(&monthlyReport).Create(&monthlyReport).Error;err!=nil{
		return model.MonthlyReport{}, err
	}
	return
}
func UpdateMonthlyReport(c *gin.Context,ID uint)(monthlyReport model.MonthlyReport,err error){
	_ = c.BindJSON(&monthlyReport)
	if err = config.DB.Model(&monthlyReport).Where("id = ?",ID).Update(&monthlyReport).Error;err !=nil{
		return model.MonthlyReport{}, err
	}
	return
}
func DeleteMonthlyReport(ID uint)(affe int64){
	affe =config.DB.Where("id = ?",ID).First(&model.MonthlyReport{}).Delete(&model.MonthlyReport{}).RowsAffected
	return
}

func GetAllStatisticReports() (statisticReports []model.StatisticReport,err error) {
	if err = config.DB.Find(&statisticReports).Error;err !=nil{
		return nil, err
	}
	return
}
func GetStatisticReportById(ID uint) (statisticReport model.StatisticReport,err error) {
	if err =config.DB.Where("id = ?",ID).First(&statisticReport).Error;err != nil{
		return model.StatisticReport{}, err
	}
	return
}
func CreateStatisticReport(c *gin.Context)(statisticReport model.StatisticReport,err error){
	_ = c.BindJSON(&statisticReport)
	if err = config.DB.Model(&statisticReport).Create(&statisticReport).Error;err!=nil{
		return model.StatisticReport{}, err
	}
	return
}
func UpdateStatisticReport(c *gin.Context,ID uint)(statisticReport model.StatisticReport,err error){
	_ = c.BindJSON(&statisticReport)
	if err = config.DB.Model(&statisticReport).Where("id = ?",ID).Update(&statisticReport).Error;err !=nil{
		return model.StatisticReport{}, err
	}
	return
}
func DeleteStatisticReport(ID uint)(affe int64){
	affe =config.DB.Where("id = ?",ID).First(&model.StatisticReport{}).Delete(&model.StatisticReport{}).RowsAffected
	return
}