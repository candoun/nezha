package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Tag 标签结构体
type Tag struct {
	gorm.Model
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	Deleted   uint   `json:"deteled"`
	State     int    `json:"state"`
	Name      string `json:"name"`
}

//BeforeCreate CreatedOn赋值
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//BeforeUpdate ModifiedOn赋值
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
