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
var (
	ErrNoRecord          = errors.New("models: no matching record found")
	ErrInvalidCredential = errors.New("models: invalid credential")
	ErrDuplicatedEmail   = errors.New("models: duplicated email")
)

var (
	trans ut.Translator
)

// TranslatorSetup initiate translator config
func TranslatorSetup() bool {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	valid, ok := binding.Validator.Engine().(*validator.Validate)

	en_translations.RegisterDefaultTranslations(valid, trans)

	return ok
}

func validate(c *gin.Context, vs interface{}) validator.ValidationErrorsTranslations {
	if err := c.ShouldBind(vs); err != nil {
		errs := err.(validator.ValidationErrors).Translate(trans)

		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return errs
		}

		return errs
	}
	return nil
}

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

// ValidateSnippet validate snippet model
func ValidateSnippet(c *gin.Context, s *SnippetValidator) validator.ValidationErrorsTranslations {
	return validate(c, s)
}

// User struct model for authentication
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// UserValidator struct for validate user account form input
type UserValidator struct {
	Name     string `form:"name" json:"name" binding:"required,lt=255"`
	Email    string `form:"email" json:"email" binding:"required,lt=255"`
	Password string `form:"password" json:"password" binding:"required,lt=60,gt=6"`
}

// ValidateUser validate user form data
func ValidateUser(c *gin.Context, user *UserValidator) validator.ValidationErrorsTranslations {
	return validate(c, user)
}
