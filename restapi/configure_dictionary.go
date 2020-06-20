// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"dictionary/models"
	"github.com/go-openapi/swag"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"dictionary/restapi/operations"
	"dictionary/restapi/operations/words"
)

//go:generate swagger generate server --target ..\..\dictionary --name Dictionary --spec ..\api\swagger.yaml

func configureFlags(api *operations.DictionaryAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DictionaryAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.WordsAddWordHandler = words.AddWordHandlerFunc(func(params words.AddWordParams) middleware.Responder {
		id := new(uint64)
		*id = uint64(1)
		return words.NewAddWordOK().WithPayload(&words.AddWordOKBody{ID: id})
	})

	api.WordsGetWordsHandler = words.GetWordsHandlerFunc(func(params words.GetWordsParams) middleware.Responder {
		w := make([]*models.Word, *params.Limit)

		return words.NewGetWordsOK().WithPayload(w)
	})

	api.WordsDeleteWordHandler = words.DeleteWordHandlerFunc(func(params words.DeleteWordParams) middleware.Responder {
		id := new(uint64)
		*id = params.WordID

		return words.NewDeleteWordOK().WithPayload(&words.DeleteWordOKBody{ID: id})
	})

	api.WordsGetWordHandler = words.GetWordHandlerFunc(func(params words.GetWordParams) middleware.Responder {
		w := models.Word{
			ID:            params.WordID,
			Origin:        swag.String("Угорщина"),
			Samples:       make([]string, 2),
			Transcription: swag.String("[с'ійо]"),
			Translations:  make([]*models.Translation, 2),
			Word:          swag.String("Сійо"),
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
