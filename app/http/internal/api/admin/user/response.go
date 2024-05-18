package user

import (
	"github.com/save95/go-pkg/http/types"
)

type entity struct {
	types.ResponseModel

	CheckedGenre  uint8  `json:"checkedGenre"`
	Account       string `json:"account"`
	Nickname      string `json:"nickname"`
	Gender        uint8  `json:"gender"`
	GenderText    string `json:"genderText" copy:"Gender.String"`
	AvatarURL     string `json:"avatarUrl" copy:"ShowAvatarURL"`
	HasPassword   bool   `json:"hasPassword"`
	State         int8   `json:"state"`
	StateText     string `json:"stateText"`
	LastLoginTime string `json:"lastLoginTime" copy:"LastLoginAt"`
	LastLoginIP   string `json:"lastLoginIp" copy:"LastLoginIP"`
	LastDriverNo  string `json:"lastDriverNo"`
	Genres        []int8 `json:"genres" copy:"Genres"`
}
