package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerTwo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello　World")
}

type Files struct {
	OriginalFileName string `json:"OriginalFileName"`
	OutputFileName   string `json:"OutputFileName"`
}

type Body struct {
	RequestBody  Files             `json:"request_body"`
	Replacements map[string]string `json:"replacements"`
	Insertions   map[string]string `json:"insertions"`
}

func BodyTestHandler(w http.ResponseWriter, r *http.Request) {
	var body Body
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"message": "存在しないURLです"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		json.NewEncoder(w).Encode(map[string]string{"status": "500", "message": "Internal Server Error: "})
		return
	}

	fmt.Println("my event: ", body)
	json.NewEncoder(w).Encode(body.Insertions)
}
