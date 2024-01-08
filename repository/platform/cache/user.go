package cache

import (
	"context"
	"encoding/json"
	"time"

	"server-api/global"
	"server-api/repository/platform"
	"server-api/repository/platform/dao"

	"github.com/save95/go-pkg/framework/dbcache"
	"github.com/save95/go-pkg/model/pager"
	"github.com/save95/xerror"
)

type user struct {
	name string
}

func NewUser() *user {
	return &user{
		name: "user",
	}
}

func (s *user) Paginate(ctx context.Context, opt pager.Option) ([]*platform.VWUser, uint, error) {
	key := s.name
	data, err := dbcache.NewDefault(key, global.CacheManager).
		WithExpiration(time.Hour). // 默认5分钟，改成1小时缓存
		Paginate(ctx, opt, func() (interface{}, uint, error) {
			// 原始方法
			return dao.NewVWUser().Paginate(opt)
		})
	if nil != err {
		return nil, 0, err
	}

	// 解码
	var res []*platform.VWUser
	if err := json.Unmarshal(data.DataBytes, &res); nil != err {
		return nil, 0, xerror.Wrap(err, "data convert error")
	}

	return res, data.Total, nil
}

func (s *user) First(ctx context.Context, id uint) (*platform.VWUser, error) {
	if id == 0 {
		return nil, xerror.New("id error")
	}

	key := s.name
	data, err := dbcache.NewDefault(key, global.CacheManager).
		WithExpiration(time.Hour). // 默认5分钟，改成1小时缓存
		First(ctx, id, func() (interface{}, error) {
			// // 原始方法
			return dao.NewVWUser().First(id)
		})
	if nil != err {
		return nil, err
	}

	// 解码
	var res platform.VWUser
	if err := json.Unmarshal([]byte(data), &res); nil != err {
		return nil, xerror.Wrap(err, "data convert error")
	}

	return &res, nil
}

func (s *user) ClearAll(ctx context.Context) error {
	return dbcache.NewDefault(s.name, global.CacheManager).ClearAll(ctx)
}

func (s *user) ClearPaginate(ctx context.Context) error {
	return dbcache.NewDefault(s.name, global.CacheManager).ClearPaginate(ctx)
}

func (s *user) ClearFirst(ctx context.Context, id uint) error {
	return dbcache.NewDefault(s.name, global.CacheManager).ClearFirst(ctx, id)
}
