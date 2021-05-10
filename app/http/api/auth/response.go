package auth

type tokenEntity struct {
	AccessToken  string   `json:"accessToken"`
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	ID           uint     `json:"id"`
	Avatar       string   `json:"avatar"`
	Name         string   `json:"name"`
}
