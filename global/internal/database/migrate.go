package database

import (
	"server-api/global/internal/database/internal"

	"github.com/save95/go-pkg/framework/dbmanager"
)

func Migrate(dbs dbmanager.IDatabaseManager) error {
	if err := internal.NewMigrate().Platform(dbs); nil != err {
		return err
	}

	if err := internal.NewInit().Platform(dbs); nil != err {
		return err
	}

	return nil
}
