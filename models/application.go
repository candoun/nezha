package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Application 应用结构体
type Application struct {
	gorm.Model
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	Deleted     uint   `json:"deteled"`
	State       uint   `json:"state"`
	Name        string `json:"name"`
	CnName      string `json:"cn_name"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	ProjectID   uint   `json:"project_id"`
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
