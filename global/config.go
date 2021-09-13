package global

// 项目配置
type projectConfig struct {
	Log           logConfig    // 日志配置
	Server        serverConfig // 服务配置
	Database      *dBConfig    // 数据库配置
	ElasticSearch *esConfig    // es配置
	App           appConfig    // App
	Job           jobConfig    // JOB
}

type logConfig struct {
	Dir              string // 存储目录路径
	Category         string // 日志分类目录
	Level            string // 日志级别
	StdPrint         bool   `toml:"std_print"`           // 是否打印到控制台
	HttpLog          bool   `toml:"http_log"`            // 是否打印 http 日志
	HttpLogOnlyError bool   `toml:"http_log_only_error"` // 是否仅打印 http 错误日志
}

type serverConfig struct {
	Addr    string        `toml:"addr"`
	Swagger swaggerConfig // Swagger 配置
}

type swaggerConfig struct {
	// 是否启用 Swagger
	Enabled bool
}

type appConfig struct {
	ClearExampleFile bool           `toml:"clear_example_file"` // 是否自动删除样例文件
	Resource         resourceConfig // 资源配置
	Admin            adminConfig    // 管理后台配置
}

type jobConfig struct {
	Enable bool // 是否启用
}

type dBConfig struct {
	Platform dBConnectConfig `toml:"platform"`
}

type dBConnectConfig struct {
	Dsn         string // 连接
	Driver      string `toml:"type"`          // 数据库类型
	MaxIdle     int    `toml:"max_idle"`      // 最大空闲连接数
	MaxOpen     int    `toml:"max_open"`      // 最大连接数
	LogMode     bool   `toml:"log_mode"`      // 是否打印SQL
	MaxLifeTime int    `toml:"max_life_time"` // 连接存活时间
}

type esConfig struct {
	Urls        []string
	SniffEnable bool
	DebugEnable bool
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
