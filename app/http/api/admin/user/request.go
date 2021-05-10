package user

import "github.com/save95/go-pkg/http/types"

type paginateRequest struct {
	types.SearchRequest

	Account string `form:"account"`
}

type createRequest struct {
	Genre    uint8  `json:"genre"`
	Account  string `json:"account"`
	IsBoss   bool   `json:"isBoss"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
}

type modifyRequest struct {
	Account  string `json:"account"`
	IsBoss   bool   `json:"isBoss"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
	State    int8   `json:"state"`
}
