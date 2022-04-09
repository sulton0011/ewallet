package handlers

import (
	"encoding/json"
	"ewallet/storage/models"
	"io"
	"net/http"
)

// @Summary Wallet Fill
// @Description По этой конечной точке вы можете заполнить или пополнить свой кошелек
// @Security Digest
// @Accept json
// @Produce json
// @Param body body models.WalletFill true "fill wallet"
// @Success 200 {object} models.Wallet
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /wallet/fill [post]
func (h Handlers) WalletFill(w http.ResponseWriter, r *http.Request) {
	walletReq := models.WalletFill{}

	walletByte, err := io.ReadAll(r.Body)
	if err != nil {
		WriteError(w, "412 PAGE NOT FAUND", http.StatusNotFound, err)
		return
	}

	err = json.Unmarshal(walletByte, &walletReq)
	if err != nil {
		WriteError(w, "400 error while bidnding json", http.StatusBadRequest, err)
		return
	}
	
	wallet, err := h.repo.WalletCheckExists(models.Wallet{Id: walletReq.Id})
	if wallet.Id == "" || err != nil {
		WriteError(w, "no such wallet, please check it", http.StatusBadRequest, err)
		return
	}

	wallet, err = h.repo.WalletFill(walletReq)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError, err)
		return
	}

	walletByte, err = json.Marshal(wallet)
	if err != nil {
		WriteError(w, "something went wrong, please try agan", http.StatusInternalServerError, err)
		return
	}

	bodyHex := h.auth.HashBody(walletByte)
	r.Header.Set("X-Digest", bodyHex)

	w.Header().Set("Content-Type", "application/json")
	w.Write(walletByte)
}
