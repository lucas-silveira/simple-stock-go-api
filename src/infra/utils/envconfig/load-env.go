package envconfig

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Load a env configuration
func Load() error {
	AppEnv := os.Getenv("APP_ENV")

	if AppEnv != "test" {
		envFile := fmt.Sprintf("%s.env", AppEnv)
		err := godotenv.Load(envFile)

		if err != nil {
			return err
		}
	}

	return nil
}
