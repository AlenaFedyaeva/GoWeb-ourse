package config

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func (cfg *FileConfig) ParseFlags() (error) {
	logNameF:=flag.String("log","log.tmp","file name")
	portF := flag.String("port","9090","a string")

	flag.Parse()

	// 1) Read port from args 
	if *portF==""{
		str:="Empty flag port"
		log.Fatal(str)
		return errors.New(str)
	}

	if *logNameF==""{
		str:="Empty flag logName"
		log.Fatal(str)
		return errors.New(str)
	}
	cfg.Port=*portF
	cfg.LogName=*logNameF
	// fmt.Println("Your flags: ", *logNameF,"port: ",*portF)
	return nil
}

func ValidateConfigPath(path string) error {
    s, err := os.Stat(path)
    if err != nil {
        return err
    }
    if s.IsDir() {
        return fmt.Errorf("'%s' is a directory, not a normal file", path)
    }
    return nil
}