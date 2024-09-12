package models

type User struct {
	UID   string `json:"uid"`
	Email string `json:"email"`
	IP    string `json:"ip"`
}
