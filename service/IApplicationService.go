package service

import (
	"github.com/aguncn/nezha/models"
)

//IApplicationService Application接口定义
type IApplicationService interface {
	//GetApplication 根据id获取Application
	GetApplication(id uint) *models.Application

	//AddApplication 新增Application
	AddApplication(application *models.Application) bool
	//GetApplications 获取文章信息
	GetApplications(PageNum, PageSize uint, total *uint64, where interface{}) *[]models.Application
}
