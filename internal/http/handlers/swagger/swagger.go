package swagger

import (
	"net/http"

	_ "github.com/HironixRotifer/test-case-postgres-jwt/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Swagger() http.HandlerFunc {
	return httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)
}
