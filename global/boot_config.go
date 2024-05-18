package global

import (
	"fmt"
)

type BootConfig struct {
	ConfigFilename  string
	RegisterServers []BootServerType

	CMDConfig *CMDBootConfig
}

func (m BootConfig) LogCategory() string {
	if len(m.RegisterServers) > 1 {
		return ""
	}
	if m.CMDConfig != nil {
		return "command"
	}

	switch m.RegisterServers[0] {
	case InitServerTypeCronjob:
		return "cronjob"
	case InitServerTypeListener:
		return "listener"
	default:
		return ""
	}
}

type BootServerType int

const (
	InitServerTypeWeb BootServerType = iota
	InitServerTypeCronjob
	InitServerTypeListener
)

type CMDBootConfig struct {
	Name    string
	Timeout int
	Args    []string
}

type FlagSlice []string

func (f *FlagSlice) String() string {
	return fmt.Sprintf("%v", []string(*f))
}

func (f *FlagSlice) Set(value string) error {
	*f = append(*f, value)
	return nil
}
