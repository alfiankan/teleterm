package main

import (
	"context"
	"log"
	"os"

	"github.com/alfiankan/teleterm/teleterm"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(path)
	ctx := context.Background()
	db := teleterm.NewSqliteConnection(path + "/teleterm.db")
	teleterm.Start(ctx, db, "5438204586:AAFxj59Op6BX9IcHqFiS0su7zziH9QIsQNk")
}
