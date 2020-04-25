package travellist

import "os"

func IsProduction() bool {
	return os.Getenv("APP_ENV") == "production"
}
