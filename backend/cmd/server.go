package cmd

import (
	"tnp-service/internal/handler/document"
	documenthttp "tnp-service/internal/handler/document"
	"tnp-service/internal/handler/user"
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
	app.POST("/v1/document", dh.CreateDocument)
	app.GET("/v1/document/{docId}", dh.GetDocument)

	uh := setupUserAPI()
	app.POST("/v1/register", uh.CreateUser)
	app.POST("v1/login", uh.CreateUser)

}

func setupUserAPI() *user.Handler {
	return user.New()
}

func setupDocumentAPI() *documenthttp.Handler {
	ds := documentstore.New()
	dsvc := documentsvc.New(ds)
	dh := document.New(dsvc)

	return dh
}
