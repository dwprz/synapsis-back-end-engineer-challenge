package v

import (
	e "book-category-service/src/model/entity"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()

	Validate.RegisterValidation("bookcategory", validateBookCategory)
}

func validateBookCategory(fl validator.FieldLevel) bool {
	category := e.Category(fl.Field().String())
	switch category {
	case e.FICTION, e.NON_FICTION, e.SCIENCE, e.HISTORY, e.BIOGRAPHY, e.TECHNOLOGY,
		e.FANTASY, e.MYSTERY, e.THRILLER, e.CHILDREN, e.YOUNG_ADULT, e.ROMANCE, e.ADVENTURE:
		return true
	}
	return false
}
