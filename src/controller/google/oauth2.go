package google

import (
	"encoding/json"
	"net-http/myapp/utils"
	"net-http/myapp/utils/google"
	"net/http"
)

func GoogleOauth(w http.ResponseWriter, r *http.Request) {
	// GETメソッド以外受け付けない
	header := w.Header()
	header.Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(utils.MyError{Message: "Method Not Allowed"})
		return
	}
	config := google.GetConnect()

	url := config.AuthCodeURL("")

	json.NewEncoder(w).Encode(utils.MyError{Message: url})

	//w.Header().Set("location", url)
	//w.WriteHeader(http.StatusMovedPermanently) // 301 Moved Permanently
}
