package models

import (
	"time"

	"github.com/agrotention/backend/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"column:id"`
	Email     string         `gorm:"column:email"`
	Password  string         `gorm:"column:password"`
	UserInfo  *UserInfo      `gorm:"foreignKey:UserID;references:ID"`
	UserRole  *UserRole      `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

type UserInfo struct {
	ID            string     `gorm:"column:id"`
	UserID        string     `gorm:"column:user_id"`
	DisplayName   *string    `gorm:"column:display_name"`
	PhotoUrl      *string    `gorm:"column:photo_url"`
	Gender        *string    `gorm:"column:gender"`
	DateOfBirth   *time.Time `gorm:"column:date_of_birth;type:date"`
	Comany        *string    `gorm:"column:company"`
	RoleAtCompany *string    `gorm:"column:role_at_company"`
}

type UserRole struct {
	ID       string `gorm:"column:id"`
	UserID   string `gorm:"column:user_id"`
	IsAdmin  bool   `gorm:"column:is_admin"`
	IsStaff  bool   `gorm:"column:is_staff"`
	IsExpert bool   `gorm:"column:is_expert"`
}

// Before Create User Hook
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// Hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		utils.LogErr.Println(err)
		return err
	}
	u.Password = string(hashPassword)

	// Generate UUID
	userID, err := uuid.NewV7()
	if err != nil {
		utils.LogErr.Println(err)
		return err
	}
	infoID, err := uuid.NewV7()
	if err != nil {
		utils.LogErr.Println(err)
		return err
	}
	roleID, err := uuid.NewV7()
	if err != nil {
		utils.LogErr.Println(err)
		return err
	}

	// Create user ID
	u.ID = userID.String()

	// Create User Info
	u.UserInfo.ID = infoID.String()

	// Create User Role
	u.UserRole.ID = roleID.String()
	u.UserRole.IsAdmin = false
	u.UserRole.IsExpert = false
	u.UserRole.IsStaff = false

	return nil

}
