package user

import "github.com/save95/go-pkg/http/types"

type entity struct {
	types.ResponseModel

	Genre       uint8  `json:"genre"`
	Account     string `json:"account"`
	IsBoss      bool   `json:"isBoss"`
	IsAi        bool   `json:"isAi"`
	Avatar      string `json:"avatar"`
	State       int8   `json:"state"`
	LastLoginAt string `json:"lastLoginAt"`
	LastLoginIp string `json:"lastLoginIp"`
	DriverNo    string `json:"driverNo"`
}
