package conf

import (
	"github.com/joho/godotenv"
	"os"
	"selling-management-be/defined"
	"selling-management-be/pkg/logger"
	"strconv"
	"time"
)

var EnvConfig *Config

func init() {
	if err := godotenv.Load(); err != nil {
		logger.Log().Error(defined.SystemDomain, "godotenv.Load", err)
		panic(err)
	}

	config := new(Config)
	config.reloadConfig()
	EnvConfig = config
}

type Config struct {
	AppPort string

	SecurityAccessSecret string
	SecurityExpiry       time.Duration
	SecurityTLL          bool

	DatabaseHost      string
	DatabaseUser      string
	DatabasePassword  string
	DatabaseDBName    string
	DatabasePort      string
	DatabaseMigration bool

	DefaultUsername    string
	DefaultPassword    string
	DefaultEmail       string
	DefaultRole        string
	DefaultFirstName   string
	DefaultLastName    string
	DefaultPhoneNumber string
	DefaultAddress     string
}

func (cfg *Config) reloadConfig() {
	cfg.AppPort = os.Getenv(appPort)

	cfg.SecurityAccessSecret = os.Getenv(securityAccessSecret)
	cfg.SecurityExpiry, _ = time.ParseDuration(os.Getenv(securityExpiryDuration))

	cfg.DatabaseHost = os.Getenv(databaseHost)
	cfg.DatabaseUser = os.Getenv(databaseUser)
	cfg.DatabasePassword = os.Getenv(databasePassword)
	cfg.DatabaseDBName = os.Getenv(databaseDBName)
	cfg.DatabasePort = os.Getenv(databasePort)
	cfg.DatabaseMigration, _ = strconv.ParseBool(os.Getenv(databaseMigration))

	cfg.DefaultUsername = os.Getenv(defaultUsername)
	cfg.DefaultPassword = os.Getenv(defaultPassword)
	cfg.DefaultEmail = os.Getenv(defaultEmail)
	cfg.DefaultRole = os.Getenv(defaultRole)
	cfg.DefaultFirstName = os.Getenv(defaultFirstName)
	cfg.DefaultLastName = os.Getenv(defaultLastName)
	cfg.DefaultPhoneNumber = os.Getenv(defaultPhoneNumber)
	cfg.DefaultAddress = os.Getenv(defaultAddress)
}
