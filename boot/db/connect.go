package db

import (
	"server-api/global"

	"github.com/save95/go-pkg/framework/dbutil"
)

func connect() error {
	if !global.Config.Database.Enabled {
		global.Log.Debug("database disabled, skip")
		return nil
	}

	for _, db := range global.Config.Database.Connects {
		c, err := dbutil.Connect(&dbutil.Option{
			Name: db.Name,
			Config: &dbutil.ConnectConfig{
				Dsn:         db.Dsn,
				Driver:      db.Driver,
				MaxIdle:     db.MaxIdle,
				MaxOpen:     db.MaxOpen,
				LogMode:     db.LogMode,
				MaxLifeTime: db.MaxLifeTime,
			},
			Logger: global.Log,
		})
		if err != nil {
			return err
		}

		if err := global.Database().Register(db.Name, c); nil != err {
			return err
		}
	}

	return nil
}
