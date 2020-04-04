package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Application 应用结构体
type Application struct {
	gorm.Model
	Name        string `json:"name"`
	CnName      string `json:"cn_name"`
	Description string `json:"description"`
	UserID      int    `json:"user_id"`
	ProjectID   int    `json:"project_id"`
	CreatedBy   string `json:"created_by"`
	ModifiedBy  string `json:"modified_by"`
	State       int    `json:"state"`
}

//BeforeCreate CreatedOn赋值
func (application *Application) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//BeforeUpdate ModifiedOn赋值
func (application *Application) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
