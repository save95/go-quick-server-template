package user

type entity struct {
	ID          uint   `json:"id" copy:"ID"`
	Genre       uint8  `json:"genre"`
	Account     string `json:"account"`
	IsBoss      bool   `json:"isBoss"`
	IsAi        bool   `json:"isAi"`
	Avatar      string `json:"avatar"`
	State       int8   `json:"state"`
	LastLoginAt string `json:"lastLoginAt"`
	LastLoginIp string `json:"lastLoginIp"`
	DriverNo    string `json:"driverNo"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
