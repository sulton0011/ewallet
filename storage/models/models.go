package models

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone_number"`
}

type NewWallet struct {
	UserId   string `json:"user_id"`
	WalletId string `json:"wallet_id"`
}

type Wallet struct {
	Id      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type WalletFill struct {
	Id     string  `json:"id"`
	UserId string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type Err struct {
	Error string `json:"error"`
}
