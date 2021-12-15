package db

import "github.com/pkg/errors"

func Connect() error {
	if err := initMysql(); nil != err {
		return errors.Wrap(err, "mysql init failed")
	}

	// 初始化数据
	if err := new(dataBuilder).Init(); nil != err {
		return errors.Wrap(err, "data builder init failed")
	}

	if err := initLocker(); nil != err {
		return errors.Wrap(err, "locker init failed")
	}

	if err := initRedis(); nil != err {
		return errors.Wrap(err, "redis init failed")
	}

	if err := initCache(); nil != err {
		return errors.Wrap(err, "cache init failed")
	}

	return nil
}
