package main

import (
	"fmt"
	"inherited/internal"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化依赖模块
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	var r *gin.Engine
	r = gin.Default()

	// to solve the cross domain
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})

	// setup listen
	err := r.Run(":8000")
	if err != nil {
		fmt.Printf("run failed: %v\n", err)
		return

	}
}
