module server-api

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/eko/gocache/v2 v2.1.0
	github.com/fsnotify/fsnotify v1.7.0
	github.com/gin-contrib/sessions v0.0.4
	github.com/gin-gonic/gin v1.7.7
	github.com/go-redis/redis/v8 v8.11.5
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/save95/go-pkg v1.2.1-0.20231012044021-6c2227dfe851
	github.com/save95/go-utils v1.0.4
	github.com/save95/xerror v1.1.3
	github.com/save95/xlog v0.0.1
	github.com/stretchr/testify v1.7.1
	github.com/zywaited/xcopy v1.1.0
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
	gorm.io/gorm v1.21.12
)

//replace github.com/save95/go-pkg => /Users/royee/Develop/PoeticalSoft/save95/go-pkg
