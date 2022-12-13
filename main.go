// main.go
package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	a := App{}

	username := os.Getenv("APP_DB_USERNAME")
	password := os.Getenv("APP_DB_PASSWORD")
	db_name := os.Getenv("APP_DB_NAME")

	a.Initialize(
		username,
		password,
		db_name,
	)

	a.Run(":8010")
}
