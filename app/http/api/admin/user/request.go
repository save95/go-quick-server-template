package user

import (
	"github.com/save95/go-pkg/http/types"
	"github.com/save95/xerror"
)

type paginateRequest struct {
	types.SearchRequest

	Account string `form:"account"`
}

type createRequest struct {
	Genres   []uint8 `json:"genres"`
	Account  string  `json:"account"`
	IsBoss   bool    `json:"isBoss"`
	Avatar   string  `json:"avatar"`
	Password string  `json:"password"`
}

func (in *createRequest) GetGenres() []int8 {
	res := make([]int8, 0)
	for _, genre := range in.Genres {
		if genre == 0 {
			continue
		}
		res = append(res, int8(genre))
	}
	return res
}

func (in *createRequest) Validate() error {
	if len(in.Account) == 0 || len(in.Genres) == 0 {
		return xerror.New("帐号、类型 不能为空")
	}

	if len(in.Password) == 0 {
		return xerror.New("密码不能为空")
	}

	return nil
}

type modifyRequest struct {
	Account   string `json:"account"`
	IsBoss    bool   `json:"isBoss"`
	AvatarURL string `json:"avatar"`
	Password  string `json:"password"`
	State     int8   `json:"state"`
}

func (in *modifyRequest) Validate() error {
	return nil
}
