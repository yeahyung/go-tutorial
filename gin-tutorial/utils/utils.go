package utils

import "os"

func GetEnvOr(name string, fallback string) string {
	env := os.Getenv(name)

	if env == "" {
		env = fallback
	}

	return env
}
