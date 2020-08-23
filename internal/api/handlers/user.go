package handlers

import (
	"../../data"
	"encoding/json"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defaultResponse := data.DefaultResponse{Code: 1000, Status: "OK", Content: []string{"User endpoint"}}
	response, err := json.Marshal(defaultResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(response)
}
