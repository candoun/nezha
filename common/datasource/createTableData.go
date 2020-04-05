package datasource

import (
	"fmt"

	"github.com/aguncn/nezha/models"
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
		ID:       1,
		UserID:   1,
		UserName: "admin",
		Value:    "admin",
	}
	mockData["testRole"] = &models.Role{
		ID:       2,
		UserID:   1,
		UserName: "admin",
		Value:    "test",
	}
	mockData["adminTestRole"] = &models.Role{
		ID:       3,
		UserID:   2,
		UserName: "test",
		Value:    "test",
	}

	mockData["tag1"] = &models.Tag{
		ID:   1,
		Name: "1",
	}
	mockData["tag2"] = &models.Tag{
		ID:   2,
		Name: "2",
	}
	mockData["tag3"] = &models.Tag{
		ID:   3,
		Name: "3",
	}
	mockData["article1"] = &models.Article{
		ID:      1,
		TagID:   1,
		Title:   "Title1",
		Content: "Content1",
	}
	mockData["article2"] = &models.Article{
		ID:      2,
		TagID:   2,
		Title:   "Title2",
		Content: "Content2",
	}
	mockData["article3"] = &models.Article{
		ID:      3,
		TagID:   3,
		Title:   "Title3",
		Content: "Content3",
	}
}
