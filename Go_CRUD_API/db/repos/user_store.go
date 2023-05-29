package repos

import (
	"github.com/fardinabir/Go_CRUD_API/database"
	"github.com/fardinabir/Go_CRUD_API/model"
	"gorm.io/gorm"
	"log"
)

type UserStore struct {
	DB *gorm.DB
}

func NewUserStore() *UserStore {
	d := database.DatabaseConnection()
	return &UserStore{DB: d}
}

func (s *UserStore) Create(u *model.User) error {
	res := s.DB.Create(u)
	if res.Error != nil {
		log.Println("Error while creating user in DB", res.Error)
		return res.Error
	}
	return nil
}

func (s *UserStore) Update(u *model.User) error {
	return nil
}

func (s *UserStore) UpdateById(u *model.User) error {
	return nil
}

func (s *UserStore) GetUsers(q map[string]interface{}) ([]model.User, error) {
	return nil, nil
}

func (s *UserStore) GetUserByName(name string) (*model.User, error) {
	usr := &model.User{}
	res := s.DB.Where("user_name = ?", name).First(usr)
	if res.Error != nil {
		log.Println("Error while getting user in DB", name, res.Error)
		return nil, res.Error
	}
	return usr, nil
}

func (s *UserStore) GetUserById(id int) (*model.User, error) {
	usr := &model.User{}
	res := s.DB.Where("id = ?", id).First(usr)
	if res.Error != nil {
		log.Println("Error while getting user in DB", id, res.Error)
		return nil, res.Error
	}
	return usr, nil
}
