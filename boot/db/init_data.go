package db

import (
	"server-api/boot/db/internal"
	"server-api/global"
)

type dataBuilder struct {
}

func (id dataBuilder) Init() error {
	if !global.Config.Database.AutoMigrate {
		global.Log.Debug("database auto migrate disabled, skip")
		return nil
	}

	if err := internal.NewMigrate().Platform(); nil != err {
		return err
	}

	if err := internal.NewInit().Platform(); nil != err {
		return err
	}

	return nil
}
