package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
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
	Title   string `validate:"required,lt=100"`
	Content string `validate:"required"`
	Expires int    `validate:"required,eq=1|eq=7|eq=365"`
}

// ValidateSnippet validate snippet model
func ValidateSnippet(s *SnippetValidator) error {
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return err
		}
		return err
	}

	return nil
}
