package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

func (app *application) addDefauldData(td *templateData, c *gin.Context) *templateData {
	if td == nil {
		td = &templateData{}
	}

	td.CurrentYear = time.Now().Year()

	return td
}

// The serverError helper writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(c *gin.Context, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	app.errorLog.Output(2, trace)

	c.AbortWithError(
		http.StatusInternalServerError,
		err,
	)
}

// The clientError helper sends a specific status code and corresponding description
// to the user. We'll use this later in the book to send responses like 400 "Bad
// Request" when there's a problem with the request that the user sent.
func (app *application) clientError(c *gin.Context, status int) {
	c.AbortWithStatus(status)
	// http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found response to
// the user.
func (app *application) notFound(c *gin.Context) {
	app.clientError(c, http.StatusNotFound)
}

func (app *application) render(c *gin.Context, name string, td *templateData) {
	// Retrieve cached template using name as keyword
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(c, fmt.Errorf("The template %s dose not exist", name))
		return
	}

	// Render template with dynamic data to buffer
	buf := new(bytes.Buffer)
	err := ts.Execute(buf, app.addDefauldData(td, c))
	if err != nil {
		app.serverError(c, err)
		return
	}

	// If no template error then pass rendered data to client
	buf.WriteTo(c.Writer)
}
