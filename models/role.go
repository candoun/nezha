package models

import "github.com/jinzhu/gorm"

//Role 身份信息结构体
type Role struct {
	gorm.Model
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	Deleted   uint   `json:"deteled"`
	State     uint   `json:"state"`
	UserID    uint   `json:"user_id"`
	UserName  string `json:"user_name"`
	Value     string `json:"value"`
}
