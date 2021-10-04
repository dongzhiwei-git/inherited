package dao

import (
	"fmt"
	"inherited/internal/conf"

	"github.com/jinzhu/gorm"
)

func Init() (err error) {
	// 初始化db

	conn := conf.Get().Mysql
	DB, err := gorm.Open("mysql", conn.My.Dns)
	if err != nil {
		fmt.Printf("failed to connecte mysql: %v", err)
		return
	}
	return DB.DB().Ping()

	// 初始化oss

	return
}
