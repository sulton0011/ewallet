package handlers

import (
	"encoding/json"
	"ewallet/storage/models"
	"io"
	"net/http"
)

// @Summary Check Wallet Existanse
// @Description Вы можете узнать что есть или нет кашелек
// @Security Digest
// @Accept json
// @Produce json
// @Param body body models.Wallet true "wallet check existance"
// @Success 200 {object} models.Wallet
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /check/wallet/exist [post]
func (h Handlers) WalletCheckExists(w http.ResponseWriter, r *http.Request) {
	wallet := models.Wallet{}

	walletByte, err := io.ReadAll(r.Body)
	if err != nil {
		WriteError(w, "412 PAGE NOT FAUND", http.StatusNotFound, err)
		return
	}

	err = json.Unmarshal(walletByte, &wallet)
	if err != nil {
		WriteError(w, "400 error while bidnding json", http.StatusBadRequest, err)
		return
	}

	res, err := h.repo.WalletCheckExists(wallet)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError, err)
		return
	}
	if res.Balance == 0 && res.Id == "" {
		WriteMessage(w, "such wallet does not exist")
		return
	}

	walletByte, err = json.Marshal(res)
	if err != nil {
		WriteError(w, "something went wrong,please try again", http.StatusInternalServerError, err)
		return
	}

	historyHex := h.auth.HashBody(walletByte)
	r.Header.Set("X-Digest", historyHex)

	w.Header().Set("Content-Type", "application/json")
	w.Write(walletByte)
}
