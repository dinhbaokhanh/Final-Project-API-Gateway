package config

import "os"

const (
	defaultPort          = "8080"
	defaultUsersBackend  = "http://localhost:9001"
	defaultOrdersBackend = "http://localhost:9002"
)

// Config contains runtime settings for the API gateway.
type Config struct {
	Port          string
	UsersBackend  string
	OrdersBackend string
}

func Load() Config {
	return Config{
		Port:          getEnv("PORT", defaultPort),
		UsersBackend:  getEnv("BACKEND_USERS", defaultUsersBackend),
		OrdersBackend: getEnv("BACKEND_ORDERS", defaultOrdersBackend),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
