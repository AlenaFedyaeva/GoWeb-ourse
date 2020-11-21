package config

import (
	"fmt"
	"os"
)


func LookupEnv(env string) (string,bool) {
    fpath, exists := os.LookupEnv(env)

    if exists {
        // Print the value of the environment variable
    	fmt.Print(fpath)
   	}
   return fpath,exists
}

func SetEnv(env string, val string) {
    // Set the USERNAME environment variable to "MattDaemon"
    os.Setenv(env, val)
}