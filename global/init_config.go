package global

type InitConfig struct {
	ConfigFilename  string
	RegisterServers []InitServerType
}

type InitServerType int

const (
	InitServerTypeWeb InitServerType = iota
	InitServerTypeJob
	InitServerTypeListener
)
