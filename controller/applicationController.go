package controller

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/aguncn/nezha/common/codes"
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	"github.com/aguncn/nezha/service"
)

// Application 注入IApplicationService
type Application struct {
	Log     logger.ILogger              `inject:""`
	Service service.IApplicationService `inject:""`
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
		application.UpdatedAt = application.CreatedAt
		valid := validation.Validation{}
		valid.Required(application.Name, "name").Message("名称不能为空")
		valid.Required(application.Description, "description").Message("简述不能为空")
		valid.Required(application.Git, "git").Message("Git不能为空")
		valid.Required(application.Jenkins, "jenkins").Message("Jenkins不能为空")
		valid.Required(application.CreatedBy, "created_by").Message("创建人不能为空")
		valid.Range(application.State, 0, 1, "state").Message("状态只允许0或1")
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

//GetApplications 获取Applications信息
func (a *Application) GetApplications(c *gin.Context) {
	res := make(map[string]interface{}, 2)
	var total uint64
	code := codes.SUCCESS
	page, pagesize := GetPage(c)
	applications := a.Service.GetApplications(page, pagesize, &total, "")
	res["list"] = &applications
	res["total"] = total
	RespData(c, http.StatusOK, code, &res)
}
