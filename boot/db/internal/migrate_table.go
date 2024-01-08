package internal

import (
	"fmt"
	"server-api/boot/db/internal/platform"
	"server-api/global"
)

type migrate struct {
}

func NewMigrate() IDatabase {
	return &migrate{}
}

func (m *migrate) Platform() error {
	dbPlatform, err := global.Database().Get("platform")
	if nil != err {
		return err
	}

	tables := map[interface{}]string{
		platform.FailedJob{}:    "失败系统任务记录",
		platform.Lang{}:         "语言包",
		platform.User{}:         "用户表",
		platform.UserLoginLog{}: "用户登录日志",
		platform.UserRole{}:     "用户角色",
		platform.UserStat{}:     "用户统计",
	}

	for table, comment := range tables {
		opt := fmt.Sprintf("COMMENT='%s'", comment)
		if err := dbPlatform.Set("gorm:table_options", opt).AutoMigrate(table); nil != err {
			return err
		}
	}

	for name, sql := range views {
		if err := dbPlatform.Exec(fmt.Sprintf(`DROP VIEW IF EXISTS %s`, name)).Error; nil != err {
			return err
		}

		if err := dbPlatform.Exec(sql).Error; nil != err {
			return err
		}
	}

	return nil
}
