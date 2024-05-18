package global

import (
	"context"
	"server-api/global/internal/database"
	inlocker "server-api/global/internal/locker"
	"strings"
	"time"

	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"

	"github.com/save95/go-pkg/framework/dbmanager"
	"github.com/save95/go-pkg/framework/dbutil"
	"github.com/save95/go-utils/locker"
	"github.com/save95/xerror"
)

func Database() dbmanager.IDatabaseManager {
	return database.Database()
}

var (
	RedisClient *redis.Client

	SessionStoreClient *redis.Client
)

var (
	Locker       locker.ILocker
	CacheManager *cache.Cache
)

func InitDataBase() error {
	if err := connectDB(); nil != err {
		return xerror.Wrap(err, "database init failed")
	}

	if err := initLocker(); nil != err {
		return xerror.Wrap(err, "locker init failed")
	}

	if err := initRedis(); nil != err {
		return xerror.Wrap(err, "redis init failed")
	}

	if err := initCache(); nil != err {
		return xerror.Wrap(err, "cache init failed")
	}

	// 初始化数据
	if err := initData(); nil != err {
		return xerror.Wrap(err, "data builder init failed")
	}

	return nil
}

func connectDB() error {
	if !Config.Database.Enabled {
		Log.Debug("database disabled, skip")
		return nil
	}

	for _, db := range Config.Database.Connects {
		c, err := dbutil.Connect(&dbutil.Option{
			Name: db.Name,
			Config: &dbutil.ConnectConfig{
				Dsn:         db.Dsn,
				Driver:      db.Driver,
				MaxIdle:     db.MaxIdle,
				MaxOpen:     db.MaxOpen,
				LogMode:     db.LogMode,
				MaxLifeTime: db.MaxLifeTime,
			},
			Logger: Log,
		})
		if err != nil {
			return err
		}

		if err := Database().Register(db.Name, c); nil != err {
			return err
		}
	}

	return nil
}

func initRedis() error {
	if !Config.Redis.Enabled {
		Log.Debug("redis disabled, skip")
		return nil
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     Config.Redis.Addr,
		Password: Config.Redis.Password,
		DB:       Config.Redis.DB,
	})
	if err := RedisClient.Ping(context.Background()).Err(); nil != err {
		return xerror.Wrap(err, "redis client connect failed")
	}

	SessionStoreClient = redis.NewClient(&redis.Options{
		Addr:     Config.Redis.Addr,
		Password: Config.Redis.Password,
		DB:       10,
	})
	if err := SessionStoreClient.Ping(context.Background()).Err(); nil != err {
		return xerror.Wrap(err, "redis client connect failed")
	}

	Log.Debug("redis enabled, init success")

	return nil
}

func initCache() error {
	if !Config.Cache.Enabled {
		Log.Debug("cache disabled, skip")
		return nil
	}

	// 获得不同驱动的存储
	var stored store.StoreInterface
	switch Config.Cache.Drive {
	case "redis":
		cnf := Config.Cache.Redis
		if len(cnf.Addr) == 0 || !strings.Contains(cnf.Addr, ":") {
			return xerror.New("cache redis config not exist")
		}

		stored = store.NewRedis(redis.NewClient(&redis.Options{
			Addr:     Config.Cache.Redis.Addr,
			Password: Config.Cache.Redis.Password,
			DB:       Config.Cache.Redis.DB,
		}), nil)
	default:
		return xerror.New("cache drive not support")
	}

	cacheManager := cache.New(stored)

	// 设置测试缓存
	if err := cacheManager.Set(context.Background(), "cacheMangerTest", "test cache", &store.Options{
		Expiration: 10 * time.Minute, // Override default value of 10 seconds defined in the store
	}); nil != err {
		return xerror.Wrap(err, "cache manager failed")
	}

	CacheManager = cacheManager
	Log.Debug("cache manger init ... success")
	return nil
}

func initLocker() error {
	if !Config.Locker.Enabled {
		Log.Debug("locker disabled, skip")
		return nil
	}

	var (
		err  error
		lock locker.ILocker
	)
	switch Config.Locker.Drive {
	case "redis":
		cnf := Config.Locker.Redis
		lock, err = inlocker.RedisLocker(&redis.Options{
			Addr:     cnf.Addr,
			Password: cnf.Password,
			DB:       cnf.DB,
		})
	default:
		return xerror.New("locker drive not support")
	}
	if nil != err {
		return err
	}

	Locker = lock
	Log.Debug("locker enabled, init success")

	return nil
}

func initData() error {
	if !Config.Database.AutoMigrate {
		Log.Debug("database auto migrate disabled, skip")
		return nil
	}

	if err := database.Migrate(Database()); nil != err {
		return err
	}

	return nil
}
