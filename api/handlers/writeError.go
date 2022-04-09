package handlers

import (
	"encoding/json"
	"ewallet/storage/models"
	"fmt"
	"net/http"
)

func WriteError(w http.ResponseWriter, msg string, code int, err error) {
	body := models.Err{
		Error: msg,
	}
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := json.Marshal(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(bytes)
}

func WriteMessage(w http.ResponseWriter, msg string) {
	type message struct {
		Message string `json:"message"`
	}

	body := message{Message: msg}

	bytes, _ := json.Marshal(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(bytes)
}
