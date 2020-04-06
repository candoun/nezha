package repository

import (
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
)

//ApplicationRepository 注入IDb
type ApplicationRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

//GetApplication 根据id获取Application
func (a *ApplicationRepository) GetApplication(where interface{}) *models.Application {
	var application models.Application
	if err := a.Base.First(where, &application); err != nil {
		a.Log.Errorf("未找到相关文章", err)
	}
	return &application
}

//AddApplication 新增Application
func (a *ApplicationRepository) AddApplication(application *models.Application) bool {
	if err := a.Base.Save(application); err != nil {
		a.Log.Errorf("添加文章失败", err)
	}
	return true
}

//GetArticles 获取文章
func (a *ApplicationRepository) GetApplications(PageNum uint, PageSize uint, total *uint64, where interface{}) *[]models.Application {
	var applications []models.Application
	err := a.Base.GetPages(&models.Application{}, &applications, PageNum, PageSize, total, where, "ID desc")
	if err != nil {
		a.Log.Errorf("获取文章信息失败", err)
	}
	return &applications
}
