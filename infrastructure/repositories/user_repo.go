package repositories

import (
	"ApiHEXaGONAL/config"
	"ApiHEXaGONAL/domain"
)

func GetUsers() []domain.User {
	var users []domain.User
	config.DB.Find(&users)
	return users
}

func CreateUser(user *domain.User) error {
	result := config.DB.Create(user)
	return result.Error
}

func GetUserByID(id uint) (domain.User, error) {
	var user domain.User
	result := config.DB.First(&user, id)
	return user, result.Error
}

func UpdateUser(user *domain.User) error {
	result := config.DB.Save(user)
	return result.Error
}

func DeleteUser(id uint) error {
	result := config.DB.Delete(&domain.User{}, id)
	return result.Error
}

func GetUserProducts(userID uint) ([]domain.Product, error) {
	var products []domain.Product
	result := config.DB.Where("user_id = ?", userID).Find(&products) // Busca productos por user_id
	return products, result.Error
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	result := config.DB.Create(user)
	return result.Error
}

type UserRepository struct {
	// Puedes agregar campos aquí si es necesario
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUserProducts(userID uint) ([]domain.Product, error) {
	var products []domain.Product
	result := config.DB.Where("user_id = ?", userID).Find(&products)
	return products, result.Error
}

func (r *UserRepository) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	result := config.DB.Preload("Products").First(&user, id)
	return &user, result.Error
}
