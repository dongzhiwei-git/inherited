package test

import (
	"inherited/internal"
	"inherited/internal/dao"
	"inherited/internal/models"
	"inherited/internal/services"
	"log"
	"testing"
)

func TestCreateSysUser(t *testing.T) {

	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	defer dao.Orm.Close()
	//DB, err := gorm.Open("mysql", conn.Master.Dsn)
	//if err != nil {
	//	fmt.Printf("failed to connecte mysql: %v", err)
	//	return
	dao.Orm.AutoMigrate(models.SysUser{})

	SysUser := new(services.SysUser)
	SysUser.CreateSysUser("root", "234")

}
