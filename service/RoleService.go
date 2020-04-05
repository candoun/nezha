package service

import (
	"github.com/aguncn/nezha/models"
	"github.com/aguncn/nezha/repository"
)

// RoleService IRoleRepository
type RoleService struct {
	Repository repository.IRoleRepository `inject:""`
}

//GetUserRoles 分页返回Articles获取用户身份信息
func (c *RoleService) GetUserRoles(userName string) *[]models.Role {
	where := models.Role{UserName: userName}
	return c.Repository.GetUserRoles(&where)
}
