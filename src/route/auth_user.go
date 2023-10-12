package route

import (
	"net-http/myapp/controller"
	"net-http/myapp/controller/google"
)

func (router *Router) GetAuthRouter() {
	router.Mutex.HandleFunc("/api/v1/auth", controller.HandlerTwo)
	router.Mutex.HandleFunc("/api/v1/google/login", google.GoogleOauth)
	router.Mutex.HandleFunc("/api/v1/google/callback", google.GoogleLoginHandler)

	// コメント アクション

	// ユーザーのコメント一覧
}
