package application

import (
	"github.com/aguncn/nezha/models"
)

//IApplicationService Application接口定义
type IApplicationService interface {
	//GetApplication 根据id获取Application
	GetApplication(id uint) *models.Application

	//AddApplication 新增Application
	AddApplication(application *models.Application) bool
	//UpdateApplication 更新Application
	UpdateApplication(application *models.Application) bool
	//GetApplications 获取文章信息
	GetApplications(page, pagesize uint, maps interface{}) interface{}
	//DeleteApplication 删除Application
	DeleteApplication(id uint) bool
}
