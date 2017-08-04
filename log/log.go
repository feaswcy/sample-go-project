package log

import (
	"fmt"
	"log"
	"os"
)

func registerLog() *log.Logger {

	logfile, err := os.OpenFile("/Users/didi/didi.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)

	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}

	//defer logfile.Close()

	Logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)

	Logger.Println("logger start...")

	return Logger

}
