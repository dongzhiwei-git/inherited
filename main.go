package main

import (
	"fmt"
	"inherited/internal"
	"inherited/internal/dao"
	"inherited/internal/models"
	"inherited/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化依赖模块
	fmt.Printf("...........wew.qeqweqw.e.")
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	dao.Orm.AutoMigrate(models.SysUser{})
	SysUser := new(services.SysUser)
	err := SysUser.CreateSysUser("root", "234")
	if err != nil {
		fmt.Printf("err.........")
	}
	fmt.Printf("err34.........")

	var r *gin.Engine
	r = gin.Default()

	// to solve the cross domain
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})

	// setup listen
	err = r.Run(":8000")
	if err != nil {
		fmt.Printf("run failed: %v\n", err)
		return

	}
}
