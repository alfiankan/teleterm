package main

import (
	"github.com/joho/godotenv"
	"teleterm/handler"
)

var teleToken string

func main() {
	godotenv.Load(".env")
	handler.Begin()
}
