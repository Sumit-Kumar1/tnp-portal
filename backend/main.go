package main

import (
	"tnp-service/migrations"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.Migrate(migrations.All()) // run DB migrations

	app.Run()
}
