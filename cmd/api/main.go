package main

import (
	"log"

	app "github.com/Infamous003/smol-url/internal/app"
)

func main() {
	a := app.New()

	app.SetupRoutes(a.Router())

	if err := a.Start(":9090"); err != nil {
		log.Fatal(err)
	}
}
