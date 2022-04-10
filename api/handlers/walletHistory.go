package handlers

import (
	"encoding/json"
	"ewallet/storage/models"
	"io"
	"net/http"
)

// @Summary Get Wallet history
// @Security Digest
// @Accept json
// @Produce json
// @Param body body models.Wallet true "wallet history"
// @Success 200 {object} models.WalletHistory
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /wallet/history [post]
func (h Handlers) WalletHistory(w http.ResponseWriter, r *http.Request) {
	walletReq := models.Wallet{}

	walletByte, err := io.ReadAll(r.Body)
	if err != nil {
		WriteError(w, "412 PAGE NOT FAUND", http.StatusNotFound, err)
		return
	}

	err = json.Unmarshal(walletByte, &walletReq)
	if err != nil {
		WriteError(w, "error while binding json", http.StatusBadRequest, err)
		return
	}

	wallet, err := h.repo.WalletCheckExists(walletReq)
	if wallet.Id == "" || err != nil {
		WriteError(w, "no such wallet", http.StatusBadRequest, err)
		return
	}

	history, err := h.repo.WalletHistory(walletReq)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError, err)
		return
	}
	historyByte, err := json.Marshal(history)
	if err != nil {
		WriteError(w, "something went wrong, please try agan", http.StatusInternalServerError, err)
		return
	}

	bodyHex := h.auth.HashBody(historyByte)
	r.Header.Set("X-Digest", bodyHex)

	w.Header().Set("Content-Type", "application/json")
	w.Write(historyByte)
}
