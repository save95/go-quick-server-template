package config

import (
	"strings"

	"server-api/global"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"github.com/save95/go-pkg/constant"
	"github.com/save95/go-utils/fsutil"
)

func Init() error {
	// 如果文件不存在，自动复制
	filename := strings.ReplaceAll(constant.ExampleConfigFilename, ".example.", ".")
	if !fsutil.Exist(filename) {
		if !fsutil.Exist(constant.ExampleConfigFilename) {
			return errors.New("配置模板文件不存在")
		}

		if _, err := fsutil.Copy(constant.ExampleConfigFilename, filename); nil != err {
			return errors.Wrap(err, "初始化db文件失败")
		}
	}

	_, err := toml.DecodeFile(filename, &global.Config)
	if err != nil {
		return err
	}
	return nil
}
