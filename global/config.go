package global

// 项目配置
type projectConfig struct {
	Server struct {
		AppID string `toml:"app_id"`
		Addr  string `toml:"addr"`
		Host  string `toml:"host"`

		Swagger struct {
			Enabled bool // 是否启用 Swagger
		} // Swagger 配置
	} `toml:"server"` // 服务配置

	Log           logConfig      // 日志配置
	Database      dBConfig       // 数据库配置
	ElasticSearch esConfig       // es配置
	Locker        lockerConfig   // locker 配置
	Cache         cacheConfig    // cache 配置
	Redis         redisConfig    // redis 配置
	App           appConfig      // App
	Job           jobConfig      // JOB
	Listener      listenerConfig // Listener
}

type logConfig struct {
	Dir              string // 存储目录路径
	Category         string // 日志分类目录
	Level            string // 日志级别
	StdPrint         bool   `toml:"std_print"`           // 是否打印到控制台
	HttpLog          bool   `toml:"http_log"`            // 是否打印 http 日志
	HttpLogOnlyError bool   `toml:"http_log_only_error"` // 是否仅打印 http 错误日志
}

type appConfig struct {
	Env              string         // 系统环境: prod/production-生产环境，local-本地环境
	ClearExampleFile bool           `toml:"clear_example_file"` // 是否自动删除样例文件
	Secret           string         // 密钥：jwt 认证等
	Resource         resourceConfig // 资源配置
	Admin            adminConfig    // 管理后台配置
}

type jobConfig struct {
	Enabled bool // 是否启用
}

type listenerConfig struct {
	Enabled bool // 是否启用
}

type dBConfig struct {
	Enabled     bool // 是否启用
	AutoMigrate bool `toml:"auto_migrate"`

	Connects []dBConnectConfig `toml:"connects"`
}

type dBConnectConfig struct {
	Name        string // 连接名称
	Dsn         string // 连接
	Driver      string `toml:"type"`          // 数据库类型
	MaxIdle     int    `toml:"max_idle"`      // 最大空闲连接数
	MaxOpen     int    `toml:"max_open"`      // 最大连接数
	LogMode     bool   `toml:"log_mode"`      // 是否打印SQL
	MaxLifeTime int    `toml:"max_life_time"` // 连接存活时间
}

type esConfig struct {
	Urls         []string
	SniffEnabled bool
	DebugEnabled bool
}

type lockerConfig struct {
	Enabled bool
	Drive   string
	Redis   redisConfig
}

type cacheConfig struct {
	Enabled bool
	Drive   string
	Redis   redisConfig
}

type redisConfig struct {
	Enabled        bool
	Addr           string // 地址
	Password       string // 密码
	Db             int    // 数据库
	Idle           int    // 最大连接数
	Active         int    // 一次性活跃
	Wait           bool   // 是否等待空闲连接
	ConnectTimeout int64  `toml:"connect_timeout"` // 连接超时时间， 毫秒
}

type resourceConfig struct {
	Host           string // 资源域名
	Path           string // 资源上传目录
	ExaminationDir string `toml:"examination_dir"` // 教务资源存储目录
}

type adminConfig struct {
	Account  string // 管理员帐号
	Password string // 管理员密码
}
