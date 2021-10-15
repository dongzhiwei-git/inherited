package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

type SysUser struct {
}

func (s *SysUser) CreateSysUser(name, password string) {
	var adminUser = models.SysUser{
		UserName: name,
		Password: password,
	}

	dao.Orm.Create(&adminUser)

}
