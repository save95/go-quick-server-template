package db

import (
	"fmt"
	"server-api/global"
	"strings"

	"github.com/save95/go-pkg/constant"

	"github.com/save95/go-pkg/utils/fsutil"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
)

func initSqlite() error {
	var err error

	// 如果文件不存在，自动复制
	filename := strings.ReplaceAll(constant.ExampleLocalDBFilename, ".example.", ".")
	if !fsutil.Exist(filename) {
		if !fsutil.Exist(constant.ExampleLocalDBFilename) {
			return errors.New("本地数据库模板文件不存在")
		}

		if _, err := fsutil.Copy(constant.ExampleLocalDBFilename, filename); nil != err {
			return errors.Wrap(err, "初始化db文件失败")
		}
	}

	connectStr := fmt.Sprintf("%s?charset=utf8mb4&parseTime=true&loc=Local", filename)
	db, err := gorm.Open("sqlite3", connectStr)
	if nil != err {
		return errors.Wrap(err, "open db connect failed")
	}
	//defer db.Close()
	db.LogMode(true)

	global.DbPlatform = db

	return nil
}
