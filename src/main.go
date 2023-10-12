package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/rs/cors"
	"net-http/myapp/crontab"
	"net-http/myapp/route"
	"net/http"
	"os"
)

func main() {
	// crontabでジョブの実行
	// cronスケジューラを用いて定期的に実行するジョブをスタートさせます
	crontab.ToStartCron()

	// 環境変数の読み込み
	// ".env" ファイルから環境変数を読み込みます

	// ルーターの設定
	router := route.Router{Mutex: http.NewServeMux()}
	router.GetRouter()
	router.GetAuthRouter()

	// CORS (Cross Origin Resource Sharing)の設定
	// アクセスを許可するドメイン等を設定します
	corsOrigin := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("ALLOW_ORIGIN")},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	////全てを許可する Access-Control-Allow-Origin: *
	//corsOrigin := cors.Default()
	handler := corsOrigin.Handler(router.Mutex)
	// Webサーバー起動時のエラーハンドリング
	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}

	// AWS Lambdaとの連携設定
	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}
