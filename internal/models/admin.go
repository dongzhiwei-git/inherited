package models

type SysUser struct {
	Id       int    `gorm:"id"`
	UserName string `gorm:"user_name"`
	Password string `gorm:"password"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
