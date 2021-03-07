package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	r := gin.Default()
	r.Use(secureHeaders())
	r.Use(sessions.Sessions("mysession", *app.store))

	r.GET("/", app.home)
	r.GET("/snippet/show/:id", app.showSnippet)
	r.GET("/snippet/create", app.createSnippetForm)
	r.POST("/snippet/create", app.createSnippet)

	r.Static("/static", "./ui/static")
	return r
}
