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

type WalletHistory struct {
	Id                     string `json:"id"`
	CurrentBalance         float64 `json:"current_balance"`
	TotalIncome            float64 `json:"total_income"`
	TotalExpense           float64 `json:"total_expense"`
	TotalIncomeOperations  int64  `json:"total_income_operations"`
	TotalExpenseOperations int64  `json:"tatal_expense_operations"`
	TotalOperations        int64  `json:"total_operations"`
}

type Err struct {
	Error string `json:"error"`
}
