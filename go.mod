module server-api

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/eko/gocache/v2 v2.1.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.4
	github.com/go-redis/redis/v8 v8.11.5
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/save95/go-pkg v1.1.2-0.20230517092524-7034f041d111
	github.com/save95/go-utils v1.0.4-0.20221116061429-f98c5aee1649
	github.com/save95/xerror v1.1.1
	github.com/save95/xlog v0.0.1
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.7.0
	github.com/zywaited/xcopy v1.0.1-0.20220105081048-cf992328857e
	gorm.io/gorm v1.21.12
)

//replace github.com/save95/go-pkg => /Users/royee/Develop/PoeticalSoft/save95/go-pkg
