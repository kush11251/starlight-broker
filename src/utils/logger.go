package main

import (
	"log"
)

func InitLogger() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}