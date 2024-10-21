package presenters

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/validator"
)

func JwtSuccess(data string) map[string]interface{} {
	return output{
		"error":   nil,
		"message": data,
		"success": true,
	}
}

func JwtError(data entity.ValidateJwt) errorOuput {
	return errorOuput{
		"error": "Invalid: You must provide a valid token.",
		"message": map[string]interface{}{
			"token": data.Token,
		},
		"success": false,
	}
}

func JwtValidField(data validator.FieldValidation) errorOuput {
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

func HeaderFailed() errorOuput {
	return errorOuput{
		"error":   "Invalid: Authorization header format",
		"message": "",
		"success": false,
	}
}
