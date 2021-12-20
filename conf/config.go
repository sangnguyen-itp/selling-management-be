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
	SecurityAccessSecret string
	SecurityExpiry       time.Duration
	SecurityTLL          bool

	DatabaseHost      string
	DatabaseUser      string
	DatabasePassword  string
	DatabaseDBName    string
	DatabasePort      string
	DatabaseMigration bool
}

func (cfg *Config) reloadConfig() {
	cfg.SecurityAccessSecret = os.Getenv(SecurityAccessSecret)
	cfg.SecurityExpiry, _ = time.ParseDuration(os.Getenv(SecurityExpiryDuration))

	cfg.DatabaseHost = os.Getenv(DatabaseHost)
	cfg.DatabaseUser = os.Getenv(DatabaseUser)
	cfg.DatabasePassword = os.Getenv(DatabasePassword)
	cfg.DatabaseDBName = os.Getenv(DatabaseDBName)
	cfg.DatabasePort = os.Getenv(DatabasePort)
	cfg.DatabaseMigration, _ = strconv.ParseBool(os.Getenv(DatabaseMigration))
}
