package dao

import (
	"fmt"

	"server-api/global"
	"server-api/repository/platform"
	"server-api/repository/types/platformtypes"

	"github.com/pkg/errors"

	"github.com/save95/go-pkg/model/pager"

	"github.com/save95/xerror"
	"github.com/save95/xerror/xcode"
	"gorm.io/gorm"
)

type vwUser struct {
	db *gorm.DB
}

func NewVWUser(options ...interface{}) *vwUser {
	impl := vwUser{}
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

func (u *vwUser) First(id uint, preloads ...string) (*platform.VWUser, error) {
	if id == 0 {
		return nil, xerror.WithXCode(xcode.DBRequestParamError)
	}

	db := u.db.Where("id = ?", id)

	for _, preload := range preloads {
		db = db.Preload(preload)
	}

	var record platform.VWUser
	if err := db.First(&record).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerror.WithXCode(xcode.DBRecordNotFound)
		}
		return nil, xerror.WrapWithXCode(err, xcode.DBFailed)
	}

	return &record, nil
}

func (u *vwUser) FirstByAccount(account string, preloads ...string) (*platform.VWUser, error) {
	if len(account) == 0 {
		return nil, xerror.WithXCode(xcode.DBRequestParamError)
	}

	db := u.db.Model(platform.VWUser{}).
		Where("account = ?", account)

	for _, preload := range preloads {
		db = db.Preload(preload)
	}

	var record platform.VWUser
	if err := db.First(&record).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerror.WithXCode(xcode.DBRecordNotFound)
		}

		return nil, xerror.WrapWithXCode(err, xcode.DBFailed)
	}

	return &record, nil
}

// FirstThirdOpenID 获取三方平台用户 unionID/openID
// 返回：unionID, openID, error
func (u *vwUser) FirstThirdOpenID(third platformtypes.UserFromPlatform, appID string, userID uint) (string, string, error) {
	switch third {
	case platformtypes.UserFromPlatformWechat:
		var record platform.WechatUser
		if err := u.db.Model(platform.WechatUser{}).
			Where("app_id = ?", appID).
			Where("user_id = ?", userID).
			First(&record).Error; nil != err {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return "", "", xerror.WrapWithXCodeStatus(err, xcode.DBRecordNotFound)
			}
			return "", "", xerror.WrapWithXCodeStatus(err, xcode.DBFailed)
		}
		return record.UnionID, record.OpenID, nil
	//case platformtypes.UserFromPlatformAlipay:
	default:
		return "", "", xerror.New("not supported")
	}
}

func (u *vwUser) List(option pager.Option) ([]*platform.VWUser, error) {
	db := u.build(option.Filter)

	for _, preload := range option.GetPreloads() {
		db = db.Preload(preload)
	}

	var records []*platform.VWUser
	if err := u.order(db, option.Sorters).Order("id DESC").
		Offset(option.Start).Limit(option.GetLimit()).
		Find(&records).Error; nil != err {
		return nil, xerror.WrapWithXCode(err, xcode.DBFailed)
	}

	return records, nil
}

func (u *vwUser) Paginate(option pager.Option) ([]*platform.VWUser, uint, error) {
	db := u.build(option.Filter)

	var total int64
	_ = db.Count(&total).Error

	for _, preload := range option.GetPreloads() {
		db = db.Preload(preload)
	}

	var records []*platform.VWUser
	if err := u.order(db, option.Sorters).Order("id DESC").
		Offset(option.Start).Limit(option.GetLimit()).
		Find(&records).Error; nil != err {
		return nil, 0, xerror.WrapWithXCode(err, xcode.DBFailed)
	}

	return records, uint(total), nil
}

func (u *vwUser) build(filter pager.Filter) *gorm.DB {
	db := u.db.Model(platform.VWUser{})

	// 账号
	if v, ok := filter["account"]; ok {
		if vv, ok := v.(string); ok && len(vv) > 0 {
			vv = fmt.Sprintf("%%%s%%", vv)
			db = db.Where("account like ?", vv)
		}
	}

	return db
}

func (u *vwUser) order(db *gorm.DB, sorters []pager.Sorter) *gorm.DB {
	if sorters == nil {
		return db
	}

	for _, sorter := range sorters {
		switch sorter.Field {
		case "created_at":
			db = db.Order(fmt.Sprintf("%s %s", sorter.Field, sorter.Sorted))
		}
	}

	return db
}

func (u *vwUser) ListRoles(id uint) ([]*platform.UserRole, error) {
	db := u.db.Model(platform.UserRole{}).
		Where("user_id = ?", id)

	var records []*platform.UserRole
	if err := db.Order("id ASC").Find(&records).Error; nil != err {
		return nil, xerror.WrapWithXCode(err, xcode.DBFailed)
	}

	return records, nil
}
