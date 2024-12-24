package services

import (
	"github.com/agrotention/backend/dto"
	"github.com/agrotention/backend/utils"
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
func (s *UserService) ChangeEmail(data dto.ReqUserChangeEmail) (*dto.ResUserChangeEmail, *utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}
func (s *UserService) ChangePassword(data dto.ReqUserChangePassword) (*dto.ResUserChangePassword, *utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}
func (s *UserService) Delete(data dto.ReqUserDelete) (*dto.ResUserDelete, *utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}
func (s *UserService) Detail(data dto.ReqUserDetail) (*dto.ResUserDetail, *utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}
func (s *UserService) Disable(data dto.ReqUserDisable) (*dto.ResUserDisable, *utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}
func (s *UserService) List(data dto.ReqUserList) (*dto.ResUserList, *utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}
func (s *UserService) Login(data dto.ReqUserLogin) (*dto.ResUserLogin, *utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}
func (s *UserService) Register(data dto.ReqUserRegister) (*dto.ResUserRegister, *utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}
func (s *UserService) Update(data dto.ReqUserUpdate) (*dto.ResUserUpdate, *utils.HTTPError) {
	return nil, utils.ErrNotImplemented
}
