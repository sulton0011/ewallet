package handlers

import (
	"encoding/json"
	"ewallet/storage/models"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

// @Summary New User
// @здесь вы можете создать и добавить нового пользователя
// @Accept json
// @Produce json
// @Param body body models.User true "новый пользователь"
// @Success 200 {object} models.User
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /user [post]
func (h Handlers) AddUser(w http.ResponseWriter, r *http.Request) {
	bodyReq := models.User{}

	s, err := io.ReadAll(r.Body)
	if err != nil {
		WriteError(w, "412 PAGE NOT FAUND", http.StatusPreconditionFailed, err)
		return
	}

	if err = json.Unmarshal(s, &bodyReq); err != nil {
		WriteError(w, "400 error while bidnding json", http.StatusBadRequest, err)
		return
	}

	bodyReq.Id = uuid.New().String()
	user, err := h.repo.AddUser(bodyReq)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError, err)
		return
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		WriteError(w, "something went wrong,please try again", http.StatusInternalServerError, err)
		return
	}

	historyHex := h.auth.HashBody(bytes)
	fmt.Println(historyHex, r.URL.Path)

	r.Header.Set("X-Digest", historyHex)
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)

}
