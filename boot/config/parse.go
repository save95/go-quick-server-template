package config

import (
	"io/ioutil"
	"os"
	"strings"

	"server-api/global"

	"github.com/save95/xerror"

	"github.com/pkg/errors"
	"github.com/save95/go-pkg/constant"
	"github.com/save95/go-utils/fsutil"
)

var configFilename = ""

func Init(filename string) error {
	localname := strings.ReplaceAll(constant.ExampleConfigFilename, ".example.", ".")
	if len(filename) == 0 {
		// 如果文件不存在，自动复制
		if !fsutil.Exist(localname) {
			if !fsutil.Exist(constant.ExampleConfigFilename) {
				return errors.New("配置模板文件不存在")
			}

			if _, err := fsutil.Copy(constant.ExampleConfigFilename, localname); nil != err {
				return errors.Wrap(err, "初始化db文件失败")
			}
		}
	} else {
		// 如果是远程连接，则从远程下载
		if strings.HasPrefix(filename, "https://") {
			if err := fsutil.Download(localname, filename); nil != err {
				return xerror.Wrapf(err, "get config from remote failed, url=%s", filename)
			}
		} else {
			localname = filename
		}
	}

	configFilename = filename

	bs, err := ioutil.ReadFile(localname)
	if nil != err {
		return err
	}
	if err := global.ParseConfig(bs); err != nil {
		return err
	}

	// 清理配置
	if global.Config.App.ClearExampleFile {
		_ = os.Remove(constant.ExampleConfigFilename)
	}
	if global.Config.App.ClearConfigFile {
		_ = os.Remove(localname)
	}

	return nil
}
