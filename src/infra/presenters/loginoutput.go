package presenters

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/validator"
)

type LoginMongoOutput struct {
	Data map[string]interface{} `json:"data"`
}

func LoginSuccess(data map[string]interface{}) LoginMongoOutput {
	return LoginMongoOutput{
		Data: data,
	}
}

func LoginError(data entity.Users) errorOuput {
	return errorOuput{
		"Error":   data,
		"Message": data,
	}
}

func LoginValidField(data validator.FieldValidation) errorOuput {
	mHoney := []string{}

	for _, v := range data.Message {
		mHoney = append(mHoney, v.Error())
	}

	return errorOuput{
		"Error":   data.Error,
		"Message": mHoney,
	}
}
