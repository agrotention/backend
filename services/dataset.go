package services

import "gorm.io/gorm"

type DatasetService struct {
	db *gorm.DB
}
