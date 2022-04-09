package repo

import "ewallet/storage/models"

type Repo interface {
	AddUser(u models.User)(*models.User, error)
	CheckUserById(id string) (bool,error)
}
