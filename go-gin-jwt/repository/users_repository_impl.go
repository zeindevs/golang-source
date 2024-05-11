package repository

import (
	"errors"
	"ginjwt/data/request"
	"ginjwt/helper"
	"ginjwt/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepositoryImpl(db *gorm.DB) UsersRepository {
	return &UserRepositoryImpl{Db: db}
}

// Delete implements UsersRepository.
func (u *UserRepositoryImpl) Delete(userId int) {
	var users model.Users
	result := u.Db.Where("id = ?", userId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UsersRepository.
func (u *UserRepositoryImpl) FindAll() []model.Users {
	var users []model.Users
	result := u.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindById implements UsersRepository.
func (u *UserRepositoryImpl) FindById(usersId int) (model.Users, error) {
	var users model.Users
	result := u.Db.Find(&users, usersId)
	if result != nil {
		return users, nil
	} else {
		return users, errors.New("user is not found")
	}
}

// FindByUsername implements UsersRepository.
func (u *UserRepositoryImpl) FindByUsername(username string) (model.Users, error) {
	var users model.Users
	result := u.Db.First(&users, "username = ? ", username)
	if result.Error != nil {
		return users, errors.New("invalid username or password")
	}
	return users, nil
}

// Save implements UsersRepository.
func (u *UserRepositoryImpl) Save(users model.Users) {
	result := u.Db.Create(&users)
	helper.ErrorPanic(result.Error)
}

// Update implements UsersRepository.
func (u *UserRepositoryImpl) Update(users model.Users) {
	var updateUsers = request.UpdateUserRequest{
		Id:       users.Id,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}
	result := u.Db.Model(&users).Updates(updateUsers)
	helper.ErrorPanic(result.Error)
}
