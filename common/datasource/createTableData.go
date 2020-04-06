package datasource

import (
	"fmt"

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
	mockData["adminUser"] = fmt.Sprintf("INSERT INTO nezha_user(id, username, password, avatar,user_type, created_at) VALUES "+
		" (1, 'admin', '%s', 'https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG',1, '2020-04-06 13:01:09');",
		pw)
	mockData["testUser"] = fmt.Sprintf("INSERT INTO nezha_user(id, username, password, avatar,user_type, created_at) VALUES "+
		" (2, 'test', '%s', 'https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG',2, '2020-04-06 13:01:09');",
		pw)

	mockData["tag1"] = "INSERT INTO nezha_tag(id, name, created_at, created_by, updated_at, updated_by, deleted_at, state) VALUES" +
		" (1, '1', '2019-08-18 18:56:01', 'test', NULL, '', 0, 1);"
	mockData["tag2"] = "INSERT INTO nezha_tag(id, name, created_at, created_by, updated_at, updated_by, deleted_at, state) VALUES" +
		" (2, '2', '2019-08-18 18:56:01', 'test', NULL, '', 0, 1);"
	mockData["tag3"] = "INSERT INTO nezha_tag(id, name, created_at, created_by, updated_at, updated_by, deleted_at, state) VALUES" +
		" (3, '3', '2019-08-18 18:56:01', 'test', NULL, '', 0, 1);"

	mockData["role1"] = "INSERT INTO nezha_role(id, user_id, user_name, value) VALUES (1, 1, 'admin', 'admin')"
	mockData["role2"] = "INSERT INTO nezha_role(id, user_id, user_name, value) VALUES (2, 1, 'admin', 'test')"
	mockData["role3"] = "INSERT INTO nezha_role(id, user_id, user_name, value) VALUES (3, 2, 'test', 'test')"

	mockData["article1"] = "INSERT INTO nezha_article(id, tag_id, title,  content, cover_image_url, created_at, created_by, updated_at, state) VALUES (1, 1, 'test1',  'test1-content', '', '2019-08-19 21:00:39', 'test-created', '2019-08-19 21:00:39', 0);"
	mockData["article2"] = "INSERT INTO nezha_article(id, tag_id, title,  content, cover_image_url, created_at, created_by, updated_at, state) VALUES (2, 1, 'test2',  'test2-content', '', '2019-08-19 21:00:39', 'test-created', '2019-08-19 21:00:39', 0);"
	mockData["article3"] = "INSERT INTO nezha_article(id, tag_id, title,  content, cover_image_url, created_at, created_by, updated_at, state) VALUES (3, 1, 'test3',  'test3-content', '', '2019-08-19 21:00:39', 'test-created', '2019-08-19 21:00:39', 0);"

}
