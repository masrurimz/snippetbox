package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our application receives.
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := httprouter.New()
	mux.GET("/", app.home)
	mux.GET("/snippet/show/:id", app.showSnippet)
	mux.GET("/snippet/create", app.createSnippetForm)
	mux.POST("/snippet/create", app.createSnippet)

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project // directory root.
	mux.ServeFiles("/static/*filepath", http.Dir("./ui/static"))

	return standardMiddleware.Then(mux)
}
