package db

import (
	"fmt"
	"strings"

	"server-api/global"

	"github.com/pkg/errors"
	"github.com/save95/go-pkg/constant"
	"github.com/save95/go-pkg/framework/dbutil"
	"github.com/save95/go-utils/fsutil"
)

func initSqlite() error {
	var err error

	// 如果文件不存在，自动复制
	filename := strings.ReplaceAll(constant.ExampleLocalDBFilename, ".example.", ".")
	if !fsutil.Exist(filename) {
		if !fsutil.Exist(constant.ExampleLocalDBFilename) {
			return errors.New("本地数据库模板文件不存在")
		}

		if _, err := fsutil.Copy(constant.ExampleLocalDBFilename, filename); nil != err {
			return errors.Wrap(err, "初始化db文件失败")
		}
	}

	connectStr := fmt.Sprintf("%s?charset=utf8mb4&parseTime=true&loc=Local", filename)
	dbc := global.Config.Database.Platform
	global.DbPlatform, err = dbutil.Connect(&dbutil.Option{
		Name: "platform",
		Config: &dbutil.ConnectConfig{
			Dsn:         connectStr,
			Driver:      "sqlite",
			MaxIdle:     dbc.MaxIdle,
			MaxOpen:     dbc.MaxOpen,
			LogMode:     dbc.LogMode,
			MaxLifeTime: dbc.MaxLifeTime,
		},
		Logger: global.Log,
	})
	if nil != err {
		return errors.Wrap(err, "open db connect failed")
	}

	return nil
}
