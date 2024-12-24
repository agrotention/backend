package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string
	Email     string
	Password  string
	Disabled  bool
	UserInfo  UserInfo
	UserRole  UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {

	// Create new uuid
	genUuid, err := uuid.NewV7()
	if err != nil {
		log.Println(err)
		return err
	}
	u.ID = genUuid.String()

	// Create new role
	u.UserRole = UserRole{
		IsAdmin:  false,
		IsStaff:  false,
		IsExpert: false,
	}

	return nil
}

type UserInfo struct {
	DisplayName   string
	PhotoUrl      *string
	Gender        *string
	DateOfBirth   *time.Time `gorm:"type:date"`
	Comany        *string
	RoleAtCompany *string
}

type UserRole struct {
	IsAdmin  bool
	IsStaff  bool
	IsExpert bool
}
