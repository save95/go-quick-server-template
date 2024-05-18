package locker

import (
	"context"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/save95/go-utils/locker"
	"github.com/save95/xerror"
)

func RedisLocker(opt *redis.Options) (locker.ILocker, error) {
	if len(opt.Addr) == 0 || !strings.Contains(opt.Addr, ":") {
		return nil, xerror.New("locker redis config not exist")
	}

	client := redis.NewClient(opt)
	if err := client.Ping(context.Background()).Err(); nil != err {
		return nil, xerror.Wrap(err, "redis client connect failed")
	}

	return locker.NewDistributedRedisLock(client), nil
}
