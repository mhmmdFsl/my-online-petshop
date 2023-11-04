package model

type LoginRs struct {
	User         *User  `json:"user"`
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}
