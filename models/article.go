package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Article 文章结构体
type Article struct {
	gorm.Model
	CreatedBy     string `json:"created_by"`
	UpdatedBy     string `json:"updated_by"`
	Deleted       uint   `json:"deteled"`
	State         uint   `json:"state"`
	TagID         uint   `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageURL string `json:"cover_image_url"`
}

//BeforeCreate 在创建Article之前，先把创建时间赋值
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now().Unix())
	return nil
}

//BeforeUpdate 在更新Article之前，先把更新时间赋值
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}
