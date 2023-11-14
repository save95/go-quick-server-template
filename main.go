package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"server-api/boot"
	"server-api/global"
)

var (
	flagConf, flagMode string
	flagCMDArgs        global.FlagSlice
	flagCMDConf        global.CMDConfig
)

func step() {
	flag.StringVar(&flagConf, "conf", "config/config.toml", "config path, support remote url")
	flag.StringVar(&flagMode, "mode", "all", "server mode: all, web, cron/cronjob, listener, cmd/command")

	flag.StringVar(&flagCMDConf.Name, "cmd.name", "", "command task name, only support `cmd` mode")
	flag.IntVar(&flagCMDConf.Timeout, "cmd.timeout", 0, "command task run timeout, second")
	flag.Var(&flagCMDArgs, "cmd.args", "command task run args. default use `:` split key and value, e.g., `-cmd.args=ver:v1.1028` is key=`ver`, value=`v1.1028`")

	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
		_, _ = fmt.Fprint(os.Stderr, "OPTIONS:\n")
		flag.PrintDefaults()
	}

	flag.Parse()
}

// @title server API
// @version 1.0
// @description 接口文档.

func main() {
	step()

	log.Println("launcher starting...")
	log.Printf("launcher flags: conf=%s, mode=%s\n", flagConf, flagMode)

	cnf := global.InitConfig{
		ConfigFilename:  flagConf,
		RegisterServers: make([]global.InitServerType, 0, 4),
	}

	// command 命令
	if flagMode == "cmd" || flagMode == "command" {
		if len(flagCMDArgs) > 0 {
			flagCMDConf.Args = flagCMDArgs
		}
		cnf.CMDConfig = &flagCMDConf
		if err := boot.Command(cnf); err != nil {
			log.Fatalf("boot command cmd failed: %+v\n", err.Error())
		}
		return
	}

	switch flagMode {
	case "all":
		cnf.RegisterServers = append(
			cnf.RegisterServers,
			global.InitServerTypeWeb,
			global.InitServerTypeCronjob,
			global.InitServerTypeListener,
		)
	case "web":
		cnf.RegisterServers = append(cnf.RegisterServers, global.InitServerTypeWeb)
	case "cron", "cronjob":
		cnf.RegisterServers = append(cnf.RegisterServers, global.InitServerTypeCronjob)
	case "listener":
		cnf.RegisterServers = append(cnf.RegisterServers, global.InitServerTypeListener)
	default:
		log.Fatalf("boot failed: mode err, not in all, web, cron/cronjob, command, listener, is %+v\n", flagMode)
	}

	if err := boot.Boot(cnf); err != nil {
		log.Fatalf("boot failed: %+v\n", err.Error())
	}
}
