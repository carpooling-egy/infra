package internal

import "os"

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func mask(s string) string {
	if s == "" {
		return ""
	}
	return "*****"
}
