package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"selling-management-be/conf"
	"selling-management-be/defined/domain"
	"selling-management-be/pkg/logger"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.EnvConfig.DatabaseHost,
		conf.EnvConfig.DatabaseUser,
		conf.EnvConfig.DatabasePassword,
		conf.EnvConfig.DatabaseDBName,
		conf.EnvConfig.DatabasePort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logger.Log().Error(domain.SystemDomain, "gorm.Open", err)
		panic(err)
	}

	return db.Debug()
}