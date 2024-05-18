package internal

import "github.com/save95/go-pkg/framework/dbmanager"

type IDatabase interface {
	Platform(dbs dbmanager.IDatabaseManager) error
}
