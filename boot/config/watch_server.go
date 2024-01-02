package config

import (
	"io/ioutil"

	"server-api/global"

	"github.com/save95/go-pkg/application"
	"github.com/save95/xerror"

	"github.com/fsnotify/fsnotify"
)

type server struct {
	watcher *fsnotify.Watcher
}

func NewWatchServer() application.IApplication {
	return &server{}
}

func (s *server) Start() error {
	var err error
	if s.watcher, err = fsnotify.NewWatcher(); err != nil {
		return xerror.New("config file watch failed")
	}

	filename := configFilename

	go func() {
		for {
			select {
			case event, ok := <-s.watcher.Events:
				if !ok {
					return
				}

				// 是否监听
				if !global.Config.App.WatchConfigEnabled {
					global.Log.Debugf("config file(%s) %s, but watch disabled. skip", event.Name, event.Op)
					return
				}

				// 配置文件被修改，则更新全局配置
				if event.Name == filename && event.Op == fsnotify.Write {
					bs, err := ioutil.ReadFile(filename)
					if nil != err {
						global.Log.Errorf("config file(%s) %s, but read failed. err=%+v", event.Name, event.Op, err)
					}
					if err := global.ParseConfig(bs); nil != err {
						global.Log.Errorf("config file(%s) parse failed, err=%+v", event.Name, err)
					}
				}
			case err, ok := <-s.watcher.Errors:
				if !ok {
					return
				}
				global.Log.Errorf("config file watch error: %+v", err)
			}
		}
	}()

	if err := s.watcher.Add(filename); nil != err {
		return err
	}

	global.Log.Info("config watch server starting...")
	return nil
}

func (s *server) Shutdown() error {
	return s.watcher.Close()
}
