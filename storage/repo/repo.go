package repo

import "ewallet/storage/models"

type Repo interface {
	AddUser(u models.User) (*models.User, error)
	CheckUserById(id string) (bool, error)
	NewWallet(w models.NewWallet) (*models.Wallet, error)
	WalletCheckExists(w models.Wallet) (*models.Wallet, error)
	WalletCheckBalance(w models.Wallet) (*models.Wallet, error)
}
