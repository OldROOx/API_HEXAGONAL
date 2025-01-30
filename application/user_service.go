package application

import (
	"ApiHEXaGONAL/config"
	"ApiHEXaGONAL/domain"
	"ApiHEXaGONAL/infrastructure/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// GetUsers obtiene todos los usuarios
func (s *UserService) GetUsers() ([]domain.User, error) {
	var users []domain.User
	result := config.DB.Find(&users)
	return users, result.Error
}

// CreateUser crea un nuevo usuario
func (s *UserService) CreateUser(user *domain.User) error {
	result := config.DB.Create(user)
	return result.Error
}

// GetUserByID obtiene un usuario por ID
func (s *UserService) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	result := config.DB.First(&user, id)
	return &user, result.Error
}

// UpdateUser actualiza un usuario existente
func (s *UserService) UpdateUser(user *domain.User) error {
	result := config.DB.Save(user)
	return result.Error
}

// DeleteUser elimina un usuario
func (s *UserService) DeleteUser(id uint) error {
	result := config.DB.Delete(&domain.User{}, id)
	return result.Error
}
