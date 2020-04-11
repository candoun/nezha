package repository

import (
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	"github.com/jinzhu/gorm"
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
	if err := a.Base.Create(application); err != nil {
		a.Log.Errorf("添加文章失败", err)
		return false
	}
	return true
}

//UpdateApplication 更新Application
func (a *ApplicationRepository) UpdateApplication(application *models.Application) bool {
	if err := a.Base.Save(application); err != nil {
		a.Log.Errorf("更新应用失败", err)
		return false
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

//ExistApplicationByName 判断用户名是否已存在
func (a *ApplicationRepository) ExistApplicationByName(where interface{}) bool {
	var application models.Application
	var whereApplication = models.Application{Name: where.(*models.Application).Name}
	err := a.Base.First(whereApplication, &application)
	//记录不存在错误(RecordNotFound)，返回false
	if gorm.IsRecordNotFoundError(err) {
		return false
	}
	//其他类型的错误，写下日志，返回false
	if err != nil {
		a.Log.Error(err)
		return false
	}
	return true
}

//DeleteApplication 删除应用
func (a *ApplicationRepository) DeleteApplication(id uint) bool {
	var application models.Application
	if err := a.Base.DeleteByID(application, id); err != nil {
		a.Log.Errorf("删除应用失败", err)
		return false
	}
	return true
}
