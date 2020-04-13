package controller

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/aguncn/nezha/common/codes"
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	appSvc "github.com/aguncn/nezha/service/application"
)

// Application 注入IApplicationService
type Application struct {
	Log     logger.ILogger             `inject:""`
	Service appSvc.IApplicationService `inject:""`
}

//GetApplication 获取单个Application
func (a *Application) GetApplication(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	var data *models.Application
	code := codes.InvalidParams
	if !valid.HasErrors() {
		data = a.Service.GetApplication(uint(id))
		code = codes.SUCCESS
	} else {
		for _, err := range valid.Errors {
			a.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	RespData(c, http.StatusOK, code, data)
}

//AddApplication 新增Application
func (a *Application) AddApplication(c *gin.Context) {

	application := models.Application{}

	code := codes.InvalidParams
	err := c.Bind(&application)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(application.Name, "name").Message("名称不能为空")
		valid.Required(application.Description, "description").Message("简述不能为空")
		valid.Required(application.Git, "git").Message("Git不能为空")
		valid.Required(application.Jenkins, "jenkins").Message("Jenkins不能为空")
		valid.Required(application.CreatedBy, "created_by").Message("创建人不能为空")
		if !valid.HasErrors() {
			if a.Service.AddApplication(&application) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				a.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

//UpdateApplication 修改Application
func (a *Application) UpdateApplication(c *gin.Context) {
	applicaion := models.Application{}
	code := codes.InvalidParams
	err := c.Bind(&applicaion)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(applicaion.Git, "git").Message("Git不能为空")
		valid.Required(applicaion.Jenkins, "jenkins").Message("Jenkins不能为空")
		if !valid.HasErrors() {
			if a.Service.UpdateApplication(&applicaion) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			for _, err := range valid.Errors {
				a.Log.Info("err.key: %s, err.message: %s", err.Key, err.Message)
			}
		}
	}
	RespOk(c, http.StatusOK, code)
}

//GetApplications 获取Applications信息
func (a *Application) GetApplications(c *gin.Context) {
	var maps string
	code := codes.SUCCESS
	name := c.Query("name")
	if name != "" {
		maps = "name LIKE '%" + name + "%'"
	}
	page, pagesize := GetPage(c)
	data := a.Service.GetApplications(page, pagesize, maps)
	RespData(c, http.StatusOK, code, data)
}

//DeleteApplication 删除用户
func (a *Application) DeleteApplication(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := codes.SUCCESS
	if !a.Service.DeleteApplication(uint(id)) {
		code = codes.ERROR
		RespFail(c, http.StatusOK, code, "删除出错，请联系管理员!")
	} else {
		RespOk(c, http.StatusOK, code)
	}
}
