package db

import (
	"server-api/service/lang"
)

func initLang() error {
	return lang.Init()
}
