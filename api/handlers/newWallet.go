package handlers

import (
	"encoding/json"
	"ewallet/storage/models"
	"io"
	"net/http"

	"github.com/google/uuid"
)

// @Summary New Wallet
// @Description здесь вы можете создать нового кошелек
// @Accept json
// @Produce json
// @Param body body models.NewWallet true "новый кошелек"
// @Success 200 {object} models.Wallet
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /wallet/new [post]
func (h Handlers) NewWallet(w http.ResponseWriter, r *http.Request) {
	bodyReq := models.NewWallet{}

	byteReq, err := io.ReadAll(r.Body)
	if err != nil {
		WriteError(w, "412 PAGE NOT FAUND", http.StatusPreconditionFailed, err)
		return
	}

	err = json.Unmarshal(byteReq, &bodyReq)
	if err != nil {
		WriteError(w, "400 error while bidnding json", http.StatusBadRequest, err)
		return
	}

	bodyReq.WalletId = uuid.New().String()

	wallet, err := h.repo.NewWallet(bodyReq)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError, err)
		return
	}

	walletByte, err := json.Marshal(wallet)
	if err != nil {
		WriteError(w, "something went wrong,please try again", http.StatusInternalServerError, err)
		return
	}

	historyHex := h.auth.HashBody(walletByte)

	r.Header.Set("X-Digest", historyHex)
	w.Header().Set("Content-Type", "application/json")
	w.Write(walletByte)
}
