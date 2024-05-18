package internal

import (
	"server-api/global/internal/database/internal/platform"

	"github.com/save95/go-pkg/framework/dbmanager"

	"github.com/pkg/errors"
	"github.com/save95/go-utils/userutil"
)

type initData struct {
}

func NewInit() IDatabase {
	return &initData{}
}

//func (m *initData) Platform(dbs dbmanager.IDatabaseManager, admin, password string) error {
func (m *initData) Platform(dbs dbmanager.IDatabaseManager) error {
	dbPlatform, err := dbs.Get("platform")
	if nil != err {
		return err
	}

	admin := "admin"
	password := "password"
	genre := uint8(1)
	if len(admin) == 0 || len(password) < 6 {
		return errors.New("account or password error")
	}

	pwd, err := userutil.NewHasher().Sum(password)
	if nil != err {
		return errors.Wrap(err, "make user password failed")
	}

	// 初始数据
	datas := []interface{}{
		&platform.User{
			ID:      1,
			Account: admin,
			//CheckedGenre: uint8(global.RoleAdmin),
			CheckedGenre: genre,
			Nickname:     "boss",
			Password:     pwd,
			State:        1,
		},
		&platform.UserRole{
			ID:     1,
			UserID: 1,
			//Genre:  uint8(global.RoleAdmin),
			Genre: genre,
		},
	}
	for _, data := range datas {
		if err := dbPlatform.FirstOrCreate(data).Error; nil != err {
			return err
		}
	}

	return nil
}
