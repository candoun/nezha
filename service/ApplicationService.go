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
	return a.Repository.AddApplication(application)
}

//GetArticles 获取文章信息
func (a *ApplicationService) GetApplications(PageNum, PageSize uint, total *uint64, where interface{}) *[]models.Application {
	return a.Repository.GetApplications(PageNum, PageSize, total, where)
}
