package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const searchAPIAuthKeyVariableName = "SEARCH_API_AUTH_KEY"

func GetSearchAPIAuthKey() (string, error) {
	godotenv.Load("./.env")
	key := os.Getenv(searchAPIAuthKeyVariableName)
	if key == "" {
		return "", ConfigMissingError{configName: searchAPIAuthKeyVariableName}
	}
	return key, nil
}

type ConfigMissingError struct {
	configName string
}

func (e ConfigMissingError) Error() string {
	return fmt.Sprintf("Missing config: %s", e.configName)
}
