package response

// FilterOrders model info
// @Description Описание модели 'JwtCustom'
type JwtCustom struct {
	AccessToken  string `json:"access_token"`  // Access token
	RefreshToken string `json:"refresh_token"` // Refresh token
}
