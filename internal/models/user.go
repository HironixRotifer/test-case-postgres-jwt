package models

type User struct {
	UID      int    `json:"uid"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

type JwtCustom struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
