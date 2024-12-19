package util

import (
	"log"
	"os"
	"time"
)

func ErrLog(v ...any) {
	log.Println(v...)
	time.Sleep(time.Second * 10)
	os.Exit(1)
}
