package postgres

import (
	"ewallet/storage/models"
	"fmt"
	"time"

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
		now = time.Now().Format(time.RFC822)
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

	return  false, nil
}