package chanlog

import (
	"log"
)

func Print(v ...interface{}) {
	log.Print(v...)
}

func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
