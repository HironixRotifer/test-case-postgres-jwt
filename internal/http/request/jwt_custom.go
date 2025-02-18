package request

// FilterOrders model info
// @Description Описание модели 'JwtCustom'
type JwtCustom struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
