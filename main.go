// main.go

package main

func main() {
	a := App{}
	// os.Getenv("APP_DB_USERNAME"),
	// 	os.Getenv("APP_DB_PASSWORD"),
	// 	os.Getenv("APP_DB_NAME")
	a.Initialize(
		"postgres",
		"Goodman8349**",
		"smarket",
	)

	a.Run(":8010")
}
