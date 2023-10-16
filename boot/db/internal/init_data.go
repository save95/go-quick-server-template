package internal

import (
	"server-api/boot/db/internal/platform"
	"server-api/global"

	"github.com/pkg/errors"
	"github.com/save95/go-utils/userutil"
)

type initData struct {
}

func NewInit() IDatabase {
	return &initData{}
}

func (m *initData) Platform() error {
	dbPlatform, err := global.Database().Get("platform")
	if nil != err {
		return err
	}

	admin := global.Config.App.Admin
	if len(admin.Account) == 0 || len(admin.Password) < 6 {
		return errors.New("account or password error")
	}

	pwd, err := userutil.NewHasher().Sum(admin.Password)
	if nil != err {
		return errors.Wrap(err, "make user password failed")
	}

	// 初始数据
	datas := []interface{}{
		&platform.User{
			ID:           1,
			Account:      admin.Account,
			CheckedGenre: uint8(global.RoleAdmin),
			Nickname:     "boss",
			Password:     pwd,
			State:        1,
		},
		&platform.UserRole{
			ID:     1,
			UserID: 1,
			Genre:  uint8(global.RoleAdmin),
		},
	}
	for _, data := range datas {
		if err := dbPlatform.FirstOrCreate(data).Error; nil != err {
			return err
		}
	}

	return nil
}
