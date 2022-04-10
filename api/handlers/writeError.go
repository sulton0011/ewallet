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
	if err != nil && err.Error() !=  "you are not identified user, so your limit is 10 000" && err.Error() != "your limit is 100 000"{
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

	errByte, _ := json.Marshal(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(errByte)
}
