package config

import (
	"os"
	"strconv"
)

type Config struct {
	BackendPort      string
	BackendEnv       string
	LogLevel         string
	DBHost           string
	DBPort           string
	DBName           string
	DBUser           string
	DBPassword       string
	JWTSecret        string
	JWTExpiry        string
	TickIntervalSecs int64
	MaxTicks         int64
	MaxTickBox       int64
}

func NewConfig() *Config {
	return &Config{
		BackendPort:      getEnv("BACKEND_PORT", "8080"),
		BackendEnv:       getEnv("BACKEND_ENV", "development"),
		LogLevel:         getEnv("LOG_LEVEL", "info"),
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "27017"),
		DBName:           getEnv("DB_NAME", "ticks_and_tomes"),
		DBUser:           getEnv("DB_USER", ""),
		DBPassword:       getEnv("DB_PASSWORD", ""),
		JWTSecret:        getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		JWTExpiry:        getEnv("JWT_EXPIRY", "24h"),
		TickIntervalSecs: getEnvInt("TICK_INTERVAL_SECONDS", 300),
		MaxTicks:         getEnvInt("MAX_TICKS", 500),
		MaxTickBox:       getEnvInt("MAX_TICK_BOX", 200),
	}
}

func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int64) int64 {
	val := getEnv(key, "")
	if val == "" {
		return defaultVal
	}
	if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
		return intVal
	}
	return defaultVal
}
