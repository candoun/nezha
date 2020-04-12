package service

import (
	"github.com/aguncn/nezha/models"
	//pageModel "github.com/aguncn/nezha/page"
	"github.com/aguncn/nezha/repository"
	"github.com/jinzhu/gorm"
)

// ArticleService 注入IArticleRepo
type ApplicationService struct {
	Repository repository.IApplicationRepository `inject:""`
}

//GetApplication 根据id获取Application
func (a *ApplicationService) GetApplication(id uint) *models.Application {
	where := models.Application{Model: gorm.Model{ID: id}}
	return a.Repository.GetApplication(&where)
}

//AddApplication 新增Application
func (a *ApplicationService) AddApplication(application *models.Application) bool {
	if a.Repository.ExistApplicationByName(application) {
		return false
	} else {
		return a.Repository.AddApplication(application)
	}
}

//UpdateUser 更新用户
func (a *ApplicationService) UpdateApplication(application *models.Application) bool {
	if a.Repository.ExistApplicationByName(application) {
		return false
	} else {
		return a.Repository.UpdateApplication(application)
	}
}

//GetArticles 获取Application信息
func (a *ApplicationService) GetApplications(page, pagesize uint, maps interface{}) interface{} {
	res := make(map[string]interface{}, 2)
	var total uint64
	applications := a.Repository.GetApplications(page, pagesize, &total, maps)
	/*
		可用于定制输出
		var pageApplications []pageModel.Application
		var pageApplication pageModel.Application
		for _, application := range *applications {
			pageApplication.ID = application.ID
			pageApplication.Name = application.Name
			pageApplication.Git = application.Git
			pageApplication.Jenkins = application.Jenkins
			pageApplication.State = application.State
			pageApplication.CreatedAt = application.CreatedAt.Format("2006-01-02 15:04:05")
			pageApplications = append(pageApplications, pageApplication)
		}
	*/
	res["list"] = &applications
	res["total"] = total
	return &res
}

//DeleteApplication 删除用户
func (a *ApplicationService) DeleteApplication(id uint) bool {
	return a.Repository.DeleteApplication(id)
}
