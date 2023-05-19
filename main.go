package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"server-api/boot"
	"server-api/global"
)

var flagConf, flagMode string

func init() {
	flag.StringVar(&flagConf, "conf", "config/config.toml", "config path, support remote url")
	flag.StringVar(&flagMode, "mode", "all", "server mode: all, web, job, listener")

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
	log.Println("launcher starting...")
	log.Printf("launcher flags: conf=%s, mode=%s\n", flagConf, flagMode)

	cnf := global.InitConfig{
		ConfigFilename:  flagConf,
		RegisterServers: make([]global.InitServerType, 0, 4),
	}

	switch flagMode {
	case "all":
		cnf.RegisterServers = append(
			cnf.RegisterServers,
			global.InitServerTypeWeb,
			global.InitServerTypeJob,
			global.InitServerTypeListener,
		)
	case "web":
		cnf.RegisterServers = append(cnf.RegisterServers, global.InitServerTypeWeb)
	case "job":
		cnf.RegisterServers = append(cnf.RegisterServers, global.InitServerTypeJob)
	case "listener":
		cnf.RegisterServers = append(cnf.RegisterServers, global.InitServerTypeListener)
	default:
		log.Fatalf("boot failed: mode err, not in all, web, job, listener, is %+v\n", flagMode)
	}

	if err := boot.Boot(cnf); err != nil {
		log.Fatalf("boot failed: %+v\n", err.Error())
	}
}
