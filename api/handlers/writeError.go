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
	fmt.Println(err)

	bytes, _ := json.Marshal(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(bytes)
}
