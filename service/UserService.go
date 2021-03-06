package service

import (
	"github.com/aguncn/nezha/common/logger"
	"github.com/aguncn/nezha/models"
	pageModel "github.com/aguncn/nezha/page"
	"github.com/aguncn/nezha/page/emun"
	"github.com/aguncn/nezha/repository"
)

// UserService 注入IUserRepository
type UserService struct {
	Repository     repository.IUserRepository `inject:""`
	RoleRepository repository.IRoleRepository `inject:""`
	Log            logger.ILogger             `inject:""`
}

//CheckUser 身份验证
func (a *UserService) CheckUser(username string, password string) bool {
	where := models.User{Username: username, Password: password}
	return a.Repository.CheckUser(where)
}

//GetUserAvatar 获取用户头像
func (a *UserService) GetUserAvatar(username string) *string {
	where := models.User{Username: username}
	sel := "avatar"
	return a.Repository.GetUserAvatar(&sel, &where)
}

//GetUserAvatar 获取用户头像
func (a *UserService) GetUserID(username string) uint {
	userWhere := models.User{Username: username}
	userSel := "id"
	userID := a.Repository.GetUserID(&userSel, &userWhere)
	return userID
}

//GetRoles 获取用户角色
func (a *UserService) GetRoles(username string) *[]string {
	userWhere := models.User{Username: username}
	userSel := "id"
	userID := a.Repository.GetUserID(&userSel, &userWhere)
	roleWhere := models.Role{UserID: uint(userID)}
	roleSel := "value"
	return a.RoleRepository.GetRoles(&roleSel, &roleWhere)
}

//GetUsers 获取用户信息
func (a *UserService) GetUsers(page, pagesize uint, maps interface{}) interface{} {
	res := make(map[string]interface{}, 2)
	var total uint64
	users := a.Repository.GetUsers(page, pagesize, &total, maps)
	var pageUsers []pageModel.Users
	var pageUser pageModel.Users
	for _, user := range *users {
		pageUser.ID = user.ID
		pageUser.Name = user.Username
		pageUser.Password = user.Password
		pageUser.Avatar = user.Avatar
		pageUser.UserType = emun.GetUserType(user.UserType)
		pageUser.State = emun.GetStatus(user.State)
		pageUser.Deleted = emun.GetDeleted(user.Deleted)
		pageUser.CreatedAt = user.CreatedAt.Format("2006-01-02 15:04:05")
		pageUsers = append(pageUsers, pageUser)
	}
	res["list"] = &pageUsers
	res["total"] = total
	return &res
}

//RegisterUser 新建用户，并不同时新建用户角色
func (a *UserService) RegisterUser(user *models.User) bool {

	isOK := a.Repository.AddUser(user)
	if !isOK {
		return false
	}
	return true
}

//AddUser 新建用户，同时新建用户角色
func (a *UserService) AddUser(user *models.User) bool {
	//此处不能使用事务同时创建用户和角色，因为Role表中需要UserID，而UserID需要插入用户数据后才生成，所以不能用事务，否则会报错
	//用业务逻辑实现事务效果
	isOK := a.Repository.AddUser(user)
	if !isOK {
		return false
	}
	//当成功插入User数据后，user为指针地址，可以获取到ID的值。省去了查数据库拿ID的值步骤
	var role models.Role
	role.UserID = user.ID
	role.UserName = user.Username
	role.Value = "dev"
	if user.UserType == 1 {
		role.Value = "admin"
	}
	isOK = a.RoleRepository.AddRole(&role)
	if isOK {
		return true
	}
	//插入role失败后，删除新插入的用户信息，达到事务处理效果
	return a.Repository.DeleteUser(user.ID)
}

//ExistUserByName 判断用户名是否已存在
func (a *UserService) ExistUserByName(username string) bool {
	where := models.User{Username: username}
	return a.Repository.ExistUserByName(where)
}

//UpdateUser 更新用户
func (a *UserService) UpdateUser(modUser *models.User) bool {
	user := a.Repository.GetUserByID(modUser.ID)
	//不允许更新用户名
	// user.Username = modUser.Username
	user.Password = modUser.Password
	user.UpdatedBy = modUser.UpdatedBy
	user.UserType = modUser.UserType
	roleWhere := models.Role{UserID: user.ID}
	role := a.RoleRepository.GetRole(&roleWhere)
	// role.UserName = user.Username
	role.Value = "test"
	if user.UserType == 1 {
		role.Value = "admin"
	}
	return a.Repository.UpdateUser(user, role)
}

//DeleteUser 删除用户
func (a *UserService) DeleteUser(id uint) bool {
	user := a.Repository.GetUserByID(id)
	if user.Username == "admin" {
		a.Log.Errorf("删除用户失败:不能删除admin账号")
		return false
	}
	return a.Repository.DeleteUser(id)
}
