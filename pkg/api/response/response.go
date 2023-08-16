package response

type Error struct {
	Error string `json:"error"`
}

type Login struct {
	AccessToken     string  `json:"accessToken"`     // 系统jwt
	MiniAccessToken string  `json:"miniAccessToken"` // 小程序access_token
	MiniExpiresIn   float64 `json:"miniExpiresIn"`   // 小程序expires_in
}
