package handlers

type tokenRequest struct {
	AccessToken  string `json:"token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type Request struct {
	GUID int `json:"guid" binding:"required"`
}
