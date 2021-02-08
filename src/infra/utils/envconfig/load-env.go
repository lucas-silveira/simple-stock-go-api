package envconfig

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Load a env configuration
func Load() error {
	envFile := fmt.Sprintf("%s.env", os.Getenv("APP_ENV"))
	err := godotenv.Load(envFile)

	if err != nil {
		return err
	}

	return nil
}
