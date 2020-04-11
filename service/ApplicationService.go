package service

import (
	"github.com/aguncn/nezha/models"
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
	return a.Repository.UpdateApplication(application)
}

//GetArticles 获取Application信息
func (a *ApplicationService) GetApplications(PageNum, PageSize uint, total *uint64, where interface{}) *[]models.Application {
	return a.Repository.GetApplications(PageNum, PageSize, total, where)
}

//DeleteApplication 删除用户
func (a *ApplicationService) DeleteApplication(id uint) bool {
	return a.Repository.DeleteApplication(id)
}
