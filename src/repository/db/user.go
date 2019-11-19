package db

import (
	"github.com/jinzhu/gorm"
	connection "example_app/util/helper/mysqlconnection"
	dbEntity "example_app/entity/db"
	"sync"
)

type UserRepository struct {
	DB gorm.DB
}

func UserRepositoryHandler() *UserRepository {
	return &UserRepository{DB: *connection.GetConnection()}
}

type UserRepositoryInterface interface {
	GetUserByID(id int, userData *dbEntity.User, wg *sync.WaitGroup) error
	UpdateUserByID(id int, userData *dbEntity.User) error
	GetUsersList(limit int, offset int) ([]dbEntity.User, error)
}

func (repository *UserRepository) GetUserByID(id int, userData *dbEntity.User, wg *sync.WaitGroup) error {
	query := repository.DB.Preload("UserStatus")
	query = query.Where("id=?", id)
	query = query.First(userData)
	wg.Done()
	return query.Error
}

func (repository *UserRepository) UpdateUserByID(id int, userData *dbEntity.User) error {
	query := repository.DB.Table("users")
	query = query.Where("id=?", id)
	query = query.Updates(userData)
	query.Scan(&userData)
	return query.Error
}

func (repository *UserRepository) GetUsersList(limit int, offset int) ([]dbEntity.User, error) {
	users := []dbEntity.User{}
	// user := &dbEntity.User{}
	query := repository.DB.Table("users")
	query = query.Limit(limit).Offset(offset)
	query = query.Find(&users)
	return users, query.Error
}