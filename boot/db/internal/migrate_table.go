package internal

import (
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

	return dbPlatform.AutoMigrate(
		&platform.User{}, &platform.UserLoginLog{}, &platform.Lang{},
	)
}
