package dao

import (
	"fmt"

	"server-api/global"
	"server-api/repository/platform"

	"github.com/save95/go-utils/valutil"

	"github.com/pkg/errors"
	"github.com/save95/go-pkg/model/pager"
	"github.com/save95/xerror"
	"github.com/save95/xerror/xcode"
	"gorm.io/gorm"
)

type lang struct {
	db *gorm.DB
}

func NewLang(options ...interface{}) *lang {
	impl := lang{}
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

func (u *lang) First(id uint) (*platform.Lang, error) {
	if id == 0 {
		return nil, xerror.WithXCode(xcode.DBRequestParamError)
	}

	var record platform.Lang
	if err := u.db.Where("id = ?", id).First(&record).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerror.WithXCode(xcode.DBRecordNotFound)
		}
		return nil, xerror.WrapWithXCode(err, xcode.DBFailed)
	}

	return &record, nil
}

func (u *lang) FirstByCode(code int) (*platform.Lang, error) {
	db := u.db.Model(platform.User{}).
		Where("code = ?", code)

	var record platform.Lang
	if err := db.First(&record).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerror.WithXCode(xcode.DBRecordNotFound)
		}

		return nil, xerror.WrapWithXCode(err, xcode.DBFailed)
	}

	return &record, nil
}

func (u *lang) Paginate(option pager.Option) ([]*platform.Lang, uint, error) {
	db := u.db.Model(platform.Lang{})

	if v, ok := option.Filter["codes"]; ok {
		if vv, ok := v.([]int); ok && len(vv) > 0 {
			db = db.Where("code in (?)", vv)
		}
	}

	if v, ok := option.Filter["code"]; ok {
		if vv, err := valutil.Int(v); err == nil {
			db = db.Where("code = ?", vv)
		}
	}

	var total int64
	_ = db.Count(&total).Error

	var records []*platform.Lang
	if err := u.order(db, option.Sorters).Order("id DESC").
		Offset(option.Start).Limit(option.GetLimit()).
		Find(&records).Error; nil != err {
		return nil, 0, xerror.WrapWithXCode(err, xcode.DBFailed)
	}

	return records, uint(total), nil
}

func (u *lang) order(db *gorm.DB, sorters []pager.Sorter) *gorm.DB {
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
