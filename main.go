package main

import (
	"log"
	"server-api/boot"
)

// @title server API
// @version 1.0
// @description 接口文档.

func main() {
	log.Println("launcher starting...")
	if err := boot.Boot(); err != nil {
		log.Fatalf("boot failed: %+v\n", err.Error())
	}
}
