package db

import (
	"server-api/global"
	"server-api/repository/platform"

	"github.com/pkg/errors"
	"github.com/save95/go-utils/userutil"
)

type dataBuilder struct {
}

func (id dataBuilder) Init() error {
	if err := id.migrate(); nil != err {
		return err
	}

	if err := id.initUser(); nil != err {
		return err
	}

	return nil
}

func (id dataBuilder) migrate() error {
	return global.DbPlatform.AutoMigrate(
		&platform.User{}, &platform.UserLoginLog{},
	)
}

func (id dataBuilder) initUser() error {
	admin := global.Config.App.Admin
	if len(admin.Account) == 0 || len(admin.Password) < 6 {
		return errors.New("account or password error")
	}

	pwd, err := userutil.NewHasher().Sum(admin.Password)
	if nil != err {
		return errors.Wrap(err, "make user password failed")
	}

	// 初始化用户
	global.DbPlatform.FirstOrCreate(&platform.User{
		ID:       1,
		Genre:    uint8(global.RoleAdmin),
		Account:  admin.Account,
		Nickname: "boss",
		Password: pwd,
		State:    1,
	})

	return nil
}
