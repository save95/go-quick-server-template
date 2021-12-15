package db

import (
	"strings"

	"server-api/global"

	"github.com/go-redis/redis/v8"
	"github.com/save95/go-utils/locker"
	"github.com/save95/xerror"
)

func initLocker() error {
	if !global.Config.Locker.Enabled {
		global.Log.Debug("locker disabled, skip")
		return nil
	}

	var (
		err  error
		lock locker.ILocker
	)
	switch global.Config.Locker.Drive {
	case "redis":
		lock, err = _redisLocker()
	default:
		return xerror.New("locker drive not support")
	}
	if nil != err {
		return err
	}

	global.Locker = lock
	global.Log.Debug("locker enabled, init success")

	return nil
}

func _redisLocker() (locker.ILocker, error) {
	cnf := global.Config.Locker.Redis
	if len(cnf.Addr) == 0 || !strings.Contains(cnf.Addr, ":") {
		return nil, xerror.New("locker redis config not exist")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cnf.Addr,
		Password: cnf.Password,
		DB:       3,
	})

	return locker.NewDistributedRedisLock(client), nil
}
