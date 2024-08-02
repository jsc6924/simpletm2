package config

import (
	"os"
)

var (
	DBPath          = "data/SimpleTM.db"
	JwtKey          = []byte("test-secret-key")
	CorsAllowOrigin = "http://localhost:3000"
	SetCookieDomain = ""
)

func init() {
	if key, exists := os.LookupEnv("JWT_SECRET_KEY"); exists {
		JwtKey = []byte(key)
	}
	if key, exists := os.LookupEnv("CORS_ALLOW_ORIGIN"); exists {
		CorsAllowOrigin = key
	}
	if key, exists := os.LookupEnv("SET_COOKIE_DOMAIN"); exists {
		SetCookieDomain = key
	}
	if key, exists := os.LookupEnv("DB_PATH"); exists {
		DBPath = key
	}
}
