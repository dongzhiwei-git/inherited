package internal

import (
	"inherited/internal/dao"

	"github.com/pkg/errors"
)

// 初始化dao
func Init() error {
	if err := dao.Init(); err != nil {
		return errors.WithStack(err)
	}

}
