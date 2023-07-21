package global

type InitConfig struct {
	ConfigFilename  string
	RegisterServers []InitServerType

	CMDConfig *CMDConfig
}

type InitServerType int

const (
	InitServerTypeWeb InitServerType = iota
	InitServerTypeCronjob
	InitServerTypeListener
)

type CMDConfig struct {
	Name    string
	Timeout int
	Args    string
}
