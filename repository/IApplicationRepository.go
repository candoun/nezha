package repository

import "github.com/aguncn/nezha/models"

//IApplicationRepository Application接口定义
type IApplicationRepository interface {
	//GetApplication 根据id获取Application
	GetApplication(where interface{}) *models.Application
	//AddApplication 新增Application
	AddApplication(application *models.Application) bool
	//GetApplications 获取文章
	GetApplications(PageNum, PageSize uint, total *uint64, where interface{}) *[]models.Application
}
