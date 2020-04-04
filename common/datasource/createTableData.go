package datasource

import (
	"fmt"

	"github.com/bingjian-zhu/gin-vue-admin/models"
	"github.com/jinzhu/gorm"
)

var mockData = make(map[string]interface{})

func CreateTableData(db *gorm.DB) error {
	initData()
	for _, record := range mockData {
		if err := db.Create(record).Error; err != nil {
			fmt.Println("用户数据插入失败", err)
			return err
		}
	}
	return nil
}

func initData() {
	mockData["adminUser"] = &models.User{
		ID:       1,
		Username: "admin",
		Password: "111111",
		Avatar:   "https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG",
		UserType: 1,
	}

	mockData["normalNuser"] = &models.User{
		ID:       2,
		Username: "test",
		Password: "111111",
		Avatar:   "https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG",
		UserType: 2,
	}

	mockData["adminRole"] = &models.Role{
		UserName: "admin",
		UserID:   1,
		Value:    "admin",
	}
	mockData["testRole"] = &models.Role{
		UserName: "test",
		UserID:   2,
		Value:    "test",
	}
	mockData["adminTestRole"] = &models.Role{
		UserName: "admin",
		UserID:   1,
		Value:    "test",
	}
}
