// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"dictionary/api/models"
	"dictionary/api/restapi/operations"
	"dictionary/api/restapi/operations/words"
	"dictionary/app"
	"dictionary/app/facilities"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

//go:generate swagger generate server --target ..\..\dictionary --name Dictionary --spec ..\api\swagger.yaml

func configureFlags(api *operations.DictionaryAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DictionaryAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	conf, err := facilities.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	dep := app.NewDependency(conf)

	if err := dep.DBMigrate.Run(); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.WordsAddWordHandler = words.AddWordHandlerFunc(func(params words.AddWordParams) middleware.Responder {
		return words.NewAddWordOK().WithPayload(&words.AddWordOKBody{ID: swag.Uint64(1)})
	})

	api.WordsGetWordsHandler = words.GetWordsHandlerFunc(func(params words.GetWordsParams) middleware.Responder {
		w := make([]*models.Word, *params.Limit)

		for i := range w {
			w[i] = &models.Word{
				ID:            uint64(i),
				Origin:        swag.String("Польща"),
				Samples:       []string{"Канапка дуже смачна.", "Канапка не дуже смачна."},
				Transcription: swag.String("канапка"),
				Translations: []*models.Translation{
					&models.Translation{
						Transcription: swag.String("бутерброд"),
						Translation:   swag.String("Бутерброд"),
					},
				},
				Word: swag.String("Канапка"),
			}
		}

		return words.NewGetWordsOK().WithPayload(w)
	})

	api.WordsDeleteWordHandler = words.DeleteWordHandlerFunc(func(params words.DeleteWordParams) middleware.Responder {
		return words.NewDeleteWordOK().WithPayload(&words.DeleteWordOKBody{ID: &params.WordID})
	})

	api.WordsGetWordHandler = words.GetWordHandlerFunc(func(params words.GetWordParams) middleware.Responder {
		w := models.Word{
			ID:            params.WordID,
			Origin:        swag.String("Угорщина"),
			Samples:       []string{"Сійо старий.", "Сійо, до завтра."},
			Transcription: swag.String("[с'ійо]"),
			Translations: []*models.Translation{
				&models.Translation{
					Transcription: swag.String("прив'іт"),
					Translation:   swag.String("Привіт"),
				},
				&models.Translation{
					Transcription: swag.String("пока"),
					Translation:   swag.String("Пока"),
				},
			},
			Word: swag.String("Сійо"),
		}

		return words.NewGetWordOK().WithPayload(&w)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
