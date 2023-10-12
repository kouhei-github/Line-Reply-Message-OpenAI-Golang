package controller

import (
	"encoding/json"
	"net/http"
)

func HandlerTwo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello　World")
}
