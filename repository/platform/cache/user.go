package cache

import (
	"context"
	"time"

	"server-api/global"
	"server-api/repository/platform"
	"server-api/repository/platform/dao"

	"github.com/save95/go-pkg/framework/dbcache"
	"github.com/save95/go-pkg/model/pager"
	"github.com/save95/xerror"
	"github.com/zywaited/xcopy"
)

type user struct {
	name string
}

func NewUser() *user {
	return &user{
		name: "user",
	}
}

func (s *user) Paginate(ctx context.Context, opt pager.Option) ([]*platform.User, uint, error) {
	data, err := dbcache.NewDefault(s.name, global.CacheManager).
		WithExpiration(10*time.Minute). // 默认5分钟过去，这里重新修改为10分钟过期
		Paginate(ctx, opt, func() (interface{}, uint, error) {
			return dao.NewUser().Paginate(opt)
		})
	if nil != err {
		return nil, 0, err
	}

	var res []*platform.User
	if err := xcopy.Copy(&res, data.Data); nil != err {
		return nil, 0, xerror.Wrap(err, "data convert error")
	}

	return res, data.Total, nil
}

func (s *user) First(ctx context.Context, id uint) (*platform.User, error) {
	if id == 0 {
		return nil, xerror.New("id error")
	}

	data, err := dbcache.NewDefault(s.name, global.CacheManager).
		First(ctx, id, func() (interface{}, error) {
			return dao.NewUser().First(id)
		})
	if nil != err {
		return nil, err
	}

	var res platform.User
	if err := xcopy.Copy(&res, data); nil != err {
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
