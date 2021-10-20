package api

import (
	"fmt"
	"inherited/internal/models"
	"inherited/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAdminUser(ctx *gin.Context) {

	//Parameter parsing
	adminUser := models.SysUser{}
	err := ctx.BindJSON(&adminUser)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], Parameter parsing error")
	}

	sysUser := new(services.SysUser)
	err = sysUser.CreateSysUser(adminUser.UserName, adminUser.Password)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], err: %v", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"date": "success",
	})

	return
}