package repositorye

type AuthRedisRepository interface {
	GetTokens(userId int) (accessToken, refreshToken string, err error)
	SetTokens(accessToken, refreshToken string, userId int) error
}
