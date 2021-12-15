package cache

import (
	"context"
	"time"

	"server-api/global"
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

func (s *user) Paginate(ctx context.Context, opt pager.Option) (*dbcache.Paginate, error) {
	return dbcache.NewDefault(s.name, global.CacheManager).
		WithExpiration(10*time.Minute). // 默认5分钟过去，这里重新修改为10分钟过期
		Paginate(ctx, opt, func(ctx context.Context, opt pager.Option) (interface{}, uint, error) {
			return dao.NewUser().Paginate(opt)
		})
}

func (s *user) First(ctx context.Context, id uint) (interface{}, error) {
	if id == 0 {
		return nil, xerror.New("id error")
	}

	return dbcache.NewDefault(s.name, global.CacheManager).
		First(ctx, id, func(ctx context.Context, id uint) (interface{}, error) {
			return dao.NewUser().First(id)
		})
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
