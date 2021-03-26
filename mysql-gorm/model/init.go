package model

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Self *gorm.DB
	// Docker *gorm.DB
}

var DB *Database

// openDB() 函数调用 gorm.Open() 来建立一个数据库连接
func openDB(username, password, addr, dbname string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?character_set_server=utf8&parseTime=%t&loc=%s", username, password, addr, dbname, true, "Local")
	db, err := gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		logrus.Errorf("%q. Database connection failed. Database name: %s", err, dbname)
	}

	setupDB(db)
	return db
}

// 为 mysql 连接设置高级参数（连接池等）
func setupDB(db *gorm.DB) {
	// db.LogMode(viper.GetBool("gormlog"))
	advanceDB, err := db.DB()
	if err != nil {
		logrus.Fatal("Database connection pools set failed", err)
	}
	// 连接池配置
	// SetMaxIdleConns设置空闲连接池中连接的最大数量
	advanceDB.SetMaxIdleConns(0)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	advanceDB.SetMaxOpenConns(10)
	// SetConnMaxLifetime 设置连接可复用的最大时间。
	advanceDB.SetConnMaxLifetime(time.Hour)
}

func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

// 调用 openDB() 方法创建 Database的数据库对象
func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

// func InitDockerDB() *gorm.DB {
// 	return openDB(viper.GetString("docker_db.username"),
// 		viper.GetString("docker_db.password"),
// 		viper.GetString("docker_db.addr"),
// 		viper.GetString("docker_db.name"))
// }

// func GetDockerDB() *gorm.DB {
// 	return InitDockerDB()
// }

// 初始化连接：Init()-->GetSelfDB()-->InitSelfDB()-->openDB()<--setupDB()
func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
		// Docker: GetDockerDB(),
	}
}

// 关于 Close() 函数：对于大多数应用 v1 版本也不需要 Close，用连接池就好了，好多人用错，于是 v2 版本直接把 Close 方法去掉了，直接用连接池不需要考虑 Close，刚需 Close 的人翻文档。
