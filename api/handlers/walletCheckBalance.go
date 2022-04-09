package handlers

import (
	"encoding/json"
	"ewallet/storage/models"
	"io"
	"net/http"
)

// @Summary Wallet Check Balance
// Description Вы можете узнать. Ваше чоте сколько осталось балансе
// @Security Digest
// @Accept json
// @Produce json
// @Param body body models.Wallet true "wallet check balance"
// @Success 200 {object} models.Wallet
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /wallet/balance [post]
func (h Handlers) WalletCheckBalance(w http.ResponseWriter, r *http.Request) {
	wallet := models.Wallet{}

	walletBate, err := io.ReadAll(r.Body)
	if err != nil {
		WriteError(w, "412 PAGE NOT FAUND", http.StatusNotFound, err)
		return
	}

	err = json.Unmarshal(walletBate, &wallet)
	if err != nil {
		WriteError(w, "error while binding json", http.StatusBadRequest, err)
		return
	}

	walletCheck, err := h.repo.WalletCheckExists(wallet)
	if walletCheck.Id == "" || err != nil {
		WriteError(w, "no such wallet", http.StatusBadRequest, err)
		return
	}

	walletBate, err = json.Marshal(walletCheck)
	if err != nil {
		WriteError(w, "something went wrong, please try agan", http.StatusInternalServerError, err)
		return
	}

	bodyHex := h.auth.HashBody(walletBate)
	r.Header.Set("X-Digest", bodyHex)

	w.Header().Set("Content-Type", "application/json")
	w.Write(walletBate)
}
