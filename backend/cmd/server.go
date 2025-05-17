package cmd

import (
	documenthttp "tnp-service/internal/handler/document"
	documentsvc "tnp-service/internal/service/document"
	documentstore "tnp-service/internal/store/document"
	"tnp-service/migrations"

	"gofr.dev/pkg/gofr"
)

func Server() {
	app := gofr.New()

	app.Migrate(migrations.All())
	setupRoutes(app)
	app.Run()
}

func setupRoutes(app *gofr.App) {
	// document API
	dh := setupDocumentAPI()
	app.POST("/v1/document", dh.CreateDoc)
	app.GET("/v1/document/{docId}", dh.GetDoc)

}

func setupDocumentAPI() *documenthttp.DocHandler {
	ds := documentstore.New()
	dsvc := documentsvc.New(ds)
	dh := documenthttp.New(dsvc)

	return dh
}
