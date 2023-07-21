package db

import "server-api/app/service/lang"

func initLang() error {
	return lang.Init()
}
