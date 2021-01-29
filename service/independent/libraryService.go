package independent

import (
	"IndustrialInternetApi/config"
	"IndustrialInternetApi/model"
	"IndustrialInternetApi/model/paginate"
	"IndustrialInternetApi/service/user"
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetAllEnterPrises(c *gin.Context) (enterPriseComposer paginate.EnterPriseComposer,err error) {
	page := c.Param("page")
	pagesize := c.Param("pageSize")
	var applications []model.EnterpriseLib
	SQL:=config.DB.Model(&model.EnterpriseLib{})
	enterPriseComposer,err = paginate.EnterPrisePaginator(SQL,page,pagesize, applications)
	return
}
func GetEnterPriseById(ID uint) (enterpriseLib model.EnterpriseLib,err error) {
	if err =config.DB.Where("id = ?",ID).First(&enterpriseLib).Error;err != nil{
		return model.EnterpriseLib{}, err
	}
	return
}

func GetAllProjects(c *gin.Context) (projectComposer paginate.ProjectComposer,err error) {
	page := c.Param("page")
	pagesize := c.Param("pageSize")
	var applications []model.ProjectLib
	SQL:=config.DB.Model(&model.ProjectLib{})
	projectComposer,err = paginate.ProjectPaginator(SQL,page,pagesize, applications)
	return
}
func GetProjectById(ID uint) (projectLib model.ProjectLib,err error) {
	if err =config.DB.Where("id = ?",ID).First(&projectLib).Error;err != nil{
		return model.ProjectLib{}, err
	}
	return
}

func CreateEnterPrise(c *gin.Context)(enterprise model.EnterpriseLib,err error){
	_ = c.BindJSON(&enterprise)
	if err = config.DB.Model(&enterprise).Create(&enterprise).Error;err!=nil{
		return model.EnterpriseLib{}, err
	}
	return
}
func UpdateEnterPrise(c *gin.Context,ID uint)(enterprise model.EnterpriseLib,err error){
	_ = c.BindJSON(&enterprise)
	if err = config.DB.Model(&enterprise).Where("id = ?",ID).Update(&enterprise).Error;err !=nil{
		return model.EnterpriseLib{}, err
	}
	return
}
func DeleteEnterPrise(ID uint)(affe int64){
	affe =config.DB.Model(&model.EnterpriseLib{}).Where("id = ?",ID).First(&model.EnterpriseLib{}).Delete(&model.EnterpriseLib{}).RowsAffected
	return
}

func CreateProject(c *gin.Context)(projectLib model.ProjectLib,err error){
	_ = c.BindJSON(&projectLib)
	if err = config.DB.Model(&projectLib).Create(&projectLib).Error;err!=nil{
		return model.ProjectLib{}, err
	}
	return
}
func UpdateProject(c *gin.Context,ID uint)(projectLib model.ProjectLib,err error){
	_ = c.BindJSON(&projectLib)
	if err = config.DB.Model(&projectLib).Where("id = ?",ID).Update(&projectLib).Error;err !=nil{
		return model.ProjectLib{}, err
	}
	return
}
func DeleteProject(ID uint)(affe int64){
	affe =config.DB.Where("id = ?",ID).First(&model.ProjectLib{}).Delete(&model.ProjectLib{}).RowsAffected
	return
}

func GetAllParks(c *gin.Context) (parkComposer paginate.ParkComposer,err error) {
	page := c.Param("page")
	pagesize := c.Param("pageSize")
	var parkLibs []model.ParkLib
	SQL:=config.DB.Model(&model.ParkLib{})
	parkComposer,err = paginate.ParkPaginator(SQL,page,pagesize, parkLibs)
	return
}
func GetParkById(ID uint) (parkLib model.ParkLib,err error) {
	if err =config.DB.Where("id = ?",ID).First(&parkLib).Error;err != nil{
		return model.ParkLib{}, err
	}
	return
}
func CreatePark(c *gin.Context)(parkLib model.ParkLib,err error){
	_ = c.BindJSON(&parkLib)
	if err = config.DB.Model(&parkLib).Create(&parkLib).Error;err!=nil{
		return model.ParkLib{}, err
	}
	return
}
func UpdatePark(c *gin.Context,ID uint)(parkLib model.ParkLib,err error){
	_ = c.BindJSON(&parkLib)
	if err = config.DB.Model(&parkLib).Where("id = ?",ID).Update(&parkLib).Error;err !=nil{
		return model.ParkLib{}, err
	}
	return
}
func DeletePark(ID uint)(affe int64){
	affe =config.DB.Where("id = ?",ID).First(&model.ParkLib{}).Delete(&model.ParkLib{}).RowsAffected
	return
}

func GetAllLibs(c *gin.Context) (libComposer paginate.LibComposer,err error) {
	libType := c.Param("type")
	page := c.Param("page")
	pagesize := c.Param("pageSize")
	var userList []model.Lib
	SQL:=config.DB.Model(&model.Lib{}).Where("type = ?",libType)
	libComposer,err = paginate.LibPaginator(SQL,page,pagesize,userList)
	return libComposer,err
}
func GetLibById(ID uint) (lib model.Lib,err error) {
	if err =config.DB.Where("id = ?",ID).First(&lib).Error;err != nil{
		return model.Lib{}, err
	}
	return
}
func CreateLib(c *gin.Context)(lib model.Lib,err error){
	_ = c.BindJSON(&lib)
	if err = config.DB.Model(&lib).Create(&lib).Error;err!=nil{
		return model.Lib{}, err
	}
	return
}
func UpdateLib(c *gin.Context,ID uint)(lib model.Lib,err error){
	_ = c.BindJSON(&lib)
	if err = config.DB.Model(&lib).Where("id = ?",ID).Update(&lib).Error;err !=nil{
		return model.Lib{}, err
	}
	return
}
func DeleteLib(ID uint)(affe int64){
	affe =config.DB.Where("id = ?",ID).First(&model.Lib{}).Delete(&model.Lib{}).RowsAffected
	return
}
func UploadLib(c *gin.Context)(rows [][]string,msg string,err error){
	libType := c.Param("type")
	//获取表单数据 参数为name值
	f, err := c.FormFile("file")
	//错误处理
	if err != nil {
		msg = "表单上传字段名错误,应为file"
		return
	}
	// 隔离缺乏文件字段上传
	if f == nil {
		msg = "缺少上传文件字段"
		return
	}
	// 过滤不接受的上传文件类型[用文件后缀过滤]
	fileSurfix := strings.SplitAfterN(f.Filename, ".", -1)
	if !(fileSurfix[1] == "xlsx" || fileSurfix[1] == "xls") {
		msg = "上传文件类型错误,请选择文件后缀为xlsx或xls的文件上传"
		err = errors.New("文件类型错误")
		return
	}
	// 限制上传文件大小不超过10M
	if f.Size/1024/1024 > 5 {
		msg ="上传文件不能大于5M"
		err = errors.New("文件过大")
		return
	}
	//将文件保存至本项目根目录中
	uploadFilePath := "./excel/upload/" + strconv.FormatInt(time.Now().Unix(), 10) + f.Filename
	saveErr := c.SaveUploadedFile(f, uploadFilePath)
	if saveErr != nil {
		c.JSON(500, gin.H{
			"err": err,
		})
		return
	}
	//上传成功读取文件内容
	excel, _ := excelize.OpenFile(uploadFilePath)
	// 判断是否是规定格式的excel文件 包括表头为固定N列以及每列命名固定
	cols, err := excel.GetCols("Sheet1")
	// 判定通过,将数据写入数据库  cols == 4 && A1 == 区县 && B1 == 类别 && C1 == 名称
	if !(len(cols) == 4){
		msg = "上传的excel数据表字段不符"
		err = errors.New("")
		return
	}
	rows, _ = excel.GetRows("Sheet1")
	//将每一行数据封装成一行记录存入数据库
	loginUser,_:=user.GetLoginUser(c)
	for _,item := range rows[1:]{
		// 将每一个item封装成一条可插入数据库的ORM记录
		libCol := model.Lib{Model: gorm.Model{}, Name: item[0], City: item[1], Type: libType, Category: item[3],Owner:loginUser.User.Account }
		// 插入数据库中
		config.DB.Model(&model.Lib{}).Create(&libCol)
	}
	//保存成功返回正确的Json数据
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"rows": rows,
		"cols": len(cols),
	})
	return
}
