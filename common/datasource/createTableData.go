package datasource

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var mockData = make(map[string]string)

func CreateTableData(db *gorm.DB) error {
	initData()
	for _, record := range mockData {
		db.Exec(record)
	}
	return nil
}

func initData() {
	pw, _ := bcrypt.GenerateFromPassword([]byte("111111"), bcrypt.DefaultCost)
	mockData["adminUser"] = fmt.Sprintf("INSERT INTO nezha_user(id, username, password, avatar,user_type, state, created_at) VALUES "+
		" (1, 'admin', '%s', 'https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG',1, 1, '2020-04-06 13:01:09');",
		pw)
	mockData["devUser"] = fmt.Sprintf("INSERT INTO nezha_user(id, username, password, avatar,user_type, state, created_at) VALUES "+
		" (2, 'dev', '%s', 'https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG',2, 1, '2020-04-06 13:01:09');",
		pw)

	mockData["tag1"] = "INSERT INTO nezha_tag(id, name, created_at, created_by, updated_at, updated_by, deleted_at, state) VALUES" +
		" (1, '1', '2019-08-18 18:56:01', 'test', NULL, '', 0, 1);"
	mockData["tag2"] = "INSERT INTO nezha_tag(id, name, created_at, created_by, updated_at, updated_by, deleted_at, state) VALUES" +
		" (2, '2', '2019-08-18 18:56:01', 'test', NULL, '', 0, 1);"
	mockData["tag3"] = "INSERT INTO nezha_tag(id, name, created_at, created_by, updated_at, updated_by, deleted_at, state) VALUES" +
		" (3, '3', '2019-08-18 18:56:01', 'test', NULL, '', 0, 1);"

	mockData["role1"] = "INSERT INTO nezha_role(id, user_id, user_name, value) VALUES (1, 1, 'admin', 'admin')"
	mockData["role2"] = "INSERT INTO nezha_role(id, user_id, user_name, value) VALUES (2, 1, 'admin', 'dev')"
	mockData["role3"] = "INSERT INTO nezha_role(id, user_id, user_name, value) VALUES (3, 2, 'dev', 'dev')"

	mockData["project1"] = "INSERT INTO nezha_project(id, user_id, name, cn_name, description) VALUES (1, 1, 'AI-PROJECT', '算法项目', '用于算法')"
	mockData["project2"] = "INSERT INTO nezha_project(id, user_id, name, cn_name, description) VALUES (2, 2, 'BI-PROJECT', '商务智能项目', '用于BI智能')"
	mockData["project3"] = "INSERT INTO nezha_project(id, user_id, name, cn_name, description) VALUES (3, 1, 'CI-PROJECT', '持续集成项目', '用于JKINS')"

	for i := 1; i < 31; i++ {
		key := "article " + string(i)
		value := fmt.Sprintf("INSERT INTO nezha_article(id, tag_id, title,  content, cover_image_url, created_at, created_by, updated_at, state) VALUES (%d, 1, '%s',  'test1-content', '', '2019-08-19 21:00:39', 'test-created', '2019-08-19 21:00:39', 0);", i, "title-"+strconv.Itoa(i))
		mockData[key] = value
	}

	for i := 1; i < 31; i++ {
		key := "application " + string(i)
		value := fmt.Sprintf("INSERT INTO nezha_application(id, name, cn_name,  description, git, jenkins, project_id, user_id, state) VALUES (%d,  'application-%s' , 'app-cn-name-%s',  'description', 'http://git', 'http://jenkins', 1, 1, 0);", i, strconv.Itoa(i), strconv.Itoa(i))
		mockData[key] = value
	}

	for i := 3; i < 31; i++ {
		key := "user- " + string(i)
		value := fmt.Sprintf("INSERT INTO nezha_user(id, username, password, avatar,user_type, state, created_at) VALUES (%d, 'user-%s', '%s', 'https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG', 2, 1, '2020-04-06 13:01:09');", i, strconv.Itoa(i), pw)
		mockData[key] = value
	}

}
