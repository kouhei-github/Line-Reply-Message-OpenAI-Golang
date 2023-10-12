package google

import (
	"context"
	"encoding/json"
	v2 "google.golang.org/api/oauth2/v2"
	"net-http/myapp/controller"
	"net-http/myapp/utils"
	"net-http/myapp/utils/google"
	"net/http"
)

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	// GETメソッド以外受け付けない
	header := w.Header()
	header.Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(utils.MyError{Message: "Method Not Allowed"})
		return
	}

	code := r.URL.Query().Get("code")

	// ユーザーの情報の取得
	config := google.GetConnect()
	ctx := context.Background()
	tok, err := config.Exchange(ctx, code)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(utils.MyError{Message: "Internal Server Error"})
		return
	}

	if tok.Valid() == false {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(utils.MyError{Message: "Internal Server Error"})
		return
	}

	service, _ := v2.New(config.Client(ctx, tok))
	tokenInfo, _ := service.Tokeninfo().AccessToken(tok.AccessToken).Context(ctx).Do()

	response := controller.Response{Text: tokenInfo.Email}

	json.NewEncoder(w).Encode(response)
}
