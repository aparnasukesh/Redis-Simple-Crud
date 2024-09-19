package repository

import (
	"redis_sample_project/config"
	"redis_sample_project/models"
)

func CreateUser(user models.User) error {
	return config.DB.Create(&user).Error
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return user, err
}

func UpdateUser(user models.User) error {
	return config.DB.Save(&user).Error
}

func DeleteUser(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}
