package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// ErrNoRecord Error message for no database records with the given ID
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet struct model
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetValidator struct to be used with golang validator
type SnippetValidator struct {
	Title   string `form:"title" json:"title" binding:"required,lt=100"`
	Content string `form:"content" json:"content" binding:"required"`
	Expires int    `form:"expires" json:"expires" binding:"required,eq=1|eq=7|eq=365"`
}

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

// ValidateSnippet validate snippet model
func ValidateSnippet(c *gin.Context, s SnippetValidator) validator.ValidationErrorsTranslations {
	v, ok := binding.Validator.Engine().(*validator.Validate)

	if !ok {
		return nil
	}

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(v, trans)

	if err := c.ShouldBind(&s); err != nil {
		errs := err.(validator.ValidationErrors).Translate(trans)

		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return errs
		}

		return errs
	}

	return nil
}
