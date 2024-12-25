package services

import (
	"github.com/agrotention/backend/dto"
	"github.com/agrotention/backend/models"
	"github.com/agrotention/backend/utils"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// =========== Type
type UserService struct {
	db *gorm.DB
}

// =========== Handlers
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

// =========== Services
func (s *UserService) ChangeEmail(data dto.ReqUserChangeEmail) (*dto.ResUserChangeEmail, utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}
func (s *UserService) ChangePassword(data dto.ReqUserChangePassword) (*dto.ResUserChangePassword, utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}
func (s *UserService) Delete(data dto.ReqUserDelete) (*dto.ResUserDelete, utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}
func (s *UserService) Detail(data dto.ReqUserDetail) (*dto.ResUserDetail, utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}
func (s *UserService) Disable(data dto.ReqUserDisable) (*dto.ResUserDisable, utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}
func (s *UserService) List(data dto.ReqUserList) (*dto.ResUserList, utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}
func (s *UserService) Login(data dto.ReqUserLogin) (*dto.ResUserLogin, utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}

func (s *UserService) Register(data dto.ReqUserRegister) (*dto.ResUserRegister, utils.HTTPError) {
	// Validate Request
	err := utils.Validate.Struct(data)
	if err != nil {
		if validationError, ok := err.(validator.ValidationErrors); ok {
			return nil, utils.TranslateValidationError(validationError)
		} else {
			return nil, utils.ErrInternal
		}
	}
	// Check Unique Email
	count, err := s.countEmail(data.Email)
	if err != nil {
		return nil, utils.ErrInternal
	}
	if count != 0 {
		return nil, utils.NewErrWithMessage(409, "email already exist")
	}

	// Create New user
	user := models.User{
		Email:    data.Email,
		Password: data.Password,
		UserInfo: &models.UserInfo{
			DisplayName: &data.DisplayName,
		},
	}

	// Query
	err = s.db.Create(&user).Error
	if err != nil {
		return nil, utils.ErrInternal
	}
	return &dto.ResUserRegister{ID: user.ID}, nil
}

func (s *UserService) Update(data dto.ReqUserUpdate) (*dto.ResUserUpdate, utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}

// Helper
func (s *UserService) countEmail(email string) (int64, error) {
	var count int64
	err := s.db.Model(models.User{}).Where("email = ?", email).Count(&count).Error
	return count, err
}
