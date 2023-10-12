package route

import (
	"net-http/myapp/controller"
	"net-http/myapp/controller/line"
	"net/http"
)

type Router struct {
	Mutex *http.ServeMux
}

func (router *Router) GetRouter() {
	//router.Mutex.HandleFunc("/", controller.HandlerTwo)
	// 練習
	router.Mutex.HandleFunc("/two", controller.HandlerTwo)

	// line-web-hook
	router.Mutex.HandleFunc("/api/v1/webhook", line.MessageHandler)

	// 記事をDBに保存

	// 人気の記事一覧

	// タグの管理
}
