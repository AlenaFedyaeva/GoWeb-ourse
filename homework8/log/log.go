package log

import (
	"os"

	"github.vom/rs/zerolog"
)
	

func NewFileLogger(fname string) (*zerolog.Logger, error){
	f,err:=os.OpenFile(fname,os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil{
		return nil,err
	}
	logger:= zerolog.New(f).With().Timestamp().Logger()
	return &logger,nil
}