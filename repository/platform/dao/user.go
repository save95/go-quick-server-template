package dao

import (
	"fmt"

	"server-api/global"
	"server-api/repository/platform"

	"github.com/pkg/errors"

	"github.com/save95/go-pkg/model/pager"

	"github.com/save95/xerror"
	"github.com/save95/xerror/xcode"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func NewUser(options ...interface{}) *user {
	impl := user{}
	for _, option := range options {
		if db, ok := option.(*gorm.DB); ok {
			impl.db = db
		}
	}

	if impl.db == nil {
		impl.db, _ = global.Database().Get("platform")
	}

	return &impl
}

func (u *user) Create(record *platform.User) error {
	if record.ID > 0 {
		return xerror.WithXCode(xcode.DBRequestParamError)
	}

	if err := u.db.Create(record).Error; nil != err {
		return xerror.WrapWithXCode(err, xcode.DBFailed)
	}

	return nil
}

func (u *user) Save(record *platform.User) error {
	if record.ID == 0 {
		return xerror.WithXCode(xcode.DBRecordNotFound)
	}

	if err := u.db.Save(record).Error; nil != err {
		return xerror.WrapWithXCode(err, xcode.DBFailed)
	}

	return nil
}

func (u *user) First(id uint) (*platform.User, error) {
	if id == 0 {
		return nil, xerror.WithXCode(xcode.DBRequestParamError)
	}

	var record platform.User
	if err := u.db.Where("id = ?", id).First(&record).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerror.WithXCode(xcode.DBRecordNotFound)
		}
		return nil, xerror.WrapWithXCode(err, xcode.DBFailed)
	}

	return &record, nil
}

func (u *user) FirstByAccount(genre uint8, account string) (*platform.User, error) {
	if genre == 0 || len(account) == 0 {
		return nil, xerror.WithXCode(xcode.DBRequestParamError)
	}

	db := u.db.Model(platform.User{}).
		Where("genre = ?", genre).
		Where("account = ?", account)

	var record platform.User
	if err := db.First(&record).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerror.WithXCode(xcode.DBRecordNotFound)
		}

		return nil, xerror.WrapWithXCode(err, xcode.DBFailed)
	}

	return &record, nil
}

func (u *user) Paginate(option pager.Option) ([]*platform.User, uint, error) {
	db := u.db.Model(platform.User{})

	// 账号
	if v, ok := option.Filter["account"]; ok {
		if vv, ok := v.(string); ok && len(vv) > 0 {
			vv = fmt.Sprintf("%%%s%%", vv)
			db = db.Where("account like ?", vv)
		}
	}

	var total int64
	_ = db.Count(&total).Error

	var records []*platform.User
	if err := db.Order("id DESC").
		Offset(option.Start).Limit(option.GetLimit()).
		Find(&records).Error; nil != err {
		return nil, 0, xerror.WrapWithXCode(err, xcode.DBFailed)
	}

	return records, uint(total), nil
}
