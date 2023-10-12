package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"net-http/myapp/route"
	"net-http/myapp/utils/cors"
	"net/http"
)

func main() {
	// ルーターの設定
	router := route.Router{Mutex: http.NewServeMux()}
	router.GetRouter()
	router.GetAuthRouter()

	// CORS (Cross Origin Resource Sharing)の設定
	// アクセスを許可するドメイン等を設定します
	corsOrigin := cors.NewCorOrigin()

	handler := corsOrigin.Handler(router.Mutex)
	// Webサーバー起動時のエラーハンドリング => localhostの時コメントイン必要
	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}

	// AWS Lambdaとの連携設定
	lambda.Start(httpadapter.NewV2(handler).ProxyWithContext)
}
