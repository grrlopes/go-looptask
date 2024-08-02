package presenters

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/validator"
)

func LoginSuccess(data map[string]interface{}) map[string]interface{} {
	return output{
		"error":   nil,
		"message": data,
		"success": true,
	}
}

func LoginError(data entity.Users) errorOuput {
	return errorOuput{
		"error": "Invalid: You must provide a valid credencials.",
		"message": map[string]interface{}{
			"email":    data.Email,
			"password": data.Password,
		},
		"success": false,
	}
}

func LoginValidField(data validator.FieldValidation) errorOuput {
	errorout := []string{}

	for _, v := range data.Message {
		errorout = append(errorout, v.Error())
	}

	return errorOuput{
		"error":   data.Error,
		"message": errorout,
		"success": false,
	}
}
