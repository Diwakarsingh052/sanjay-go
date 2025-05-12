package stores

import (
	"app/internal/stores/models"
)

type DataBase interface {
	Create(u models.User) (models.User, bool)
	//CreateSimple(string) (models.User, bool)
	Update(int, string) (models.User, bool)
	Delete(int) bool
	FetchAll() (map[int]models.User, bool)
	FetchUser(int) (models.User, bool)
}
