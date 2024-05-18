package config

import (
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/save95/go-utils/fsutil"
	"github.com/save95/xerror"
)

// ParseConfig 解析配置
func ParseConfig(content []byte) (*ProjectConfig, error) {
	var cfg ProjectConfig
	if _, err := toml.Decode(string(content), &cfg); nil != err {
		return nil, err
	}

	return &cfg, nil
}

const exampleConfigFilename = "config/config.example.toml" // APP 配置样例文件

func GetConfigFilename(filename string) (string, error) {
	localFilename := strings.ReplaceAll(exampleConfigFilename, ".example.", ".")
	if len(filename) == 0 {
		if fsutil.Exist(localFilename) {
			return localFilename, nil
		}

		// 如果文件不存在，自动复制
		if !fsutil.Exist(exampleConfigFilename) {
			return "", xerror.New("配置模板文件不存在")
		}

		if _, err := fsutil.Copy(exampleConfigFilename, localFilename); nil != err {
			return "", xerror.Wrap(err, "复制配置模板失败")
		}

		return localFilename, nil
	}

	// 如果是远程连接，则从远程下载
	if strings.HasPrefix(filename, "https://") || strings.HasPrefix(filename, "http://") {
		if err := fsutil.Download(localFilename, filename); nil != err {
			return "", xerror.Wrapf(err, "get config from remote failed, url=%s", filename)
		}
		return localFilename, nil
	}

	return filename, nil
}

func ClearConfigExampleFile() {
	_ = os.Remove(exampleConfigFilename)
}
