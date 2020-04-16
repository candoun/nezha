package datasource

import (
	"fmt"
	"log"
	"database/sql"

	"github.com/aguncn/nezha/common/setting"
	"github.com/aguncn/nezha/models"
	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Migrate bool
)

//Db gormDB
type Db struct {
	Conn *gorm.DB
}


func addLocation(dbURL string) string {
        // https://stackoverflow.com/questions/30074492/what-is-the-difference-between-utf8mb4-and-utf8-charsets-in-mysql
        return fmt.Sprintf("%s?charset=utf8mb4&loc=%s", dbURL, "Asia%2FShanghai")
}



//Connect 初始化数据库配置
func (d *Db) Connect() error {
	var (
		dbType, dbName, user, pwd, host string
	)

	conf := setting.Config.Database
	dbType = conf.Type
	dbName = conf.Name
	user = conf.User
	pwd = conf.Password
	host = conf.Host

	DBTns := fmt.Sprintf("tcp(%s)", host)
    dbURL := fmt.Sprintf("%s:%s@%s/", user, pwd, DBTns)
	dbForCreateDatabase, err := sql.Open(dbType, addLocation(dbURL))
	fmt.Print("111connecting create db: xxxxxxxxxxxxxxxxxxxxxx",err)
	fmt.Print("\n")
	_, err = dbForCreateDatabase.Exec(fmt.Sprintf("CREATE DATABASE `%s` CHARACTER SET utf8 COLLATE utf8_general_ci;", dbName))
	fmt.Print("connecting create db: xxxxxxxxxxxxxxxxxxxxxx",err)
	fmt.Print("\n")

	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, dbName))
	if err != nil {
		log.Fatal("connecting mysql error: ", err)
		return err
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.TablePrefix + defaultTableName
	}
	db.LogMode(true) //打印SQL语句
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	d.Conn = db

	log.Println("Connect Mysql Success.")

	if Migrate {
		d.DropAndCreateTb()
		log.Println("Drop And Create Table...")
		if err := CreateTableData(db); err != nil {
			log.Fatal("CreateTableData error: ", err)
			return err
		}
		log.Println("Create Table...")
	}

	return nil
}

//DB 返回DB
func (d *Db) DB() *gorm.DB {
	return d.Conn
}

func (d *Db) DropAndCreateTb() {
	d.DropAndCreateSingleTb(&models.Application{})
	d.DropAndCreateSingleTb(&models.Article{})
	d.DropAndCreateSingleTb(&models.Project{})
	d.DropAndCreateSingleTb(&models.Role{})
	d.DropAndCreateSingleTb(&models.Tag{})
	d.DropAndCreateSingleTb(&models.User{})
}

func (d *Db) DropAndCreateSingleTb(Tb interface{}) {
	d.Conn.DropTableIfExists(Tb)
	d.Conn.CreateTable(Tb)

}
