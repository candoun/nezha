package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Project 项目结构体
type Project struct {
	ID          int       `gorm:"primary_key" json:"id"`
	CreatedOn   time.Time `json:"created_on"`
	ModifiedOn  time.Time `json:"modified_on"`
	Name        string    `json:"name"`
	CnName      string    `json:"cn_name"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id"`
	CreatedBy   string    `json:"created_by"`
	ModifiedBy  string    `json:"modified_by"`
	State       int       `json:"state"`

	Application []Application
}

//BeforeCreate CreatedOn赋值
func (project *Project) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//BeforeUpdate ModifiedOn赋值
func (project *Project) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
