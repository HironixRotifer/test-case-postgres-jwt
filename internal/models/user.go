package models

type User struct {
	UID          int    `json:"uid"`
	Email        string `json:"email"`
	IP           string `json:"ip"`
	RefreshToken string `json:"refresh_token"`
}
