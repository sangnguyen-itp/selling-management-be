package service

import (
	"gorm.io/gorm"
	"selling-management-be/conf"
	"selling-management-be/defined"
	"selling-management-be/model"
	"selling-management-be/pkg/database"
	"selling-management-be/pkg/logger"
)

var mainService *Service

type Service struct {
	db *gorm.DB
}

type TransactionService struct {
	tx *gorm.DB
}

func NewService() error {
	_db := database.ConnectDB()
	mainService = &Service{db: _db}
	if conf.EnvConfig.DatabaseMigration {
		logger.Log().Info(defined.SystemDomain, "MigrationSchema")
		mainService.MigrationSchema()
	}
	return nil
}

func (s *Service) MigrationSchema() error {
	return s.db.AutoMigrate(&model.User{})
}