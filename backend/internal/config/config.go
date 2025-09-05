package config

import (
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	Environment string // Values other than "development" are considered production
	UserDataDir string // Path to TextBro user data directory
	AppDataDir  string // Path to TextBro application data directory
	Password    string // Application password
	ServerPort  string // Backend URL or port for development
}

func Load() (*Config, error) {
	cfg := &Config{
		Environment: getEnv("TEXTBRO_ENVIRONMENT"),
		UserDataDir: getEnv("TEXTBRO_USER_DATA_DIR"),
		AppDataDir:  getEnv("TEXTBRO_APP_DATA_DIR"),
		Password:    getEnv("TEXTBRO_PASSWORD"),
		ServerPort:  getEnv("TEXTBRO_SERVER_PORT"),
	}

	if cfg.Environment == "development" {
		exe, err := os.Executable()
		if err != nil {
			return nil, errors.New("failed determining the current working directory")
		}

		exePath := filepath.Dir(exe)

		if cfg.UserDataDir == "" {
			cfg.UserDataDir = filepath.Join(exePath, "user_data")
		}

		if cfg.AppDataDir == "" {
			cfg.AppDataDir = filepath.Join(exePath, "app_data")
		}

		if cfg.ServerPort == "" {
			cfg.ServerPort = "8080"
		}
	} else {
		if cfg.UserDataDir == "" {
			return nil, errors.New("root directory must be set (TEXTBRO_USER_DATA_DIR)")
		}

		if cfg.AppDataDir == "" {
			return nil, errors.New("app data directory must be set (TEXTBRO_APPDATA_DIR)")
		}

		if cfg.Password == "" {
			return nil, errors.New("password must be set in production mode (TEXTBRO_PASSWORD)")
		}
	}

	return cfg, nil
}

func getEnv(key string) string {
	value, _ := os.LookupEnv(key)
	return value
}
