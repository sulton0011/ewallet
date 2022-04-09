package postgres

import (
	"ewallet/storage/models"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

// Этот метод для создания нового пользователя
func (d Database) AddUser(u models.User) (*models.User, error) {
	// запрос для не идентифицированных пользователей
	query := `INSERT INTO users (user_id, first_name, last_name, email, created_at) 
	VALUES ( $1, $2, $3, $4, $5 )`

	// запрос для идентифицированных пользователей, которые предоставляют номер телефона
	queryIden := `INSERT INTO users (user_id, first_name, last_name, email, phone, is_detected, created_at) 
	VALUES ( $1, $2, $3, $4, $5, $6, $7 )`

	var (
		now = time.Now().UTC().Format("2006-01-02 15:04:05")
		err error
	)

	// Первое утверждение для неидентифицированных пользователей, второе для идентифицированных
	if u.Phone == "" {
		err = d.db.QueryRow(query, u.Id, u.FirstName, u.LastName, u.Email, now).Err()
	} else {
		err = d.db.QueryRow(queryIden, u.Id, u.FirstName, u.LastName, u.Email, u.Phone, true, now).Err()
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}

// Этот метод проверяет, есть ли у нас такой пользователь или нет, если нет, то результат ложный
func (d Database) CheckUserById(id string) (bool, error) {
	var count int
	query := `SELECT COUNT(1) FROM users WHERE user_id = $1 AND deleted_at IS NULL`
	if err := d.db.QueryRow(query, id).Scan(&count); err != nil {
		fmt.Println(err)
		return false, err
	}

	if count == 0 {
		return true, nil
	}
	return false, nil
}

// Этот метод создает новый кошелек
func (d Database) NewWallet(w models.NewWallet) (*models.Wallet, error) {
	isIden, err := d.isUserIdentified(w.UserId)
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO wallets (wallet_id, user_id, is_detected, created_at)
	VALUES ($1, $2, $3, $4)`

	err = d.db.QueryRow(
		query,
		w.WalletId,
		w.UserId,
		isIden,
		time.Now().UTC().Format("2006-01-02 15:04:05"),
	).Err()
	if err != nil {
		return nil, err
	}

	return &models.Wallet{Id: w.WalletId, Balance: 0}, nil
}

// Этот метод проверяет, идентифицирован ли пользователь или нет
func (d Database) isUserIdentified(id string) (bool, error) {
	query := `SELECT is_detected FROM users WHERE user_id = $1`
	var isUserIdentified bool

	if err := d.db.QueryRow(query, id).Scan(&isUserIdentified); err != nil {
		return false, err
	}
	return isUserIdentified, nil
}

// Этот метод проверяет, существует ли кошелек или нет, если нет, он возвращает ошибку и кошелек с нулевыми полями
func (d Database) WalletCheckExists(w models.Wallet) (*models.Wallet, error) {
	var count int
	query := `SELECT COUNT(*) FROM wallets WHERE wallet_id = $1 AND deleted_at IS NULL`

	err := d.db.QueryRow(query, w.Id).Scan(&count)
	if err != nil {
		return nil, err
	}

	if count != 1 {
		return &models.Wallet{Id: "", Balance: 0}, err
	}

	return d.WalletCheckBalance(w)
}

// Этот метод возвращает баланс кошелька
func (d Database) WalletCheckBalance(w models.Wallet) (*models.Wallet, error) {
	query := `SELECT balance FROM wallets WHERE wallet_id = $1 AND deleted_at IS NULL`

	err := d.db.QueryRow(query, w.Id).Scan(&w.Balance)
	if err != nil {
		return nil, err
	}

	return &w, nil
}

// Этот метод используется для заполнения кошелька
func (d Database) WalletFill(w models.WalletFill) (*models.Wallet, error) {
	isIden, err := d.isUserIdentified(w.UserId)
	if err != nil {
		return nil, err
	}

	wallet, err := d.WalletCheckBalance(models.Wallet{Id: w.Id})
	if err != nil {
		return nil, err
	}

	sumBalance := wallet.Balance + w.Amount

	if !isIden {
		if sumBalance > 10000 {
			return nil, fmt.Errorf("you are not identified user, so your limit is 10 000")
		}
	}

	if sumBalance > 100000 {
		return nil, fmt.Errorf("your limit is 100 000")
	}

	tx, err := d.db.Begin()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC().Format("2006-01-02 15:04:05")

	query := `UPDATE wallets SET balance = balance + $2, updated_at = $3 
		WHERE wallet_id = $1 AND deleted_at IS NULL`

	_, err = tx.Exec(query, w.Id, w.Amount, now)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	query = `INSERT INTO wallets_incomes (history_id, wallet_id, amount, created_at)
		VALUES ($1, $2, $3, $4)`

	_, err = tx.Exec(query, uuid.New().String(), w.Id, w.Amount, now)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return d.WalletCheckBalance(models.Wallet{Id: w.Id})
}
