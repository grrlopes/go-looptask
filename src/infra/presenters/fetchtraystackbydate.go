package presenters

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
)

func FetchTrayStackByDateSuccess(label []entity.LabelStackAggSet) output {
	return output{
		"error":   nil,
		"message": label,
		"success": true,
	}
}

func FetchTrayStackByDateError(err error) output {
	return output{
		"error":   err.Error(),
		"message": nil,
		"success": false,
	}
}
