package validate

import (
	"errors"

	"github.com/pulpfree/univsales-pdf-url/model"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

// RequestInput function
func RequestInput(r *model.Request) (err error) {

	if r == nil {
		err = errors.New("Missing Request values")
		return err
	}

	validate = validator.New()

	err = validate.Struct(r)
	if err != nil {
		return err
	}

	// Ensure that if type is quote, we have Version
	if r.Type == "quote" && r.Version <= 0 {
		err = errors.New("Missing Version value")
		return err
	}

	return nil
}
