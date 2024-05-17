package utils

import (
	"log"
)

func LogSomething(message string, err any, logType int) {

	log.Println(message, err)
}
