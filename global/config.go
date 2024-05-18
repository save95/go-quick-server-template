package global

import (
	"io/ioutil"
	"log"
	"os"
	"server-api/global/internal/config"
)

var Config config.ProjectConfig

func GetConfigFilename(filename string) (string, error) {
	return config.GetConfigFilename(filename)
}

// ParseConfig 解析配置
func ParseConfig(filename string) error {
	localFilename, err := config.GetConfigFilename(filename)
	if err != nil {
		return err
	}

	bs, err := ioutil.ReadFile(localFilename)
	if nil != err {
		return err
	}
	cnf, err := config.ParseConfig(bs)
	if err != nil {
		return err
	}

	Config = *cnf

	// 清理文件
	if Config.App.ClearExampleFile {
		config.ClearConfigExampleFile()
	}
	if Config.App.ClearConfigFile {
		if Config.App.WatchConfigEnabled {
			log.Println("set clear config file, but enabled watch it, skip delete")
			return nil
		}
		_ = os.Remove(localFilename)
	}
	return nil
}
