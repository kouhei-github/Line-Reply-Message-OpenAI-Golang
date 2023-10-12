package cors

import (
	"github.com/rs/cors"
	"net/http"
	"os"
)

func NewCorOrigin() *cors.Cors {
	// CORS (Cross Origin Resource Sharing)の設定
	// アクセスを許可するドメイン等を設定します
	////全てを許可する Access-Control-Allow-Origin: *
	corsOrigin := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("ALLOW_ORIGIN")},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	return corsOrigin
}
