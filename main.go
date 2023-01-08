// main.go
package main

import (
	"os"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

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

	a.Router.HandleFunc("/swagger", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8010/swagger/doc.json"), //The url pointing to API definition
	))

}
