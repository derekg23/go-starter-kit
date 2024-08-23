package controllers

import (
	"errors"
	"gorm.io/gorm"

	"untitledgoproject/models"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}

func (uc *UserController) GetAllUsers() []models.User {
	var users []models.User
	uc.db.Find(&users)
	return users
}

func (uc *UserController) CreateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("Name and email are required")
	}
	var existingUser models.User
	if err := uc.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("User with this email already exists")
	}
	return uc.db.Create(user).Error
}

func (uc *UserController) EditUser(id uint, user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("Name and email are required")
	}
	var existingUser models.User
	if err := uc.db.Where("id = ?", id).First(&existingUser).Error; err != nil {
		return err
	}
	if existingUser.Email != user.Email {
		if err := uc.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
			return errors.New("User with this email already exists")
		}
	}
	if err := uc.db.Model(&existingUser).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (uc *UserController) DeleteUser(id uint) error {
	var user models.User
	result := uc.db.Where("id = ?", id).Delete(&user)
	return result.Error
}