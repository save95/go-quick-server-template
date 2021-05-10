package auth

import (
	"server-api/global"
)

type createTokenRequest struct {
	Genre    global.Role `json:"genre"`
	Account  string      `json:"account"`
	Password string      `json:"password"`
}
