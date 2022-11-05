package logger

import (
	"fmt"
	"time"
)

func Log(message string) {
	fmt.Printf("[%s] %s\n", time.Now().Format(time.RFC3339), message)
}

func Error(err error) {
	Log("ERR! " + err.Error())
}
