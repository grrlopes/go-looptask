package presenters

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/validator"
)

type errorOuput map[string]interface{}
type output map[string]interface{}

type MongoOutput struct {
	TotalRows int              `json:"total_rows"`
	Offset    int              `json:"offset"`
	Data      []entity.Labeled `json:"data"`
}

func ValidFieldResponse(data validator.FieldValidation) errorOuput {
	looptask := []string{}

	for _, v := range data.Message {
		looptask = append(looptask, v.Error())
	}

	return errorOuput{
		"Error":   data.Error,
		"Message": looptask,
	}
}

func SuccessResponse(data entity.MongoResul) output {
  return output{
    "message": data.Reason,
  }
}

func ErrorResponse(data entity.MongoResul) errorOuput {
	return errorOuput{
		"Error":   data.Error,
		"Message": data.Reason,
	}
}
