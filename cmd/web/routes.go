package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	r := gin.Default()
	r.Use(secureHeaders())

	r.GET("/", app.home)
	r.GET("/snippet/show/:id", app.showSnippet)
	r.GET("/snippet/create", app.createSnippetForm)
	r.POST("/snippet/create", app.createSnippet)

	r.Static("/static", "./ui/static")
	return r
}
