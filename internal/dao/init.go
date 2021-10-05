package dao

import (
	"fmt"
	"inherited/internal/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Init() (err error) {
	// 初始化db

	conn := conf.Get().Mysql
	//DB, err := gorm.Open("mysql", "root:beego@tcp(121.36.216.191:3306)/inherited?charset=utf8mb4&parseTime=True&loc=Local")
	DB, err := gorm.Open("mysql", conn.Master.Dsn)
	if err != nil {
		fmt.Printf("failed to connecte mysql: %v", err)
		return
	}
	return DB.DB().Ping()

	// 初始化oss

	return
}
