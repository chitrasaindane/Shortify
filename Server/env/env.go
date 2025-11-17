package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// # Load Env
func LoadEnv() {
	// # Load Env Variables
	err := godotenv.Load()
	if err != nil {
		// # Load Env Error
		fmt.Println("🚫 Load Env Error : ", err)
	}
}

// # Get Env
func GetEnv(key, defaultValue string) string {
	// # Get Value
	value := os.Getenv(key)
	if value == "" {
		// # Return Default Value
		return defaultValue
	}
	// # Return Value
	return value
}
