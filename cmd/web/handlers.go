package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"masrurimz/snippetbox/pkg/models"
)

// home handle home routes then display home page
// which contain all latest not expired snippets
func (app *application) home(c *gin.Context) {
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(c, err)
		return
	}

	app.render(c, "home.page.tmpl", &templateData{
		Snippets: s,
	})
}

// Add a showSnippet handler function.
func (app *application) showSnippet(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id < 1 {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	s, err := app.snippets.Get(id)
	fmt.Println(err)
	if err == models.ErrNoRecord {
		app.notFound(c)
		return
	} else if err != nil {
		app.serverError(c, err)
		return
	}

	app.render(c, "show.page.tmpl", &templateData{
		Snippet: s,
	})
}

// Login blbala
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func (app *application) createSnippet(c *gin.Context) {
	var form models.SnippetValidator

	// Do validation
	if err := models.ValidateSnippet(c, form); err != nil {
		app.render(c, "create.page.tmpl", &templateData{
			FormData: c.Request.PostForm,
			FormErrors: map[string]string{
				"title":   err["SnippetValidator.Title"],
				"content": err["SnippetValidator.Content"],
				"expires": err["SnippetValidator.Expires"],
			},
		})
		return
	}

	// Insert to database
	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(c, err)
		return
	}

	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/snippet/show/%d", id))
}

func (app *application) createSnippetForm(c *gin.Context) {
	app.render(c, "create.page.tmpl", nil)
}
