package users

import "github.com/fardinabir/Go_CRUD_API/model"

type UserStore interface {
	Create(u *model.User) error
	Update(u *model.User) error
	UpdateById(u *model.User) error
	GetUsers(q map[string]interface{}) ([]model.User, error)
	GetUserById(id int) (*model.User, error)
	GetUserByName(name string) (*model.User, error)
}
