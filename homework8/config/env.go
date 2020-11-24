package config

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func GetFileNameFromEnv(env string) (string, error) {
	fpath, exists := os.LookupEnv(env)

    if !exists{
        str:=fmt.Sprintf("environment variable %s do not exist", env)
        log.Println(str)
        return fpath, errors.New(str)
    }
    validErr := ValidateConfigPath(fpath)
	return fpath, validErr
}

func SetEnv(env string, val string) {
	// Set the USERNAME environment variable to "MattDaemon"
	os.Setenv(env, val)
}
