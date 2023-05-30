package repos

import (
	"Go_CRUD_API/database"
	"Go_CRUD_API/model"
	"errors"
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

func (s *UserStore) UpdateById(id int, u *model.User) (*model.User, error) {
	res := s.DB.Where("id = ?", id).Updates(&u)
	if res.Error != nil {
		log.Println("Error while updating user in DB", id, res.Error)
		return nil, res.Error
	}
	return u, nil
}

func (s *UserStore) Delete(id int) (*model.User, error) {
	usr := &model.User{}
	res := s.DB.Delete(&usr, id)
	if res.Error != nil {
		log.Println("Error while deleting user in DB", id, res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		log.Println("Record with ", id, " not found.")
		return nil, errors.New("not Found")
	}
	return usr, nil
}

func (s *UserStore) GetUsers(q map[string]interface{}) ([]model.User, error) {
	var users []model.User
	res := s.DB.Where(q).Find(&users)
	if res.Error != nil {
		log.Println("Fetching Users list, ", q, res.Error)
		return nil, res.Error
	}
	return users, nil
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
