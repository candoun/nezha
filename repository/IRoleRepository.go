package repository

import "github.com/aguncn/nezha/models"

//IRoleRepository Role接口定义
type IRoleRepository interface {
	//GetUserRoles 获取用户身份信息
	GetUserRoles(where interface{}) *[]models.Role
	//GetRoles 获取用户角色
	GetRoles(sel *string, where interface{}) *[]string
	//AddRole 添加用户角色
	AddRole(role *models.Role) bool
	//GetRole 获取角色
	GetRole(where interface{}) *models.Role
}
