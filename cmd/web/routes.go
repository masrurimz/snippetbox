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

	snippet := r.Group("/snippet")
	{
		snippet.GET("/show/:id", app.showSnippet)
		snippet.GET("/create", app.createSnippetForm)
		snippet.POST("/create", app.createSnippet)
	}

	auth := r.Group("/user")
	{
		auth.GET("/signup", app.signupUserForm)
		auth.POST("/signup", app.signupUser)
		auth.GET("/login", app.loginUserForm)
		auth.POST("/login", app.loginUser)
		auth.POST("/logout", app.logoutUser)
	}

	r.Static("/static", "./ui/static")
	return r
}
