package db

import (
	"server-api/global"

	"github.com/save95/go-pkg/framework/dbutil"
)

func initMysql() error {
	if !global.Config.Database.Enabled {
		global.Log.Debug("database disabled, skip")
		return nil
	}

	var err error
	dbc := global.Config.Database.Platform
	dbPlatform, err := dbutil.Connect(&dbutil.Option{
		Name: "platform",
		Config: &dbutil.ConnectConfig{
			Dsn:         dbc.Dsn,
			Driver:      dbc.Driver,
			MaxIdle:     dbc.MaxIdle,
			MaxOpen:     dbc.MaxOpen,
			LogMode:     dbc.LogMode,
			MaxLifeTime: dbc.MaxLifeTime,
		},
		Logger: global.Log,
	})
	if err != nil {
		return err
	}
	if err := global.Database().Register("platform", dbPlatform); nil != err {
		return err
	}

	return nil
}
